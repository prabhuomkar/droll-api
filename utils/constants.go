package utils

import (
	"github.com/prabhuomkar/droll-api/model"
)

var (
	// APIVersion ...
	APIVersion = "1.1.0"
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
		model.ComicInfo{
			Name:  "phdcomic",
			About: "\"Piled Higher and Deeper\" (PhD) is the comic strip about life (or the lack thereof) in academia.",
			Link:  "http://phdcomics.com",
			Logo:  "http://phdcomics.com/images/phd_logo.png",
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

	// PHDComicName ...
	PHDComicName = "phdcomic"
	// PHDComicBaseURL ...
	PHDComicBaseURL = "http://phdcomics.com"
	// PHDComicBaseRoute ...
	PHDComicBaseRoute = "/comics/archive_print.php?comicid="
	// PHDComicLink ...
	PHDComicLink = "http://phdcomics.com/comics/archive.php?comicid="
)
