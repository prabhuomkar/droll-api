package resolvers

import (
	"encoding/json"
	"sort"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/schema"
)

// XKCDQueryResolver : Resolver for query { xkcd }
var XKCDQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	limit, ok := p.Args["limit"].(int)
	if !ok || limit > schema.Limit {
		limit = schema.Limit
	}
	offset, ok := p.Args["offset"].(int)
	if !ok || offset < 1 {
		offset = 0
	}

	// send parallel HTTP requests to xkcd API
	semaphoreChan := make(chan struct{}, limit)
	resultsChan := make(chan *schema.XKCD)
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
	var comics []*schema.XKCD
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
func fetchXKCDComic(num int) (*schema.XKCD, error) {
	apiURL := BuildAPIURL(schema.XKCDComicName, num)
	body, err := FetchResponse(apiURL)
	if err != nil {
		return nil, err
	}
	var xkcd schema.XKCD
	err = json.Unmarshal(body, &xkcd)
	if err != nil {
		return nil, err
	}
	return &xkcd, err
}
