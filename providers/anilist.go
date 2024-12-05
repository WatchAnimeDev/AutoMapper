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

func SearchAnilistByName(title string) (bool, interfaces.AniListSearchResponse) {
	var responseObject interfaces.AniListSearchResponseWrapper
	hasError := false

	requestBody, err := json.Marshal(map[string]string{
		"query": fmt.Sprintf(`{
			Page(page: 1, perPage: 50) {
				pageInfo {
				total
				perPage
				currentPage
				lastPage
				hasNextPage
				}
				media(search: "%s", sort: SEARCH_MATCH, type: ANIME) {
				id
				idMal
				title {
					romaji
					english
					native
					userPreferred
				}
				type
				format
				status
				description
				startDate {
					year
					month
					day
				}
				endDate {
					year
					month
					day
				}
				season
				seasonYear
				seasonInt
				episodes
				duration
				chapters
				volumes
				countryOfOrigin
				isLicensed
				source
				hashtag
				trailer {
					id
				}
				updatedAt
				coverImage {
					extraLarge
					large
					medium
					color
				}
				bannerImage
				genres
				synonyms
				averageScore
				meanScore
				popularity
				isLocked
				trending
				favourites
				tags {
					id
				}
				relations {
					edges {
					id
					}
				}
				characters {
					edges {
					id
					}
				}
				staff {
					edges {
					id
					}
				}
				studios {
					edges {
					id
					}
				}
				isAdult
				nextAiringEpisode {
					id
				}
				airingSchedule {
					edges {
					id
					}
				}
				trends {
					edges {
					node {
						averageScore
						popularity
						inProgress
						episode
					}
					}
				}
				externalLinks {
					id
				}
				streamingEpisodes {
					title
					thumbnail
					url
					site
				}
				rankings {
					id
				}
				mediaListEntry {
					id
				}
				reviews {
					edges {
					node {
						id
					}
					}
				}
				recommendations {
					edges {
					node {
						id
					}
					}
				}
				siteUrl
				}
			}
	}`, title),
	})
	if err != nil {
		return true, interfaces.AniListSearchResponse{}
	}

	finalUrl := "https://graphql.anilist.co/"

	response, err := http.Post(finalUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return true, interfaces.AniListSearchResponse{}
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return true, interfaces.AniListSearchResponse{}
	}

	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return true, interfaces.AniListSearchResponse{}
	}

	return hasError, responseObject.Data.Page
}

func SearchAniListByNameAndReturnBestMatch(title string) (bool, interfaces.AniListIndividualResult) {
	hasError, responseObject := SearchAnilistByName(title)

	if hasError || len(responseObject.Result) == 0 {
		return true, interfaces.AniListIndividualResult{}
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.Title.UserPreferred, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	return false, responseObject.Result[currIndex]
}

func SearchAniListByNameAndReturnBestMatchAsync(title string, result chan<- map[string]any, errors chan<- map[string]bool) {
	hasError, responseObject := SearchAnilistByName(title)

	if hasError || len(responseObject.Result) == 0 {
		errors <- map[string]bool{"anilist": true}
		return // Exit early on error
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.Title.UserPreferred, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	// Send result
	result <- map[string]any{"anilist": responseObject.Result[currIndex]}
}
