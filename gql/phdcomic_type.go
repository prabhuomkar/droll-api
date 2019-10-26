package gql

import (
	"github.com/graphql-go/graphql"
)

// PHDComicType : GraphQL type for phd comic
var PHDComicType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "phdcomic",
		Fields: graphql.Fields{
			"comicid": &graphql.Field{
				Type:        graphql.Int,
				Description: "ID for PHD comic",
			},
			"date": &graphql.Field{
				Type:        graphql.String,
				Description: "Date on which comic was released",
			},
			"link": &graphql.Field{
				Type:        graphql.String,
				Description: "Permalink of the comic",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of the comic",
			},
			"image": &graphql.Field{
				Type:        graphql.String,
				Description: "Image URL of the comic",
			},
		},
	},
)
