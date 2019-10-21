package schema

import "github.com/graphql-go/graphql"

var (
	// CurrentVersion ...
	CurrentVersion = "1.0.0"
	// APIName ...
	APIName = "Droll GraphQL API"
	// ComicsSupported ...
	ComicsSupported = []string{XKCDComicName}
	// Limit ...
	Limit = 10

	// XKCDBaseAPIURL ...
	XKCDBaseAPIURL = "https://xkcd.com/"
	// XKCDBaseAPIPath ...
	XKCDBaseAPIPath = "/info.0.json"
	// XKCDStartIndex ...
	XKCDStartIndex = 1

	// Args : Pagination Arguments
	Args = graphql.FieldConfigArgument{
		"limit": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"offset": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
)
