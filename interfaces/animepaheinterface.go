package interfaces

type AnimepaheIndividualResult struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Type     string  `json:"type"`
	Episodes int     `json:"episodes"`
	Status   string  `json:"status"`
	Season   string  `json:"season"`
	Year     int     `json:"year"`
	Score    float64 `json:"score"`
	Poster   string  `json:"poster"`
	Session  string  `json:"session"`
}

type AnimepaheSearchResponse struct {
	Total        int                         `json:"total"`
	Per_page     int                         `json:"per_page"`
	Current_page int                         `json:"current_page"`
	Last_page    int                         `json:"last_page"`
	From         int                         `json:"from"`
	To           int                         `json:"to"`
	Result       []AnimepaheIndividualResult `json:"data"`
}
