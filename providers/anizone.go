package providers

import (
	"log"
	"net/http"
	"strings"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/interfaces"

	"github.com/PuerkitoBio/goquery"
)

func SearchAnizoneByName(title string) (bool, interfaces.AnizoneSearchResponse) {
	res, err := http.Get("https://anizone.to/anime?search=" + strings.Join(strings.Split(title, " "), "+"))
	if err != nil {
		log.Println("Error fetching URL:", err)
		return true, interfaces.AnizoneSearchResponse{}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
		return true, interfaces.AnizoneSearchResponse{}
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Error loading HTML document:", err)
		return true, interfaces.AnizoneSearchResponse{}
	}

	resp := make([]interfaces.AnizoneIndividualResult, 0)

	// Find the search result items
	doc.Find(".h-6.inline.truncate").Each(func(i int, s *goquery.Selection) {
		attrVal, exists := s.Find("a").Attr("href")
		if exists {
			animeId := strings.Split(strings.TrimPrefix(attrVal, "/"), "?")[0]
			title := strings.TrimSpace(s.Find("a").Text())
			resp = append(resp, interfaces.AnizoneIndividualResult{
				Title: title, // Assuming title in English as an example
				ID:    animeId,
			})
		}
	})

	// Construct the final response object
	responseObject := interfaces.AnizoneSearchResponse{Result: resp}
	return false, responseObject
}

func SearchAnizoneByNameAndReturnBestMatch(title string) (bool, any) {
	hasError, responseObject := SearchAnizoneByName(title)

	if hasError || len(responseObject.Result) == 0 {
		return true, map[string]any{}
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.Title, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	return false, responseObject.Result[currIndex]
}

func SearchAnizoneByNameAndReturnBestMatchAsync(title string, result chan<- map[string]any, errors chan<- map[string]bool) {
	hasError, responseObject := SearchAnizoneByName(title)

	if hasError || len(responseObject.Result) == 0 {
		errors <- map[string]bool{"anizone": true}
		return
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.Title, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	// Send result
	result <- map[string]any{"anizone": responseObject.Result[currIndex]}
}
