package interfaces

type KitsuAlgoliaKeyResponse struct {
	Users      kitsuApiKeyResource `json:"users"`
	Media      kitsuApiKeyResource `json:"media"`
	Groups     kitsuApiKeyResource `json:"groups"`
	Characters kitsuApiKeyResource `json:"characters"`
}

type kitsuApiKeyResource struct {
	Key   string `json:"key"`
	Index string `json:"index"`
}

type KitsuIndividualResult struct {
	Titles         map[string]string `json:"titles"`
	CanonicalTitle string            `json:"canonicalTitle"`
	Subtype        string            `json:"subtype"`
	Slug           string            `json:"slug"`
	PosterImage    struct {
		Tiny     string `json:"tiny"`
		Large    string `json:"large"`
		Small    string `json:"small"`
		Medium   string `json:"medium"`
		Original string `json:"original"`
		Meta     struct {
			Dimensions struct {
				Tiny   kistuDimension `json:"tiny"`
				Large  kistuDimension `json:"large"`
				Small  kistuDimension `json:"small"`
				Medium kistuDimension `json:"medium"`
			} `json:"dimensions"`
		} `json:"meta"`
	} `json:"posterImage"`
	Kind     string `json:"kind"`
	ID       int    `json:"id"`
	ObjectID string `json:"objectID"`
}

type kistuDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type KitsuSearchResponse struct {
	Result      []KitsuIndividualResult `json:"hits"`
	NbHits      int                     `json:"nbHits"`
	Page        int                     `json:"page"`
	NbPages     int                     `json:"nbPages"`
	HitsPerPage int                     `json:"hitsPerPage"`
}
