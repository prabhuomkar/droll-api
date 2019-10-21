package gql

import (
	"github.com/graphql-go/graphql"
)

var (
	fields = graphql.Fields{
		"version": &graphql.Field{
			Type:    VersionType,
			Resolve: VersionQueryResolver,
		},
		"xkcd": &graphql.Field{
			Type:    graphql.NewList(XKCDType),
			Args:    Args,
			Resolve: XKCDQueryResolver,
		},
	}

	schemaConfig = graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery", Fields: fields,
		}),
	}
)

// GetSchema : Returns global schema
func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(schemaConfig)
}
