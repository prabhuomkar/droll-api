package gql

import (
	"github.com/graphql-go/graphql"
)

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
