package providers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/interfaces"
)

func SearchMyanimeListByName(title string) (bool, interfaces.MyanimeListSearchResponse) {
	var responseObject interfaces.MyanimeListSearchResponse

	finalUrl := "https://api.jikan.moe/v4/anime?q=" + url.QueryEscape(title)

	response, err := http.Get(finalUrl)
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

func SearchMyanimeListByNameAndReturnBestMatch(title string) (bool, interfaces.MyanimeListIndividualResult) {
	hasError, responseObject := SearchMyanimeListByName(title)

	if hasError || len(responseObject.Result) == 0 {
		return true, interfaces.MyanimeListIndividualResult{}
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
