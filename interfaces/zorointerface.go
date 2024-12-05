package interfaces

type ZoroIndividualResult struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

type ZoroSearchResponse struct {
	Result []ZoroIndividualResult `json:"results"`
}
