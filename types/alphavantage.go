package types

import (
	"github.com/ostafen/clover/v2/query"
	"money/database"
)

// NewsSentiments AlphaVantage
type NewsSentiments struct {
	Items                    string `json:"items"`
	SentimentScoreDefinition string `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string `json:"relevance_score_definition"`
	Feeds                    []Feed `json:"feed"`
}

type Feed struct {
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
}

// TopGainersLosers AlphaVantage
type TopGainersLosers struct {
	Metadata           string   `json:"metadata"`
	LastUpdated        string   `json:"last_updated"`
	TopGainers         []Ticker `json:"top_gainers"`
	TopLosers          []Ticker `json:"top_losers"`
	MostActivelyTraded []Ticker `json:"most_actively_traded"`
}

type Ticker struct {
	Ticker           string `json:"ticker"`
	Price            string `json:"price"`
	ChangeAmount     string `json:"change_amount"`
	ChangePercentage string `json:"change_percentage"`
	Volume           string `json:"volume"`
}

func (f *Feed) Save() error {
	db, err := database.Load("alphavantage-sentiments")
	if err != nil {
		return err
	}
	if err = db.Save(f.Title, f); err != nil {
		return err
	}
	if err = db.Client.Close(); err != nil {
		return err
	}
	return nil
}

func (f *Feed) GetAll() ([]Feed, error) {
	all := make([]Feed, 0)

	db, err := database.Load("alphavantage-sentiments")
	if err != nil {
		return nil, err
	}
	posts, err := db.Client.FindAll(query.NewQuery("alphavantage-sentiments"))
	if err != nil {
		return nil, err
	}

	var event Feed
	for _, p := range posts {
		if err = p.Unmarshal(&event); err != nil {
			return nil, err
		}
		all = append(all, event)
	}

	if err = db.Client.Close(); err != nil {
		return nil, err
	}
	return all, err
}
