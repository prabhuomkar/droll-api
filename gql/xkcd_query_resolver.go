package gql

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/model"
	"github.com/prabhuomkar/droll-api/utils"
)

// XKCDQueryResolver : Resolver for query { xkcd }
var XKCDQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	limit, ok := p.Args["limit"].(int)
	if !ok || limit > utils.Limit || limit < 1 {
		limit = utils.Limit
	}
	offset, ok := p.Args["offset"].(int)
	if !ok || offset < 1 {
		offset = 0
	}

	return fetchResolveXKCDComics(limit, offset, false)
}

func fetchResolveXKCDComics(limit, offset int, forFeed bool) (interface{}, error) {
	// send parallel HTTP requests to xkcd API
	semaphoreChan := make(chan struct{}, limit)
	resultsChan := make(chan *model.XKCD)
	feedResultsChan := make(chan *model.Comic)
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
		close(feedResultsChan)
	}()

	for num := offset + 1; num <= offset+limit; num++ {
		go func(num int) {
			semaphoreChan <- struct{}{}
			feedComic, comic, _ := fetchXKCDComic(num)
			resultsChan <- comic
			feedResultsChan <- feedComic
			<-semaphoreChan
		}(num)
	}

	// create slice for comics from xkcd API responses
	var comics []*model.XKCD
	var feedComics []*model.Comic
	for {
		comic := <-resultsChan
		feedComic := <-feedResultsChan
		comics = append(comics, comic)
		feedComics = append(feedComics, feedComic)
		if len(comics) == limit && len(feedComics) == limit {
			break
		}
	}

	// sort for ordering of comics
	sort.Slice(comics, func(i, j int) bool {
		return comics[i].Num < comics[j].Num
	})
	sort.Slice(feedComics, func(i, j int) bool {
		return feedComics[i].ID < feedComics[j].ID
	})

	if forFeed {
		return &feedComics, nil
	}
	return &comics, nil
}

// fetches a single comic for given comic number
func fetchXKCDComic(num int) (*model.Comic, *model.XKCD, error) {
	apiURL := utils.BuildAPIURL(utils.XKCDComicName, num)
	body, err := utils.FetchResponse(apiURL)
	if err != nil {
		return nil, nil, err
	}
	var xkcd model.XKCD
	err = json.Unmarshal(body, &xkcd)
	if err != nil {
		return nil, nil, err
	}

	comicLink := xkcd.Link
	if comicLink == "" {
		comicLink = fmt.Sprintf("%s%d", utils.XKCDBaseAPIURL, xkcd.Num)
	}
	return &model.Comic{
		ID:          xkcd.Num,
		Title:       xkcd.Title,
		Description: xkcd.Transcript,
		ImageURL:    xkcd.Image,
		Link:        comicLink,
		Name:        utils.XKCDComicName,
		Published:   fmt.Sprintf("%s/%s/%s", xkcd.Month, xkcd.Day, xkcd.Year),
	}, &xkcd, err
}
