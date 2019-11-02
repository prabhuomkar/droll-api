package model

type (
	// Comic ...
	Comic struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"imageURL"`
		Published   string `json:"published"`
		Link        string `json:"link"`
		Name        string `json:"name"`
	}

	// Comics ...
	Comics []Comic
)
