package resolvers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prabhuomkar/droll-api/schema"
)

// BuildAPIURL ...
func BuildAPIURL(comic string, num int) string {
	switch comic {
	case "xkcd":
		if num == -1 {
			return fmt.Sprintf("%s%s", schema.XKCDBaseAPIURL, schema.XKCDBaseAPIPath)
		}
		return fmt.Sprintf("%s%d%s", schema.XKCDBaseAPIURL, num, schema.XKCDBaseAPIPath)
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
