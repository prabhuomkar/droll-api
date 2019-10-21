package schema

import (
	"github.com/graphql-go/graphql"
)

// XKCDComicName ...
var XKCDComicName = "xkcd"

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

// XKCDType : GraphQL type for xkcd comic
var XKCDType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "xkcd",
		Fields: graphql.Fields{
			"num": &graphql.Field{
				Type: graphql.Int,
			},
			"day": &graphql.Field{
				Type: graphql.String,
			},
			"month": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.String,
			},
			"link": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"safeTitle": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"alt": &graphql.Field{
				Type: graphql.String,
			},
			"transcript": &graphql.Field{
				Type: graphql.String,
			},
			"news": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
