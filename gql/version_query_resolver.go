package gql

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/model"
	"github.com/prabhuomkar/droll-api/utils"
)

// VersionQueryResolver : Resolver for query { version }
var VersionQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	return &model.Version{
		Timestamp: time.Now(),
		Name:      utils.APIName,
		Version:   utils.APIVersion,
		Comics:    utils.ComicsSupported,
	}, nil
}
