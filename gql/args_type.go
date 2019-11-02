package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/utils"
)

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

// FeedArgs : Pagination Arguments
var FeedArgs = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type:        graphql.NewEnum(nameEnumConfig),
		Description: "Name of the comic",
	},
	"limit": &graphql.ArgumentConfig{
		Type:        graphql.Int,
		Description: "Number of comics to return in results",
	},
	"offset": &graphql.ArgumentConfig{
		Type:        graphql.Int,
		Description: "Number of comics to skip in results",
	},
}

// nameEnumConfig ...
var nameEnumConfig = graphql.EnumConfig{
	Name: "name",
	Values: graphql.EnumValueConfigMap{
		"xkcd":     &graphql.EnumValueConfig{Value: utils.XKCDComicName},
		"phdcomic": &graphql.EnumValueConfig{Value: utils.PHDComicName},
	},
}
