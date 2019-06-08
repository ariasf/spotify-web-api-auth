package structures

type Album struct {
	Id        string   `json:"id"`
	AlbumType string   `json:"album_type"`
	Artists   []Artist `json:"artists"`
	Href      string   `json:"href"`
	Images    []Image  `json:"images"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
}
