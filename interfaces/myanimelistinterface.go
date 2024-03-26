package interfaces

import (
	"time"
)

// Aired struct represents the aired information.
type Aired struct {
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	String string    `json:"string"`
}

// ImageUrls struct represents URLs for different image types.
type ImageUrls struct {
	ImageUrl      string `json:"image_url"`
	SmallImageUrl string `json:"small_image_url"`
	LargeImageUrl string `json:"large_image_url"`
}

// Trailer struct represents information about the anime trailer.
type Trailer struct {
	YouTubeID string    `json:"youtube_id"`
	URL       string    `json:"url"`
	EmbedURL  string    `json:"embed_url"`
	Images    ImageUrls `json:"images"`
}

// Title struct represents information about the anime title.
type Title struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

// Producer struct represents information about a producer.
type GenericEntity struct {
	MALID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

// Brodcast struct represents information about a broadcast.
type Brodcast struct {
	Day      string `json:"day"`
	Time     string `json:"time"`
	Timezone string `json:"timezone"`
	String   string `json:"string"`
}

// Anime struct represents the main structure of the API response.
type MyanimeListIndividualResult struct {
	MalId         int                  `json:"mal_id"`
	URL           string               `json:"url"`
	Images        map[string]ImageUrls `json:"images"`
	Trailer       Trailer              `json:"trailer"`
	Approved      bool                 `json:"approved"`
	Titles        []Title              `json:"titles"`
	Title         string               `json:"title"`
	TitleEnglish  string               `json:"title_english"`
	TitleJapanese string               `json:"title_japanese"`
	TitleSynonyms []string             `json:"title_synonyms"`
	Type          string               `json:"type"`
	Source        string               `json:"source"`
	Episodes      int                  `json:"episodes"`
	Status        string               `json:"status"`
	Airing        bool                 `json:"airing"`
	Aired         Aired                `json:"aired"`
	Duration      string               `json:"duration"`
	Rating        string               `json:"rating"`
	Score         float64              `json:"score"`
	ScoredBy      int                  `json:"scored_by"`
	Rank          int                  `json:"rank"`
	Popularity    int                  `json:"popularity"`
	Members       int                  `json:"members"`
	Favorites     int                  `json:"favorites"`
	Synopsis      string               `json:"synopsis"`
	Background    string               `json:"background"`
	Season        string               `json:"season"`
	Year          int                  `json:"year"`
	Broadcast     Brodcast             `json:"broadcast"`
	Producers     []GenericEntity      `json:"producers"`
	Licensors     []GenericEntity      `json:"licensors"`
	Studios       []GenericEntity      `json:"studios"`
	Genres        []GenericEntity      `json:"genres"`
	Themes        []GenericEntity      `json:"themes"`
	Demographics  []GenericEntity      `json:"demographics"`
}

type Items struct {
	Count   int `json:"count"`
	Total   int `json:"total"`
	PerPage int `json:"per_page"`
}

type MyanimeListPagination struct {
	LastVisiblePage int   `json:"last_visible_page"`
	HasNextPage     bool  `json:"has_next_page"`
	CurrentPage     int   `json:"current_page"`
	Items           Items `json:"items_per_page"`
}

type MyanimeListSearchResponse struct {
	Pagination MyanimeListPagination         `json:"pagination"`
	Result     []MyanimeListIndividualResult `json:"data"`
}
