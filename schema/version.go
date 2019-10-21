package schema

import (
	"time"

	"github.com/graphql-go/graphql"
)

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
