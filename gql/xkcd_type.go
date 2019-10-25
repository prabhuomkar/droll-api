package gql

import (
	"github.com/graphql-go/graphql"
)

// XKCDType : GraphQL type for xkcd comic
var XKCDType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "xkcd",
		Fields: graphql.Fields{
			"num": &graphql.Field{
				Type:        graphql.Int,
				Description: "ID for XKCD comic",
			},
			"day": &graphql.Field{
				Type:        graphql.String,
				Description: "Date on which comic was released",
			},
			"month": &graphql.Field{
				Type:        graphql.String,
				Description: "Month on which comic was released",
			},
			"year": &graphql.Field{
				Type:        graphql.String,
				Description: "Year when comic was released",
			},
			"link": &graphql.Field{
				Type:        graphql.String,
				Description: "Permalink of the comic",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of the comic",
			},
			"safeTitle": &graphql.Field{
				Type:        graphql.String,
				Description: "Safe title of the comic",
			},
			"image": &graphql.Field{
				Type:        graphql.String,
				Description: "Image URL of the comic",
			},
			"alt": &graphql.Field{
				Type:        graphql.String,
				Description: "Alternate text for the image of comic",
			},
			"transcript": &graphql.Field{
				Type:        graphql.String,
				Description: "Description of the comic",
			},
			"news": &graphql.Field{
				Type:        graphql.String,
				Description: "Any associated news relative to comic",
			},
		},
	},
)
