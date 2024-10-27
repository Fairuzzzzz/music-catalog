package spotify

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
}

type SpotifyTrackObject struct {
	// album related fields
	AlbumType        string   `json:"albumType"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesUrl   []string `json:"albumImagesUrl"`
	AlbumName        string   `json:"albumName"`
	AlbumReleaseDate string   `json:"albumReleaseDate"`

	// artists related fields
	ArtistsName []string `json:"artists"`

	// track related fields
	Explicit bool   `json:"explicit"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsLiked  *bool  `json:"isLiked"`
}

type RecommendationResponse struct {
	Items []SpotifyTrackObject `json:"items"`
}
