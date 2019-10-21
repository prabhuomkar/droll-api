package resolvers

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/schema"
)

// VersionQueryResolver : Resolver for query { version }
var VersionQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	return &schema.Version{
		Timestamp: time.Now(),
		Name:      schema.APIName,
		Version:   schema.CurrentVersion,
		Comics:    schema.ComicsSupported,
	}, nil
}
