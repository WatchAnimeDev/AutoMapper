package routers

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"watchanime/auto-mapper/config"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/providers"

	"github.com/gin-gonic/gin"
)

func Automap(c *gin.Context) {

	slug := c.Query("slug")
	title := c.Query("title")

	providerList := strings.Split(c.Query("provider"), "|")

	supportedServices := config.SupportedServices

	providerFuncs := map[string]func(string, chan<- map[string]any, chan<- map[string]bool){
		"tmdb":    providers.SearchTmdbByNameAndReturnBestMatchAsync,
		"mal":     providers.SearchMyanimeListByNameAndReturnBestMatchAsync,
		"anilist": providers.SearchAniListByNameAndReturnBestMatchAsync,
		"kitsu":   providers.SearchKitsuByNameAndReturnBestMatchAsync,
		"zoro":    providers.SearchZoroByNameAndReturnBestMatchAsync,
	}

	errorList := helpers.ValidateSearchRequestAutoMap(c, providerList, supportedServices)

	if len(errorList) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorList})
		return
	}

	result := make(chan map[string]any, len(providerList))
	errors := make(chan map[string]bool, len(providerList))

	var wg sync.WaitGroup

	wg.Add(len(providerList))

	for _, provider := range providerList {
		go func(provider string) {
			defer wg.Done()
			if fn, ok := providerFuncs[provider]; ok {
				fn(title, result, errors)
			}
		}(provider)
	}

	// Close the result channel after all goroutines finish
	go func() {
		wg.Wait()
		close(result)
		close(errors)
	}()

	var data = make(map[string]any)

	for res := range result {
		for key, value := range res {
			data[key] = value
		}
	}

	// We dont do anything with errors unless all are errored.
	// Respond based on errors
	if len(errors) == len(providerList) {
		c.JSON(http.StatusNotFound, gin.H{"data": "fail"})
	} else {
		fmt.Println(len(errors))
		c.JSON(http.StatusOK, gin.H{"slug": slug, "data": data})
	}
}
