package api

import (
	"fmt"
	"money/intelligence"
	"money/types"
	"net/url"
	"os"
)

type Tiingo struct {
	rateLimit string
	url       *url.URL
}

func NewTiingo() (*Tiingo, error) {
	u, err := url.Parse("https://api.tiingo.com")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("token", os.Getenv("TIINGO_KEY"))
	u.RawQuery = q.Encode()

	return &Tiingo{
		rateLimit: "Max Requests Per Hour 50, Max Requests Per Day 1000",
		url:       u,
	}, nil
}

func (t *Tiingo) TopOfBook(symbol string) ([]types.TopOfBook, error) {
	var totb []types.TopOfBook
	t.url.Path = fmt.Sprintf("iex/%s", symbol)
	err := intelligence.HttpGet(t.url, &totb)
	if err != nil {
		return nil, err
	}
	return totb, nil
}
