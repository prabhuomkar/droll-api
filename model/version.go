package model

import "time"

// Version : Model for version
type Version struct {
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Comics    []string  `json:"comics"`
}
