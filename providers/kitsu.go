package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/interfaces"
)

func SearchKitsuByName(title string) (bool, interfaces.KitsuSearchResponse) {

	var kitsuAlgoliaKeyResponse interfaces.KitsuAlgoliaKeyResponse
	var responseObject interfaces.KitsuSearchResponse

	//Get Kitsu Algolia Key
	kitsuAlgoliaKeyUrl := "https://kitsu.io/api/edge/algolia-keys"
	response, err := http.Get(kitsuAlgoliaKeyUrl)
	if err != nil {
		return true, responseObject
	}
	kitsuAlgoliaKeyResponseData, err := io.ReadAll(response.Body)
	if err != nil {
		return true, responseObject
	}
	err = json.Unmarshal(kitsuAlgoliaKeyResponseData, &kitsuAlgoliaKeyResponse)
	if err != nil {
		return true, responseObject
	}

	//Get actual value
	kistMediaUrl := "https://awqo5j657s-1.algolianet.com/1/indexes/production_media/query?x-algolia-agent=Algolia%20for%20vanilla%20JavaScript%20(lite)%203.24.12&x-algolia-application-id=AWQO5J657S&x-algolia-api-key=" + kitsuAlgoliaKeyResponse.Media.Key
	requestBody, err := json.Marshal(map[string]string{
		"params": fmt.Sprintf("query=%s&attributesToRetrieve=[\"id\",\"slug\",\"kind\",\"canonicalTitle\",\"titles\",\"posterImage\",\"subtype\",\"posterImage\"]&hitsPerPage=4&queryLanguages=[\"en\",\"ja\"]&naturalLanguages=[\"en\",\"ja\"]&attributesToHighlight=[]&responseFields=[\"hits\",\"hitsPerPage\",\"nbHits\",\"nbPages\",\"offset\",\"page\"]&removeStopWords=false&removeWordsIfNoResults=allOptional", title),
	})
	if err != nil {
		return true, responseObject
	}
	response, err = http.Post(kistMediaUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return true, responseObject
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return true, responseObject
	}
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return true, responseObject
	}

	return false, responseObject
}

func SearchKistuByNameAndReturnBestMatch(title string) (bool, interfaces.KitsuIndividualResult) {
	hasError, responseObject := SearchKitsuByName(title)

	if hasError || len(responseObject.Result) == 0 {
		return true, interfaces.KitsuIndividualResult{}
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.CanonicalTitle, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	return false, responseObject.Result[currIndex]
}

func SearchKitsuByNameAndReturnBestMatchAsync(title string, result chan<- map[string]any, errors chan<- map[string]bool) {
	hasError, responseObject := SearchKitsuByName(title)

	if hasError || len(responseObject.Result) == 0 {
		errors <- map[string]bool{"kitsu": true}
		return // Exit early on error
	}

	// Find best match
	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.CanonicalTitle, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	// Send result
	result <- map[string]any{"kitsu": responseObject.Result[currIndex]}
}
