package gql

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

var (
	// CurrentVersion ...
	CurrentVersion = "1.0.0"
	// Name ...
	Name = "Droll GraphQL API"
	// ComicsSupported ...
	ComicsSupported = []string{XKCDComicName}
	// Limit ...
	Limit = 10

	// XKCDBaseAPIURL ...
	XKCDBaseAPIURL = "https://xkcd.com/"
	// XKCDBaseAPIPath ...
	XKCDBaseAPIPath = "/info.0.json"
	// XKCDStartIndex ...
	XKCDStartIndex = 1

	// Args : Pagination Arguments
	Args = graphql.FieldConfigArgument{
		"limit": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"offset": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
)

// BuildAPIURL ...
func BuildAPIURL(comic string, num int) string {
	switch comic {
	case "xkcd":
		if num == -1 {
			return fmt.Sprintf("%s%s", XKCDBaseAPIURL, XKCDBaseAPIPath)
		}
		return fmt.Sprintf("%s%d%s", XKCDBaseAPIURL, num, XKCDBaseAPIPath)
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

// Version : Model for version
type Version struct {
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	Comics    []string  `json:"comics"`
}

// VersionType : GraphQL type for version
var VersionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "version",
		Fields: graphql.Fields{
			"timestamp": &graphql.Field{
				Type: graphql.DateTime,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"version": &graphql.Field{
				Type: graphql.String,
			},
			"comics": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

// VersionQueryResolver : Resolver for query { version }
var VersionQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	return &Version{
		Timestamp: time.Now(),
		Name:      Name,
		Version:   CurrentVersion,
		Comics:    ComicsSupported,
	}, nil
}
