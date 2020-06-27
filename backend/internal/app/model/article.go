package model

type Article struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Text        string   `json:"text,omitempty"`
	Date        string   `json:"date,omitempty"`
	Auth        *Profile `json:"auth,omitempty"`
	Article     bool     `json:"article,omitempty"`
	Theme       int      `json:"-"`
}
