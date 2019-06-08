package structures

type AudioFeature struct {
	DurationMs       int     `json:"duration_ms"`
	Key              int     `json:"key"`
	Mode             int     `json:"mode"`
	TimeSignature    int     `json:"time_signature"`
	Acousticness     float64 `json:"acousticness"`
	Danceability     float64 `json:"danceability"`
	Energy           float64 `json:"energy"`
	Instrumentalness float64 `json:"instrumentalness"`
	Liveness         float64 `json:"liveness"`
	Loudness         float64 `json:"loudness"`
	Speechiness      float64 `json:"speechiness"`
	Valence          float64 `json:"valence"`
	Tempo            float64 `json:"tempo"`
	ID               string  `json:"id"`
	URI              string  `json:"uri"`
	TrackHref        string  `json:"track_href"`
	AnalysisURL      string  `json:"analysis_url"`
	Type             string  `json:"type"`
}
