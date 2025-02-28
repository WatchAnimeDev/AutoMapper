package interfaces

type AnizoneIndividualResult struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

type AnizoneSearchResponse struct {
	Result []AnizoneIndividualResult `json:"results"`
}
