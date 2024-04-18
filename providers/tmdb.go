package providers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/interfaces"
)

func SearchTmdbByName(title string) (bool, interfaces.TmdbSearchResponse) {

	var responseObject interfaces.TmdbSearchResponse

	types := "tv"
	includeAdult := strconv.FormatBool(false)
	language := "en-US"
	page := strconv.Itoa(1)

	finalUrl := "https://api.themoviedb.org/3/search/" + types + "?query=" + url.QueryEscape(title) + "&include_adult=" + includeAdult + "&language=" + language + "&page=" + page + "&api_key=" + os.Getenv("TMDB_API_KEY")

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

func SearchTmdbByNameAndReturnBestMatch(title string) (bool, interfaces.TmdbSearchIndividualResult) {
	hasError, responseObject := SearchTmdbByName(title)

	if hasError || len(responseObject.Results) == 0 {
		return hasError, interfaces.TmdbSearchIndividualResult{}
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Results {
		currScore := helpers.MinDistance(result.Name, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	return false, responseObject.Results[currIndex]
}

func GetTmdbImageInfoByTmdbId(tmdbId string) (bool, interfaces.TmdbImageResponse) {
	var responseObject interfaces.TmdbImageResponse

	types := "tv"

	finalUrl := "https://api.themoviedb.org/3/" + types + "/" + tmdbId + "/images" + "?api_key=" + os.Getenv("TMDB_API_KEY")

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

func SearchTmdbByNameAndReturnBestMatchAsync(title string, result chan<- map[string]any, errors chan<- map[string]bool) {
	hasError, responseObject := SearchTmdbByName(title)

	if hasError || len(responseObject.Results) == 0 {
		errors <- map[string]bool{"tmdb": true}
		return // Exit early on error
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Results {
		currScore := helpers.MinDistance(result.Name, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	// Send result
	result <- map[string]any{"tmdb": responseObject.Results[currIndex]}
}
