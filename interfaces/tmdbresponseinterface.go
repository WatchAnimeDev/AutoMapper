package interfaces

type TmdbSearchIndividualResult struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIds         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	FirstAirDate     string   `json:"first_air_date"`
	Name             string   `json:"name"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

type TmdbSearchResponse struct {
	Page         int                          `json:"page"`
	Results      []TmdbSearchIndividualResult `json:"results"`
	TotalPages   int                          `json:"total_pages"`
	TotalResults int                          `json:"total_results"`
}

type TmdbImageResponse struct {
	ID      int `json:"id"`
	Posters []struct {
		AspectRatio float64 `json:"aspect_ratio"`
		Height      int     `json:"height"`
		Iso6391     string  `json:"iso_639_1"`
		FilePath    string  `json:"file_path"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"posters"`
	Backdrops []struct {
		AspectRatio float64 `json:"aspect_ratio"`
		Height      int     `json:"height"`
		Iso6391     string  `json:"iso_639_1"`
		FilePath    string  `json:"file_path"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"backdrops"`
	Logos []struct {
		AspectRatio float64 `json:"aspect_ratio"`
		Height      int     `json:"height"`
		Iso6391     string  `json:"iso_639_1"`
		FilePath    string  `json:"file_path"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"logos"`
}
