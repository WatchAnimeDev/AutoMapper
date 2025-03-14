package routers

import (
	"net/http"
	"watchanime/auto-mapper/config"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/providers"

	"github.com/gin-gonic/gin"
)

func SearchByNameAndReturnBestMatch(c *gin.Context) {
	slug := c.Query("slug")
	title := c.Query("title")
	provider := c.Param("provider")

	supportedServices := config.SupportedServices

	errorList := helpers.ValidateSearchRequest(c, supportedServices)
	if len(errorList) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorList})
		return
	}

	var data any
	var hasError bool

	switch provider {
	case "tmdb":
		hasError, data = providers.SearchTmdbByNameAndReturnBestMatch(title)
	case "mal":
		hasError, data = providers.SearchMyanimeListByNameAndReturnBestMatch(title)
	case "anilist":
		hasError, data = providers.SearchAniListByNameAndReturnBestMatch(title)
	case "kitsu":
		hasError, data = providers.SearchKistuByNameAndReturnBestMatch(title)
	case "zoro":
		hasError, data = providers.SearchZoroByNameAndReturnBestMatch(title)
	case "animepahe":
		hasError, data = providers.SearchAnimepaheByNameAndReturnBestMatch(title)
	case "anizone":
		hasError, data = providers.SearchAnizoneByNameAndReturnBestMatch(title)
	}

	if hasError {
		c.JSON(http.StatusNotFound, gin.H{"data": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"slug": slug, "data": data})

}
