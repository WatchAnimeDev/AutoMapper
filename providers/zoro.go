package providers

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/interfaces"

	"github.com/PuerkitoBio/goquery"
)

func SearchZoroByName(title string) (bool, interfaces.ZoroSearchResponse) {
	res, err := http.Get("https://hianime.to/search?keyword=" + url.QueryEscape(title))
	if err != nil {
		log.Println("Error fetching URL:", err)
		return true, interfaces.ZoroSearchResponse{}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Status code error: %d %s", res.StatusCode, res.Status)
		return true, interfaces.ZoroSearchResponse{}
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Error loading HTML document:", err)
		return true, interfaces.ZoroSearchResponse{}
	}

	resp := make([]interfaces.ZoroIndividualResult, 0)

	// Find the search result items
	doc.Find(".flw-item").Each(func(i int, s *goquery.Selection) {
		attrVal, exists := s.Find(".film-name a").Attr("href")
		if exists {
			animeId := strings.Split(strings.TrimPrefix(attrVal, "/"), "?")[0]
			title := s.Find(".film-name a").Text()
			resp = append(resp, interfaces.ZoroIndividualResult{
				Title: title, // Assuming title in English as an example
				ID:    animeId,
			})
		}
	})

	// Construct the final response object
	responseObject := interfaces.ZoroSearchResponse{Result: resp}
	return false, responseObject
}

func SearchZoroByNameAndReturnBestMatch(title string) (bool, any) {
	hasError, responseObject := SearchZoroByName(title)

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
