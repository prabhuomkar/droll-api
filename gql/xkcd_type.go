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
