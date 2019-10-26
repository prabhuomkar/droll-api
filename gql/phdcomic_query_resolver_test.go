package gql

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
)

type phdComicTest struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var phdComicTests = []phdComicTest{}
var phdComicLengthTests = []phdComicTest{}

func init() {
	schema, _ := GetSchema()
	phdComicTests = []phdComicTest{
		{
			Query: `
				query {
					phdcomic(limit: 1) {
						comicid
						date
						title
						image
						link
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"phdcomic": []map[string]interface{}{
						{
							"comicid": 1,
							"date":    "10/27/1997",
							"title":   "Very Close - First Phd strip!",
							"image":   "http://www.phdcomics.com/comics/archive/phd1027.gif",
							"link":    "http://phdcomics.com/comics/archive.php?comicid=1",
						},
					},
				},
			},
		}, {
			Query: `
				query {
					phdcomic(limit: 1, offset: 999) {
						comicid
						date
						title
						image
						link
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"phdcomic": []map[string]interface{}{
						{
							"comicid": 1000,
							"date":    "4/7/2008",
							"title":   "And it only took 1000 strips",
							"image":   "http://www.phdcomics.com/comics/archive/phd040708s.gif",
							"link":    "http://phdcomics.com/comics/archive.php?comicid=1000",
						},
					},
				},
			},
		},
	}
	phdComicLengthTests = []phdComicTest{
		{
			Query: `
				query {
					phdcomic(limit: -1) {
						comicid
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"phdcomic": []map[string]interface{}{
						{"comicid": 1},
						{"comicid": 2},
						{"comicid": 3},
						{"comicid": 4},
						{"comicid": 5},
						{"comicid": 6},
						{"comicid": 7},
						{"comicid": 8},
						{"comicid": 9},
						{"comicid": 10},
					},
				},
			},
		},
	}
}
func TestPHDComicQueryResolver(t *testing.T) {
	for _, test := range phdComicTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testPHDComic(test, params, t)
	}
}

func TestPHDComicLength(t *testing.T) {
	for _, test := range phdComicLengthTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testPHDComicLength(test, params, t)
	}
}

func testPHDComic(test phdComicTest, p graphql.Params, t *testing.T) {
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

func testPHDComicLength(test phdComicTest, p graphql.Params, t *testing.T) {
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

func TestFetchPHDComic(t *testing.T) {
	phdComic, err := fetchPHDComic(1)
	if err != nil || phdComic == nil {
		t.Fatalf("expected error: nil, got: %v", err)
	}
}
