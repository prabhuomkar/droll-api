package gql

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/utils"
)

type versionTest struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var versionTests = []versionTest{}

func init() {
	schema, _ := GetSchema()
	versionTests = []versionTest{
		{
			Query: `
				query {
					version {
						name
						version
						comics
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"version": map[string]interface{}{
						"name":    utils.APIName,
						"version": utils.APIVersion,
						"comics":  utils.ComicsSupported,
					},
				},
			},
		},
	}
}

func TestVersionQueryResolver(t *testing.T) {
	for _, test := range versionTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testVersionGraphQL(test, params, t)
	}
}

func testVersionGraphQL(test versionTest, p graphql.Params, t *testing.T) {
	result := graphql.Do(p)
	if len(result.Errors) > 0 {
		t.Fatalf("expected: no errors, got: %v", result.Errors)
	}
	resultJSON, _ := json.Marshal(result)
	expectedJSON, _ := json.Marshal(test.Expected)
	if !reflect.DeepEqual(resultJSON, expectedJSON) {
		t.Fatalf("expected: %v, got: %v", test.Expected, result)
	}
}