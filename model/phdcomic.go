package model

// PHDComic : Model for phd comic
type PHDComic struct {
	Date    string `json:"date"`
	ComicID int    `json:"comicid"`
	Link    string `json:"link"`
	Title   string `json:"title"`
	Image   string `json:"img"`
}
