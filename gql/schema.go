package gql

import (
	"github.com/graphql-go/graphql"
)

var (
	fields = graphql.Fields{
		"version": &graphql.Field{
			Description: "Droll Version Information",
			Type:        VersionType,
			Resolve:     VersionQueryResolver,
		},
		"xkcd": &graphql.Field{
			Description: "List all XKCD Comics",
			Type:        graphql.NewList(XKCDType),
			Args:        Args,
			Resolve:     XKCDQueryResolver,
		},
		"phdcomic": &graphql.Field{
			Description: "List all PHD Comics",
			Type:        graphql.NewList(PHDComicType),
			Args:        Args,
			Resolve:     PHDComicQueryResolver,
		},
	}

	schemaConfig = graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "RootQuery",
			Fields:      fields,
			Description: "Root Query Type for GraphQL",
		}),
	}
)

// GetSchema : Returns global schema
func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(schemaConfig)
}
