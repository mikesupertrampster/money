package api

import (
	"context"
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"money/intelligence"
	"money/types"
	"os"
	"time"
)

type FinnHub struct {
	rateLimit string
	client    *finnhub.DefaultApiService
}

func NewFinHub() *FinnHub {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("FINNHUB_KEY"))
	return &FinnHub{
		client:    finnhub.NewAPIClient(cfg).DefaultApi,
		rateLimit: "30 API calls/ second limit",
	}
}

func (f *FinnHub) CompanyNews(symbol string) ([]finnhub.CompanyNews, error) {
	currentTime := time.Now().Local()
	to := currentTime.Format("2006-01-02")
	from := currentTime.Add(-7 * 24 * time.Hour).Format("2006-01-02")

	var c []finnhub.CompanyNews
	res, _, err := f.client.CompanyNews(context.Background()).Symbol(symbol).From(from).To(to).Execute()
	err = intelligence.MapToStruct(res, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (f *FinnHub) BasicMetrics(symbol string) (types.FinnhubMetric, error) {
	var metric types.FinnhubMetric
	res, _, err := f.client.CompanyBasicFinancials(context.Background()).Symbol(symbol).Metric("all").Execute()
	err = intelligence.MapToStruct(res.Metric, &metric)
	if err != nil {
		return metric, err
	}
	metric.Symbol = symbol
	return metric, nil
}

func (f *FinnHub) EarningsSurprises(symbol string) ([]finnhub.EarningResult, error) {
	var e []finnhub.EarningResult
	res, _, err := f.client.CompanyEarnings(context.Background()).Symbol(symbol).Execute()
	err = intelligence.MapToStruct(res, &e)
	if err != nil {
		return e, err
	}
	return e, nil
}

func (f *FinnHub) RecommendationTrends(symbol string) ([]finnhub.RecommendationTrend, error) {
	var r []finnhub.RecommendationTrend
	res, _, err := f.client.RecommendationTrends(context.Background()).Symbol(symbol).Execute()
	err = intelligence.MapToStruct(res, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
