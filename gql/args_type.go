package gql

import "github.com/graphql-go/graphql"

// Args : Pagination Arguments
var Args = graphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{
		Type:        graphql.Int,
		Description: "Number of comics to return in results",
	},
	"offset": &graphql.ArgumentConfig{
		Type:        graphql.Int,
		Description: "Number of comics to skip in results",
	},
}
