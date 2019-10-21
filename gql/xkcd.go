package gql

import (
	"encoding/json"
	"sort"

	"github.com/graphql-go/graphql"
)

// XKCDComicName ...
var XKCDComicName = "xkcd"

// XKCD : Model for xkcd comic
type XKCD struct {
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Image      string `json:"img"`
	Alt        string `json:"alt"`
	Transcript string `json:"transcript"`
	News       string `json:"news"`
}

// XKCDType : GraphQL type for xkcd comic
var XKCDType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "xkcd",
		Fields: graphql.Fields{
			"num": &graphql.Field{
				Type: graphql.Int,
			},
			"day": &graphql.Field{
				Type: graphql.String,
			},
			"month": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.String,
			},
			"link": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"safeTitle": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"alt": &graphql.Field{
				Type: graphql.String,
			},
			"transcript": &graphql.Field{
				Type: graphql.String,
			},
			"news": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// XKCDQueryResolver : Resolver for query { xkcd }
var XKCDQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	limit, ok := p.Args["limit"].(int)
	if !ok || limit > Limit {
		limit = Limit
	}
	offset, ok := p.Args["offset"].(int)
	if !ok || offset < 1 {
		offset = 0
	}

	// send parallel HTTP requests to xkcd API
	semaphoreChan := make(chan struct{}, limit)
	resultsChan := make(chan *XKCD)
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for num := offset + 1; num <= offset+limit; num++ {
		go func(num int) {
			semaphoreChan <- struct{}{}
			comic, _ := fetchXKCDComic(num)
			resultsChan <- comic
			<-semaphoreChan
		}(num)
	}

	// create slice for comics from xkcd API responses
	var comics []*XKCD
	for {
		comic := <-resultsChan
		comics = append(comics, comic)
		if len(comics) == limit {
			break
		}
	}

	// sort for ordering of comics
	sort.Slice(comics, func(i, j int) bool {
		return comics[i].Num < comics[j].Num
	})

	return &comics, nil
}

// fetches a single comic for given comic number
func fetchXKCDComic(num int) (*XKCD, error) {
	apiURL := BuildAPIURL(XKCDComicName, num)
	body, err := FetchResponse(apiURL)
	if err != nil {
		return nil, err
	}
	var xkcd XKCD
	err = json.Unmarshal(body, &xkcd)
	if err != nil {
		return nil, err
	}
	return &xkcd, err
}
