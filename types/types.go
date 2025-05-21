package types

import "time"

// Metric Finnhub
type Metric struct {
	TDAverageTradingVolume float64 `json:"10DayAverageTradingVolume"`
	FTWHigh                float64 `json:"52WeekHigh"`
	FTWLow                 float64 `json:"52WeekLow"`
	FTWkLowDate            string  `json:"52WeekLowDate"`
	FTWPriceReturnDaily    float64 `json:"52WeekPriceReturnDaily"`
	Beta                   float64 `json:"beta"`
}

// TopOfBook Polygon
type TopOfBook struct {
	Ticker            string    `json:"ticker"`
	Timestamp         time.Time `json:"timestamp"`
	QuoteTimestamp    time.Time `json:"quoteTimestamp"`
	LastSaleTimeStamp time.Time `json:"lastSaleTimeStamp"`
	Last              float64   `json:"last"`
	LastSize          int       `json:"lastSize"`
	TngoLast          float64   `json:"tngoLast"`
	PrevClose         float64   `json:"prevClose"`
	Open              float64   `json:"open"`
	High              float64   `json:"high"`
	Low               float64   `json:"low"`
	Mid               float64   `json:"mid"`
	Volume            int       `json:"volume"`
	BidSize           int       `json:"bidSize"`
	BidPrice          float64   `json:"bidPrice"`
	AskSize           int       `json:"askSize"`
	AskPrice          float64   `json:"askPrice"`
}

// NewsSentiments AlphaVantage
type NewsSentiments struct {
	Items                    string `json:"items"`
	SentimentScoreDefinition string `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string `json:"relevance_score_definition"`
	Feed                     []struct {
		Title                string   `json:"title"`
		Url                  string   `json:"url"`
		TimePublished        string   `json:"time_published"`
		Authors              []string `json:"authors"`
		Summary              string   `json:"summary"`
		BannerImage          string   `json:"banner_image"`
		Source               string   `json:"source"`
		CategoryWithinSource string   `json:"category_within_source"`
		SourceDomain         string   `json:"source_domain"`
		Topics               []struct {
			Topic          string `json:"topic"`
			RelevanceScore string `json:"relevance_score"`
		} `json:"topics"`
	} `json:"feed"`
}

// TopGainersLosers AlphaVantage
type TopGainersLosers struct {
	Metadata           string   `json:"metadata"`
	LastUpdated        string   `json:"last_updated"`
	TopGainers         []Ticker `json:"top_gainers"`
	TopLosers          []Ticker `json:"top_losers"`
	MostActivelyTraded []Ticker `json:"most_actively_traded"`
}

// Quote TwelveData
type Quote struct {
	Symbol              string `json:"symbol"`
	Name                string `json:"name"`
	Exchange            string `json:"exchange"`
	MicCode             string `json:"mic_code"`
	Currency            string `json:"currency"`
	Datetime            string `json:"datetime"`
	Timestamp           int    `json:"timestamp"`
	Open                string `json:"open"`
	High                string `json:"high"`
	Low                 string `json:"low"`
	Rolling1DChange     string `json:"rolling_1d_change"`
	Rolling7DChange     string `json:"rolling_7d_change"`
	RollingPeriodChange string `json:"rolling_period_change"`
	IsMarketOpen        bool   `json:"is_market_open"`
	Close               string `json:"close"`
	Volume              string `json:"volume"`
	PreviousClose       string `json:"previous_close"`
	Change              string `json:"change"`
	PercentChange       string `json:"percent_change"`
	AverageVolume       string `json:"average_volume"`
	FiftyTwoWeek        struct {
		Low               string `json:"low"`
		High              string `json:"high"`
		LowChange         string `json:"low_change"`
		HighChange        string `json:"high_change"`
		LowChangePercent  string `json:"low_change_percent"`
		HighChangePercent string `json:"high_change_percent"`
		Range             string `json:"range"`
	} `json:"fifty_two_week"`

	ExtendedChange        string `json:"extended_change"`
	ExtendedPercentChange string `json:"extended_percent_change"`
	ExtendedPrice         string `json:"extended_price"`
	ExtendedTimestamp     int    `json:"extended_timestamp"`
	LastQuoteAt           int    `json:"last_quote_at"`
}

// RealTimePrice TwelveData
type RealTimePrice struct {
	Price string `json:"price"`
}

// Rating FMP
type Rating struct {
	Symbol                  string `json:"symbol"`
	Rating                  string `json:"rating"`
	OverallScore            int    `json:"overallScore"`
	DiscountedCashFlowScore int    `json:"discountedCashFlowScore"`
	ReturnOnEquityScore     int    `json:"returnOnEquityScore"`
	ReturnOnAssetsScore     int    `json:"returnOnAssetsScore"`
	DebtToEquityScore       int    `json:"debtToEquityScore"`
	PriceToEarningsScore    int    `json:"priceToEarningsScore"`
	PriceToBookScore        int    `json:"priceToBookScore"`
}

type Grade struct {
	Symbol         string `json:"symbol"`
	Date           string `json:"date"`
	GradingCompany string `json:"gradingCompany"`
	PreviousGrade  string `json:"previousGrade"`
	NewGrade       string `json:"newGrade"`
	Action         string `json:"action"`
}

type GradeSummery struct {
	Symbol     string `json:"symbol"`
	StrongBuy  int    `json:"strongBuy"`
	Buy        int    `json:"buy"`
	Hold       int    `json:"hold"`
	Sell       int    `json:"sell"`
	StrongSell int    `json:"strongSell"`
	Consensus  string `json:"consensus"`
}

type GradesHistorical struct {
	Symbol                   string `json:"symbol"`
	Date                     string `json:"date"`
	AnalystRatingsBuy        int    `json:"analystRatingsBuy"`
	AnalystRatingsHold       int    `json:"analystRatingsHold"`
	AnalystRatingsSell       int    `json:"analystRatingsSell"`
	AnalystRatingsStrongSell int    `json:"analystRatingsStrongSell"`
}

type EarningsReport struct {
	Symbol           string      `json:"symbol"`
	Date             string      `json:"date"`
	EpsActual        interface{} `json:"epsActual"`
	EpsEstimated     interface{} `json:"epsEstimated"`
	RevenueActual    interface{} `json:"revenueActual"`
	RevenueEstimated interface{} `json:"revenueEstimated"`
	LastUpdated      string      `json:"lastUpdated"`
}

type PriceEndOfDay struct {
	Symbol        string  `json:"symbol"`
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	Volume        int     `json:"volume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Vwap          float64 `json:"vwap"`
}

