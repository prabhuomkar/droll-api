package utils

import (
	"github.com/prabhuomkar/droll-api/model"
)

var (
	// APIVersion ...
	APIVersion = "1.0.1"
	// APIName ...
	APIName = "Droll GraphQL API"
	// ComicsSupported ...
	// TODO: Update this via JSON file
	ComicsSupported = []model.ComicInfo{
		model.ComicInfo{
			Name:  "xkcd",
			About: "A webcomic of romance, sarcasm, math, and language",
			Link:  "https://xkcd.com/",
			Logo:  "https://xkcd.com/s/0b7742.png",
		},
	}
	// Limit ...
	Limit = 10

	// XKCDComicName ...
	XKCDComicName = "xkcd"
	// XKCDBaseAPIURL ...
	XKCDBaseAPIURL = "https://xkcd.com/"
	// XKCDBaseAPIPath ...
	XKCDBaseAPIPath = "/info.0.json"
	// XKCDStartIndex ...
	XKCDStartIndex = 1
)
