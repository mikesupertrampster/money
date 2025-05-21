package types

import (
	"money/database"
	"time"
)

type Metrics struct {
	Symbol              string    `clover:"symbol"`
	Price               string    `clover:"price"`
	Change              string    `clover:"change"`
	MarketCap           string    `clover:"marketCap"`
	FiftyTwoWeekRange   string    `clover:"fiftyTwoWeekRange"`
	FiftyTwoWeekHigh    string    `clover:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow     string    `clover:"fiftyTwoWeekLow"`
	FiftyTwoWeekHighPct string    `clover:"fiftyTwoWeekHighPct"`
	FiftyTwoWeekLowPct  string    `clover:"fiftyTwoWeekLowPct"`
	Volume              string    `clover:"volume"`
	RSI14               string    `clover:"rSI14"`
	Image               string    `clover:"image"`
	TimeStamp           time.Time `clover:"timestamp"`
}

func (m *Metrics) Save() error {
	db, err := database.Load()
	if err != nil {
		return err
	}
	if err = db.Save(m.Symbol, m); err != nil {
		return err
	}
	return nil
}

func (m *Metrics) Get(db database.DB, symbol string) error {
	data, err := db.FindSymbol(symbol)
	if err != nil {
		return err
	}
	if err = data.Unmarshal(&m); err != nil {
		return err
	}
	return nil
}
