package entity

type URLResult struct {
	URL          string `json:"url"`
	Availability bool   `json:"availability"`
	Error        string `json:"error"`
}