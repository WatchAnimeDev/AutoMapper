package interfaces

type AniListIndividualResult struct {
	ID    int `json:"id"`
	IDMal int `json:"idMal"`
	Title struct {
		Romaji        string `json:"romaji"`
		English       string `json:"english"`
		Native        string `json:"native"`
		UserPreferred string `json:"userPreferred"`
	} `json:"title"`
	Type        string `json:"type"`
	Format      string `json:"format"`
	Status      string `json:"status"`
	Description string `json:"description"`
	StartDate   struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"startDate"`
	EndDate struct {
		Year  interface{} `json:"year"`
		Month interface{} `json:"month"`
		Day   interface{} `json:"day"`
	} `json:"endDate"`
	Season          string      `json:"season"`
	SeasonYear      int         `json:"seasonYear"`
	SeasonInt       int         `json:"seasonInt"`
	Episodes        int         `json:"episodes"`
	Duration        int         `json:"duration"`
	Chapters        interface{} `json:"chapters"`
	Volumes         interface{} `json:"volumes"`
	CountryOfOrigin string      `json:"countryOfOrigin"`
	IsLicensed      bool        `json:"isLicensed"`
	Source          string      `json:"source"`
	Hashtag         string      `json:"hashtag"`
	Trailer         struct {
		ID string `json:"id"`
	} `json:"trailer"`
	UpdatedAt  int64 `json:"updatedAt"`
	CoverImage struct {
		ExtraLarge string `json:"extraLarge"`
		Large      string `json:"large"`
		Medium     string `json:"medium"`
		Color      string `json:"color"`
	} `json:"coverImage"`
	BannerImage  string   `json:"bannerImage"`
	Genres       []string `json:"genres"`
	Synonyms     []string `json:"synonyms"`
	AverageScore int      `json:"averageScore"`
	MeanScore    int      `json:"meanScore"`
	Popularity   int      `json:"popularity"`
	IsLocked     bool     `json:"isLocked"`
	Trending     int      `json:"trending"`
	Favourites   int      `json:"favourites"`
	Tags         []struct {
		ID int `json:"id"`
	} `json:"tags"`
	Relations struct {
		Edges []struct {
			ID int `json:"id"`
		} `json:"edges"`
	} `json:"relations"`
	Characters struct {
		Edges []struct {
			ID int `json:"id"`
		} `json:"edges"`
	} `json:"characters"`
	Staff struct {
		Edges []struct {
			ID int `json:"id"`
		} `json:"edges"`
	} `json:"staff"`
	Studios struct {
		Edges []struct {
			ID int `json:"id"`
		} `json:"edges"`
	} `json:"studios"`
	IsAdult           bool `json:"isAdult"`
	NextAiringEpisode struct {
		ID int `json:"id"`
	} `json:"nextAiringEpisode"`
	AiringSchedule struct {
		Edges []struct {
			ID interface{} `json:"id"`
		} `json:"edges"`
	} `json:"airingSchedule"`
	Trends struct {
		Edges []struct {
			Node struct {
				AverageScore interface{} `json:"averageScore"`
				Popularity   int         `json:"popularity"`
				InProgress   interface{} `json:"inProgress"`
				Episode      interface{} `json:"episode"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"trends"`
	ExternalLinks []struct {
		ID int `json:"id"`
	} `json:"externalLinks"`
	StreamingEpisodes []interface{} `json:"streamingEpisodes"`
	Rankings          []struct {
		ID int `json:"id"`
	} `json:"rankings"`
	MediaListEntry interface{} `json:"mediaListEntry"`
	Reviews        struct {
		Edges []interface{} `json:"edges"`
	} `json:"reviews"`
	Recommendations struct {
		Edges []struct {
			Node struct {
				ID int `json:"id"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"recommendations"`
	SiteURL string `json:"siteUrl"`
}

type AnilistPagination struct {
	CurrentPage int  `json:"currentPage"`
	HasNextPage bool `json:"hasNextPage"`
	LastPage    int  `json:"lastPage"`
	PerPage     int  `json:"perPage"`
	Total       int  `json:"total"`
}

type AniListSearchResponse struct {
	Pagination AnilistPagination         `json:"pageInfo"`
	Result     []AniListIndividualResult `json:"media"`
}

type AniListSearchResponseWrapper struct {
	Data struct {
		Page AniListSearchResponse `json:"Page"`
	}
}
