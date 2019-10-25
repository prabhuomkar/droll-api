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
				Type:        graphql.DateTime,
				Description: "Current timestamp for version",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "GraphQL Service Name",
			},
			"version": &graphql.Field{
				Type:        graphql.String,
				Description: "GraphQL Service Version",
			},
			"comics": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "List of comics supported by current version",
			},
		},
	},
)
