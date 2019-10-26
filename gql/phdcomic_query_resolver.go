package gql

import (
	"fmt"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/graphql-go/graphql"
	"github.com/prabhuomkar/droll-api/model"
	"github.com/prabhuomkar/droll-api/utils"
)

// PHDComicQueryResolver : Resolver for query { phdcomic }
var PHDComicQueryResolver = func(p graphql.ResolveParams) (interface{}, error) {
	limit, ok := p.Args["limit"].(int)
	if !ok || limit > utils.Limit || limit < 1 {
		limit = utils.Limit
	}
	offset, ok := p.Args["offset"].(int)
	if !ok || offset < 1 {
		offset = 0
	}

	// send parallel HTTP requests to phd comic page
	semaphoreChan := make(chan struct{}, limit)
	resultsChan := make(chan *model.PHDComic)
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for num := offset + 1; num <= offset+limit; num++ {
		go func(num int) {
			semaphoreChan <- struct{}{}
			comic, _ := fetchPHDComic(num)
			resultsChan <- comic
			<-semaphoreChan
		}(num)
	}

	// create slice for comics from phd comic page responses
	var comics []*model.PHDComic
	for {
		comic := <-resultsChan
		comics = append(comics, comic)
		if len(comics) == limit {
			break
		}
	}

	// sort for ordering of comics
	sort.Slice(comics, func(i, j int) bool {
		return comics[i].ComicID < comics[j].ComicID
	})

	return &comics, nil
}

// fetches a single comic for given comic number
func fetchPHDComic(num int) (*model.PHDComic, error) {
	apiURL := utils.BuildAPIURL(utils.PHDComicName, num)
	body, err := utils.FetchResponse(apiURL)
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(string(body))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	splitContent := []string{}
	content := " "
	doc.Find("font").Each(func(index int, item *goquery.Selection) {
		if item.AttrOr("size", "") == "-2" {
			content = strings.TrimSpace(item.Text())
			splitContent = strings.Split(content, " ")
			content = strings.Join(splitContent[11:len(splitContent)-4], " ")
			content = strings.ReplaceAll(content, "\"", "")
		}
	})

	return &model.PHDComic{
		ComicID: num,
		Title:   content,
		Image:   doc.Find(`img`).AttrOr(`src`, ``),
		Link:    fmt.Sprintf("%s%d", utils.PHDComicLink, num),
		Date:    splitContent[len(splitContent)-1],
	}, err
}
