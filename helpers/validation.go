package helpers

import (
	"slices"
	"strings"

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

func ValidateSearchRequestAutoMap(c *gin.Context, requestedServices []string, supportedServices []string) []string {
	errorList := []string{}
	if c.Query("title") == "" {
		errorList = append(errorList, "Title is required")
	}
	if len(requestedServices) > 0 || (len(requestedServices) == 1 && requestedServices[0] == "") {
		missingServices := findMissingServices(requestedServices, supportedServices)
		if len(missingServices) > 0 {
			errorList = append(errorList, "Providers "+strings.Join(missingServices, ",")+" are not supported")
		}
	} else {
		errorList = append(errorList, "Provider is required")
	}
	return errorList
}

func findMissingServices(requestedServices, supportedServices []string) []string {
	missing := make([]string, 0)

	// Create a map of supported services for faster lookup
	supportedMap := make(map[string]bool)
	for _, service := range supportedServices {
		supportedMap[service] = true
	}

	// Check if each requested service is supported
	for _, service := range requestedServices {
		if !supportedMap[service] {
			missing = append(missing, service)
		}
	}

	return missing
}
