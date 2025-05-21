package api

import (
	"money/intelligence"
	"money/types"
	"net/url"
	"os"
)

type TwelveData struct {
	rateLimit string
	url       *url.URL
}

func NewTwelveData() (*TwelveData, error) {
	u, err := url.Parse("https://api.twelvedata.com")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("apikey", os.Getenv("TWELVEDATA_KEY"))
	u.RawQuery = q.Encode()

	return &TwelveData{
		rateLimit: "8 API (800 a day) + 8 trial WS",
		url:       u,
	}, nil
}

func (t *TwelveData) Quote(symbol string) (types.Quote, error) {
	var quote types.Quote
	t.url.Path = "quote"
	q := t.url.Query()
	q.Set("symbol", symbol)
	t.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(t.url, &quote)
	if err != nil {
		return quote, err
	}
	return quote, nil
}

func (t *TwelveData) RealTimePrice(symbol string) (types.RealTimePrice, error) {
	var price types.RealTimePrice
	t.url.Path = "price"
	q := t.url.Query()
	q.Set("symbol", symbol)
	t.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(t.url, &price)
	if err != nil {
		return price, err
	}
	return price, nil
}
