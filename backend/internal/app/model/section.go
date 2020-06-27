package model

// Section
type Section struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	UUID   string   `json:"uuid"`
	Themes []*Theme `json:"themes,omitempty"`
}
