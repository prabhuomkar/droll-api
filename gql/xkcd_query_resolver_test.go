package gql

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
)

type xkcdTest struct {
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

var xkcdTests = []xkcdTest{}
var xkcdLengthTests = []xkcdTest{}

func init() {
	schema, _ := GetSchema()
	xkcdTests = []xkcdTest{
		{
			Query: `
				query {
					xkcd(limit: 1) {
						num
						day
						month
						year
						title
						safeTitle
						image
						link
						alt
						news
						transcript
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"xkcd": []map[string]interface{}{
						{
							"num":        1,
							"day":        "1",
							"month":      "1",
							"year":       "2006",
							"title":      "Barrel - Part 1",
							"safeTitle":  "Barrel - Part 1",
							"image":      "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
							"link":       "",
							"alt":        "Don't we all.",
							"news":       "",
							"transcript": "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?\n[[The barrel drifts into the distance. Nothing else can be seen.]]\n{{Alt: Don't we all.}}",
						},
					},
				},
			},
		}, {
			Query: `
				query {
					xkcd(limit: 1, offset: 999) {
						num
						day
						month
						year
						title
						safeTitle
						image
						link
						alt
						news
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"xkcd": []map[string]interface{}{
						{
							"month":     "1",
							"num":       1000,
							"link":      "https://xkcd.com/1000/large/",
							"year":      "2012",
							"news":      "",
							"safeTitle": "1000 Comics",
							"alt":       "Thank you for making me feel less alone.",
							"image":     "https://imgs.xkcd.com/comics/1000_comics.png",
							"title":     "1000 Comics",
							"day":       "6",
						},
					},
				},
			},
		},
	}
	xkcdLengthTests = []xkcdTest{
		{
			Query: `
				query {
					xkcd(limit: -1) {
						num
					}
				}
			`,
			Schema: schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"xkcd": []map[string]interface{}{
						{"num": 1},
						{"num": 2},
						{"num": 3},
						{"num": 4},
						{"num": 5},
						{"num": 6},
						{"num": 7},
						{"num": 8},
						{"num": 9},
						{"num": 10},
					},
				},
			},
		},
	}
}
func TestXKCDQueryResolver(t *testing.T) {
	for _, test := range xkcdTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testXKCD(test, params, t)
	}
}

func TestXKCDLength(t *testing.T) {
	for _, test := range xkcdLengthTests {
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
		}
		testXKCDLength(test, params, t)
	}
}

func testXKCD(test xkcdTest, p graphql.Params, t *testing.T) {
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

func testXKCDLength(test xkcdTest, p graphql.Params, t *testing.T) {
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

func TestFetchXKCDComic(t *testing.T) {
	comic, xkcdComic, err := fetchXKCDComic(0)
	if err == nil || xkcdComic != nil || comic != nil {
		t.Fatalf("expected xkcd: nil, got: %v", xkcdComic)
	}
	comic, xkcdComic, err = fetchXKCDComic(1)
	if err != nil || xkcdComic == nil || comic == nil {
		t.Fatalf("expected error: nil, got: %v", err)
	}
}
