package intelligence

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"money/intelligence"
	"net/url"
)

type StockTwits struct {
	url  *url.URL
	page playwright.Page
	dir  string
}

func NewStockTwits() (*StockTwits, error) {
	dir := "_cache/screenshots/stocktwits"
	if err := intelligence.MkDir(dir); err != nil {
		return nil, err
	}

	u, err := url.Parse("https://stocktwits.com/")
	page, err := intelligence.GetPlaywrightPage()
	if err != nil {
		return nil, err
	}
	return &StockTwits{
		url:  u,
		page: *page,
		dir:  dir,
	}, err
}

func (s *StockTwits) GetSentiments(symbol string) error {
	s.url.Path = fmt.Sprintf("symbol/%s/sentiment", symbol)
	if _, err := s.page.Goto(s.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return err
	}
	if err := s.page.Locator("button[id='onetrust-accept-btn-handler']").Click(); err != nil {
		return err
	}

	all, err := s.page.Locator("div.grid > div:first-child.tabletXxl-down\\|m-auto").All()
	if err != nil {
		return err
	}
	for i, div := range all {
		var screenshot string
		switch i {
		case 0:
			screenshot = "sentiment"
		case 1:
			screenshot = "message_volume"
		case 2:
			screenshot = "participant_ratio"
		}

		if _, err = div.Screenshot(playwright.LocatorScreenshotOptions{
			Path: playwright.String(fmt.Sprintf("%s/%s_%s.png", s.dir, symbol, screenshot)),
		}); err != nil {
			return err
		}
	}

	return nil
}
