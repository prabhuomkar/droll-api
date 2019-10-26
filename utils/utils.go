package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BuildAPIURL : Creates the API URL based on Comic and URL parameters
func BuildAPIURL(comic string, num int) string {
	switch comic {
	case "xkcd":
		if num == -1 {
			return fmt.Sprintf("%s%s", XKCDBaseAPIURL, XKCDBaseAPIPath)
		}
		return fmt.Sprintf("%s%d%s", XKCDBaseAPIURL, num, XKCDBaseAPIPath)
	case "phdcomic":
		return fmt.Sprintf("%s%s%d", PHDComicBaseURL, PHDComicBaseRoute, num)
	}
	return ""
}

// FetchResponse : Gets response for HTTP GET request made on API URL
func FetchResponse(apiURL string) ([]byte, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("there is no such comic")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
