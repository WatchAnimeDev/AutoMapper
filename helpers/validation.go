package helpers

import (
	"slices"

	"github.com/gin-gonic/gin"
)

func ValidateSearchRequest(c *gin.Context, supportedServices []string) []string {
	errorList := []string{}
	if c.Query("title") == "" {
		errorList = append(errorList, "Title is required")
	}
	if c.Param("provider") != "" {
		if !slices.Contains(supportedServices, c.Param("provider")) {
			errorList = append(errorList, "Provider "+c.Param("provider")+" is not supported")
		}
	} else {
		errorList = append(errorList, "Provider is required")
	}
	return errorList
}

func ValidateMetaImageRequest(c *gin.Context, supportedServices []string) []string {
	errorList := []string{}
	if c.Query("id") == "" {
		errorList = append(errorList, "ID is required")
	}
	if c.Param("provider") != "" {
		if !slices.Contains(supportedServices, c.Param("provider")) {
			errorList = append(errorList, "Provider "+c.Param("provider")+" is not supported")
		}
	} else {
		errorList = append(errorList, "Provider is required")
	}
	return errorList
}
