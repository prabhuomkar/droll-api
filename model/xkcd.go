package model

// XKCD : Model for xkcd comic
type XKCD struct {
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Image      string `json:"img"`
	Alt        string `json:"alt"`
	Transcript string `json:"transcript"`
	News       string `json:"news"`
}
