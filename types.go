package bmm

import "time"

type Year struct {
	Year  uint32 `json:"year"`
	Count uint32 `json:"count"`
}

type Meta struct {
	ContainedTypes []string  `json:"contained_types"`
	IsVisible      bool      `json:"is_visible"`
	ModifiedAt     time.Time `json:"modified_at"`
	ModifiedBy     string    `json:"modified_by"`
}

type Item struct {
	Meta      Meta        `json:"_meta"`
	BmmID     interface{} `json:"bmm_id"`
	Cover     string      `json:"cover"`
	ID        int         `json:"id"`
	Languages []string    `json:"languages"`
	//ParentID    interface{} `json:"parent_id"`
	PublishedAt            time.Time `json:"published_at"`
	Tags                   []string  `json:"tags"`
	Language               string    `json:"language"`
	Title                  string    `json:"title"`
	Type                   string    `json:"type"`
	Tracks                 []Item    `json:"children"`
	TranscriptionLanguages []string  `json:"transcription_languages"`
	HasTranscription       bool      `json:"has_transcription"`
}
type Overview struct {
	Name      string   `json:"name"`
	Languages []string `json:"languages"`
}
