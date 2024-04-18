package routers

import (
	"net/http"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/providers"

	"github.com/gin-gonic/gin"
)

func SearchByNameAndReturnBestMatch(c *gin.Context) {
	slug := c.Query("slug")
	title := c.Query("title")
	provider := c.Param("provider")

	supportedServices := []string{"tmdb", "mal", "anilist", "all"}

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
	}

	if hasError {
		c.JSON(http.StatusNotFound, gin.H{"data": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"slug": slug, "data": data})

}
