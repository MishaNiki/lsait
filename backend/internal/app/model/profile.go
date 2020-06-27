package model

type Profile struct {
	ID          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	Position    string `json:"position,omitempty"`
	Description string `json:"description,omitempty"`

	Articles []*Article `json:"articles,omitempty"`
	Drafts   []*Article `json:"drafts,omitempty"`
}
