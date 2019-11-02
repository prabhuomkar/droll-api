package gql

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
)

type feedComicTest struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var feedComicTests = []feedComicTest{}
var feedComicLengthTests = []feedComicTest{}

func init() {
	schema, _ := GetSchema()
	feedComicTests = []feedComicTest{
		{
			Query: `
				query {
					feed(limit: 1, name: xkcd) {
						id
						name
						title
						description
						published
						link
						imageURL
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"feed": []map[string]interface{}{
						{
							"id":          1,
							"name":        "xkcd",
							"published":   "1/1/2006",
							"title":       "Barrel - Part 1",
							"imageURL":    "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
							"link":        "https://xkcd.com/1",
							"description": "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?\n[[The barrel drifts into the distance. Nothing else can be seen.]]\n{{Alt: Don't we all.}}",
						},
					},
				},
			},
		},
	}

}
func TestFeedComicQueryResolver(t *testing.T) {
	for _, test := range feedComicTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testfeedComic(test, params, t)
	}
}

func testfeedComic(test feedComicTest, p graphql.Params, t *testing.T) {
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
