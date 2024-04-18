package main

import (
	"log"
	"os"
	"watchanime/auto-mapper/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	router := gin.Default()
	// router.Use(middlewares.Logger())
	router.GET("meta/:provider/search", routers.SearchByName)
	router.GET("meta/:provider/searchbestmatch", routers.SearchByNameAndReturnBestMatch)
	router.GET("meta/:provider/images", routers.GetImageByMetaProviderId)
	router.GET("meta/automap", routers.Automap)
	router.Run()
}
