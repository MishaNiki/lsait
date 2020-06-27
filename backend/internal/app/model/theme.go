package model

type Theme struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	UUID     string     `json:"uuid"`
	Articles []*Article `json:"articles,omitempty"`
}
