package structures

type Track struct {
	Album       Album             `json:"album"`
	Artists     []Artist          `json:"artists"`
	DurationMs  int64             `json:"duration_ms"`
	Explicit    bool              `json:"explicit"`
	ExternalIds map[string]string `json:"external_ids"`
	Href        string            `json:"href"`
	Id          string            `json:"id"`
	IsLocal     bool              `json:"is_local"`
	IsPlayable  bool              `json:"is_playable"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
}
