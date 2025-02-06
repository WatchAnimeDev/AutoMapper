package providers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"watchanime/auto-mapper/helpers"
	"watchanime/auto-mapper/interfaces"

	"github.com/redis/go-redis/v9"
)

// Initialize Redis client (you might want to move this to a separate file)
var redisClient *redis.Client

// Fetch animepaheCookie from Redis
func getAnimepaheCookie() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cookie, err := redisClient.Get(ctx, "animepaheCookie").Result()
	if err == redis.Nil {
		log.Println("animepaheCookie not found in Redis")
		return "", nil // No cookie yet, continue without it
	} else if err != nil {
		log.Println("Error fetching animepaheCookie from Redis:", err)
		return "", err
	}

	return cookie, nil
}

// Store animepaheCookie in Redis
func updateAnimepaheCookie(newCookies string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch old cookies
	oldCookies, _ := getAnimepaheCookie()

	// Merge old and new cookies
	mergedCookies := mergeCookies(oldCookies, newCookies)

	// Store merged cookies
	err := redisClient.Set(ctx, "animepaheCookie", mergedCookies, 0).Err()
	if err != nil {
		log.Println("Error storing animepaheCookie in Redis:", err)
		return err
	}

	log.Println("Updated animepaheCookie in Redis:", mergedCookies)
	return nil
}

// Extract cookie string from response headers
func extractCookies(response *http.Response) string {
	var cookieList []string
	for _, cookie := range response.Cookies() {
		cookieList = append(cookieList, cookie.Name+"="+cookie.Value)
	}
	return strings.Join(cookieList, "; ")
}

// Merge old and new cookies, preferring new values
func mergeCookies(oldCookies, newCookies string) string {
	cookieMap := make(map[string]string)

	// Add old cookies to the map
	for _, cookie := range strings.Split(oldCookies, "; ") {
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) == 2 {
			cookieMap[parts[0]] = parts[1]
		}
	}

	// Overwrite with new cookies
	for _, cookie := range strings.Split(newCookies, "; ") {
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) == 2 {
			cookieMap[parts[0]] = parts[1]
		}
	}

	// Convert map back to cookie string
	var mergedCookies []string
	for key, value := range cookieMap {
		mergedCookies = append(mergedCookies, key+"="+value)
	}

	return strings.Join(mergedCookies, "; ")
}

// Search anime on Animepahe and update the cookie in Redis
func SearchAnimepaheByName(title string) (bool, interfaces.AnimepaheSearchResponse) {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"), // Change if Redis is running on a different host/port
		Password: os.Getenv("REDIS_PASS"), // Add password if Redis requires authentication
		DB:       0,                       // Use default DB
	})

	var responseObject interfaces.AnimepaheSearchResponse

	// Fetch the stored cookie
	cookieValue, _ := getAnimepaheCookie()

	// Create the HTTP request
	request, err := http.NewRequest("GET", "https://animepahe.ru/api?m=search&q="+url.QueryEscape(title), nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return true, interfaces.AnimepaheSearchResponse{}
	}

	// Add stored cookie if available
	if cookieValue != "" {
		request.Header.Set("Cookie", cookieValue)
	}

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error fetching URL:", err)
		return true, interfaces.AnimepaheSearchResponse{}
	}
	defer response.Body.Close()

	// Extract new cookies from response and merge with existing ones
	newCookieValue := extractCookies(response)
	if newCookieValue != "" {
		_ = updateAnimepaheCookie(newCookieValue)
	}

	// Read response body
	responseData, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK || err != nil {
		log.Printf("Status code error: %d %s", response.StatusCode, response.Status)
		return true, interfaces.AnimepaheSearchResponse{}
	}

	// Parse JSON response
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return true, interfaces.AnimepaheSearchResponse{}
	}

	return false, responseObject
}

func SearchAnimepaheByNameAndReturnBestMatch(title string) (bool, interfaces.AnimepaheIndividualResult) {
	hasError, responseObject := SearchAnimepaheByName(title)

	if hasError || len(responseObject.Result) == 0 {
		return true, interfaces.AnimepaheIndividualResult{}
	}

	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.Title, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	return false, responseObject.Result[currIndex]
}

func SearchAnimepaheByNameAndReturnBestMatchAsync(title string, result chan<- map[string]any, errors chan<- map[string]bool) {
	hasError, responseObject := SearchAnimepaheByName(title)

	if hasError || len(responseObject.Result) == 0 {
		errors <- map[string]bool{"animepahe": true}
		return // Exit early on error
	}

	// Find best match
	scores := len(title)
	currIndex := 0
	for ind, result := range responseObject.Result {
		currScore := helpers.MinDistance(result.Title, title)
		if currScore < scores {
			scores = currScore
			currIndex = ind
		}
	}

	// Send result
	result <- map[string]any{"animepahe": responseObject.Result[currIndex]}
}
