package routers

import (
	"net/http"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/providers"

	"github.com/gin-gonic/gin"
)

func SearchByName(c *gin.Context) {
	slug := c.Query("slug")
	title := c.Query("title")
	supportedServices := []string{"tmdb", "mal", "anilist"}
	provider := c.Param("provider")

	errorList := helpers.ValidateSearchRequest(c, supportedServices)
	if len(errorList) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorList})
		return
	}
	switch provider {
	case "tmdb":
		hasError, data := providers.SearchTmdbByName(title)
		if hasError {
			c.JSON(http.StatusNotFound, gin.H{"data": "fail"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"slug": slug, "data": data.Results})
	case "mal":
		hasError, data := providers.SearchMyanimeListByName(title)
		if hasError {
			c.JSON(http.StatusNotFound, gin.H{"data": "fail"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"slug": slug, "data": data.Result})
	case "anilist":
		hasError, data := providers.SearchAnilistByName(title)
		if hasError {
			c.JSON(http.StatusNotFound, gin.H{"data": "fail"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"slug": slug, "data": data.Result})
	}
}
