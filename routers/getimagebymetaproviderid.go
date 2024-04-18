package routers

import (
	"net/http"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/providers"

	"github.com/gin-gonic/gin"
)

func GetImageByMetaProviderId(c *gin.Context) {
	metaProviderId := c.Query("id")
	supportedServices := []string{"tmdb", "mal", "anilist"}
	provider := c.Param("provider")

	errorList := helpers.ValidateMetaImageRequest(c, supportedServices)
	if len(errorList) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorList})
		return
	}
	switch provider {
	case "tmdb":
		hasError, data := providers.GetTmdbImageInfoByTmdbId(metaProviderId)
		if hasError {
			c.JSON(http.StatusNotFound, gin.H{"data": "fail"})
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	case "mal":
		// TODO
	case "anilist":
		// TODO
	}
}
