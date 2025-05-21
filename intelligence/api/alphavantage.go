package api

import (
	"money/intelligence"
	"money/types"
	"net/url"
	"os"
	"strings"
)

type AlphaVantage struct {
	rateLimit string
	url       *url.URL
}

func NewAlphaVantage() (*AlphaVantage, error) {
	u, err := url.Parse("https://www.alphavantage.co")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("apikey", os.Getenv("ALPHAVANTAGE_KEY"))
	u.RawQuery = q.Encode()

	return &AlphaVantage{
		rateLimit: "25 requests per day",
		url:       u,
	}, nil
}

func (a *AlphaVantage) NewsSentiments(symbols []string) (types.NewsSentiments, error) {
	var ns types.NewsSentiments
	a.url.Path = "query"
	q := a.url.Query()
	q.Set("function", "NEWS_SENTIMENT")
	q.Set("tickers", strings.Join(symbols, ","))
	a.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(a.url, &ns)
	if err != nil {
		return ns, err
	}
	return ns, nil
}

func (a *AlphaVantage) TopGainersLosers() (types.TopGainersLosers, error) {
	var tgl types.TopGainersLosers
	a.url.Path = "query"
	q := a.url.Query()
	q.Set("function", "TOP_GAINERS_LOSERS")
	a.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(a.url, &tgl)
	if err != nil {
		return tgl, err
	}
	return tgl, nil
}
