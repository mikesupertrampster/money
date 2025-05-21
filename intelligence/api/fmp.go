package api

import (
	"money/intelligence"
	"money/types"
	"net/url"
	"os"
)

type FMP struct {
	rateLimit string
	url       *url.URL
}

func NewFMP() (*FMP, error) {
	u, err := url.Parse("https://financialmodelingprep.com")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("apikey", os.Getenv("FMP_KEY"))
	u.RawQuery = q.Encode()

	return &FMP{
		rateLimit: "250/d",
		url:       u,
	}, nil
}

func (f *FMP) Ratings(symbol string) ([]types.Rating, error) {
	var ratings []types.Rating
	f.url.Path = "stable/ratings-snapshot"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &ratings)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (f *FMP) GradesHistorical(symbol string) ([]types.GradesHistorical, error) {
	var grades []types.GradesHistorical
	f.url.Path = "stable/grades-historical"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &grades)
	if err != nil {
		return nil, err
	}
	return grades, nil
}

func (f *FMP) GradeSummery(symbol string) ([]types.GradeSummery, error) {
	var grades []types.GradeSummery
	f.url.Path = "stable/grades-consensus"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &grades)
	if err != nil {
		return nil, err
	}
	return grades, nil
}

func (f *FMP) EarningsReport(symbol string) ([]types.EarningsReport, error) {
	var earningsReports []types.EarningsReport
	f.url.Path = "stable/earnings"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &earningsReports)
	if err != nil {
		return nil, err
	}
	return earningsReports, nil
}

func (f *FMP) PriceEndOfDay(symbol string) ([]types.PriceEndOfDay, error) {
	var prices []types.PriceEndOfDay
	f.url.Path = "stable/historical-prices-eod"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &prices)
	if err != nil {
		return nil, err
	}
	return prices, nil
}

func (f *FMP) SharesFloat(symbol string) ([]types.SharesFloat, error) {
	var floats []types.SharesFloat
	f.url.Path = "stable/shares-float"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &floats)
	if err != nil {
		return nil, err
	}
	return floats, nil
}

func (f *FMP) FinancialScores(symbol string) ([]types.FinancialScores, error) {
	var scores []types.FinancialScores
	f.url.Path = "stable/financial-scores"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &scores)
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func (f *FMP) PriceChanges(symbol string) ([]types.PriceChanges, error) {
	var changes []types.PriceChanges
	f.url.Path = "stable/stock-price-change"
	q := f.url.Query()
	q.Set("symbol", symbol)
	f.url.RawQuery = q.Encode()
	err := intelligence.HttpGet(f.url, &changes)
	if err != nil {
		return nil, err
	}
	return changes, nil
}
