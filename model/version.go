package model

import "time"

type (
	// Version : Model for version
	Version struct {
		Timestamp time.Time   `json:"timestamp"`
		Name      string      `json:"name"`
		Version   string      `json:"version"`
		Comics    []ComicInfo `json:"comics"`
	}

	// ComicInfo : Model for basic information of comic
	ComicInfo struct {
		Name  string `json:"name"`
		About string `json:"about"`
		Link  string `json:"link"`
		Logo  string `json:"logo"`
	}
)
