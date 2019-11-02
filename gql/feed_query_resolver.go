package gql

import (
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/utils"
)

// FeedQueryResolver : Resolver for query { feed }
var FeedQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	limit, ok := p.Args["limit"].(int)
	if !ok || limit > utils.Limit || limit < 1 {
		limit = utils.Limit
	}
	offset, ok := p.Args["offset"].(int)
	if !ok || offset < 1 {
		offset = 0
	}
	name, ok := p.Args["name"].(string)
	if !ok || strings.TrimSpace(name) == "" {
		name = ""
	}

	switch name {
	case utils.XKCDComicName:
		return fetchResolveXKCDComics(limit, offset, true)
	case utils.PHDComicName:
		return fetchResolvePHDComics(limit, offset, true)
	}

	return nil, nil
}
