package types

import (
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"money/database"
)

type CompanyNews struct {
	Symbol string `json:"symbol"`
	News   []finnhub.CompanyNews
}

type FinnhubMetric struct {
	TDAverageTradingVolume float64 `json:"10DayAverageTradingVolume"`
	FTWHigh                float64 `json:"52WeekHigh"`
	FTWLow                 float64 `json:"52WeekLow"`
	FTWkLowDate            string  `json:"52WeekLowDate"`
	FTWPriceReturnDaily    float64 `json:"52WeekPriceReturnDaily"`
	Beta                   float64 `json:"beta"`
	Symbol                 string  `json:"symbol"`
}

func (f *FinnhubMetric) Save() error {
	db, err := database.Load("finnhub-metrics")
	if err != nil {
		return err
	}
	if err = db.Save(f.Symbol, f); err != nil {
		return err
	}
	if err = db.Client.Close(); err != nil {
		return err
	}
	return nil
}

func (f *FinnhubMetric) Get(symbol string) error {
	db, err := database.Load("finnhub-metrics")
	if err != nil {
		return err
	}

	data, err := db.FindSymbol(symbol)
	if err != nil {
		return err
	}
	if err = data.Unmarshal(&f); err != nil {
		return err
	}
	if err = db.Client.Close(); err != nil {
		return err
	}

	return nil
}
