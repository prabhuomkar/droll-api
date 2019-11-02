package gql

import (
	"github.com/graphql-go/graphql"
)

// ComicType : GraphQL type for comic in feed
var ComicType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "feed",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "ID of the comic",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the comic",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of the comic",
			},
			"description": &graphql.Field{
				Type:        graphql.String,
				Description: "Description of the comic",
			},
			"link": &graphql.Field{
				Type:        graphql.String,
				Description: "Website link of the comic",
			},
			"published": &graphql.Field{
				Type:        graphql.String,
				Description: "Date on which comic was published",
			},
			"imageURL": &graphql.Field{
				Type:        graphql.String,
				Description: "Link to image of the comic",
			},
		},
	},
)
