package gql

import "github.com/graphql-go/graphql"

// Args : Pagination Arguments
var Args = graphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"offset": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}