type SharesFloat struct {
	Symbol            string  `json:"symbol"`
	Date              string  `json:"date"`
	FreeFloat         float64 `json:"freeFloat"`
	FloatShares       int64   `json:"floatShares"`
	OutstandingShares int64   `json:"outstandingShares"`
}

type FinancialScores struct {
	Symbol           string  `json:"symbol"`
	ReportedCurrency string  `json:"reportedCurrency"`
	AltmanZScore     float64 `json:"altmanZScore"`
	PiotroskiScore   int     `json:"piotroskiScore"`
	WorkingCapital   int64   `json:"workingCapital"`
	TotalAssets      int64   `json:"totalAssets"`
	RetainedEarnings int64   `json:"retainedEarnings"`
	Ebit             int64   `json:"ebit"`
	MarketCap        int64   `json:"marketCap"`
	TotalLiabilities int64   `json:"totalLiabilities"`
	Revenue          int64   `json:"revenue"`
}

type PriceChanges struct {
	Symbol string  `json:"symbol"`
	D      float64 `json:"1D"`
	D1     float64 `json:"5D"`
	M      float64 `json:"1M"`
	M1     float64 `json:"3M"`
	M2     float64 `json:"6M"`
	Ytd    float64 `json:"ytd"`
	Y      float64 `json:"1Y"`
	Y1     float64 `json:"3Y"`
	Y2     float64 `json:"5Y"`
	Y3     float64 `json:"10Y"`
	Max    float64 `json:"max"`
}

type Ticker struct {
	Ticker           string `json:"ticker"`
	Price            string `json:"price"`
	ChangeAmount     string `json:"change_amount"`
	ChangePercentage string `json:"change_percentage"`
	Volume           string `json:"volume"`
}

type StockNewsEvent struct {
	Symbol  string
	Article Article
}

type Article struct {
	Symbol             string `clover:"symbol"`
	Title              string `clover:"title"`
	Sentiment          string `clover:"sentiment"`
	Content            string `clover:"content"`
	Link               string `clover:"link"`
	ScreenshotFilePath string `clover:"screenshotFilePath"`
}

type TickerAlt struct {
	Rank            int    `clover:"rank"`
	Symbol          string `clover:"symbol"`
	Name            string `clover:"name"`
	Sector          string `clover:"sector"`
	MCap            string `clover:"mCap"`
	Price           string `clover:"price"`
	Change          string `clover:"change"`
	Volume          string `clover:"volume"`
	PreMarketVolume string `clover:"preMarketVolume"`
	ChartLink       string `clover:"chartLink"`
}

type Post struct {
	Rank    int    `clover:"rank"`
	Symbol  string `clover:"symbol"`
	Heading string `clover:"heading"`
	Content string `clover:"content"`
	News    []News `clover:"news"`
}

type News struct {
	Time     string `clover:"time"`
	Headline string `clover:"headline"`
	Link     string `clover:"link"`
}
