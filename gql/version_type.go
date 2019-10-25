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
				Type:        graphql.NewList(ComicInfoType),
				Description: "List of comics supported by current version",
			},
		},
	},
)

// ComicInfoType : GraphQL type for basic information of comic
var ComicInfoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "comic",
		Fields: graphql.Fields{
			"about": &graphql.Field{
				Type:        graphql.String,
				Description: "About content of comic",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of comic",
			},
			"logo": &graphql.Field{
				Type:        graphql.String,
				Description: "Logo of comic",
			},
			"link": &graphql.Field{
				Type:        graphql.String,
				Description: "Link of comic website",
			},
		},
	},
)
