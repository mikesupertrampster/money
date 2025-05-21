package intelligence

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log/slog"
	"money/intelligence"
	"money/types"
	"net/url"
	"path"
	"time"
)

type StockNewsAi struct {
	url  *url.URL
	page playwright.Page
	dir  string
}

func NewStockNewsAi() (*StockNewsAi, error) {
	dir := "_cache/screenshots/stocknews"
	if err := intelligence.MkDir(dir); err != nil {
		return nil, err
	}

	u, err := url.Parse("https://stocknews.ai/ai-events")
	page, err := intelligence.GetPlaywrightPage()
	if err != nil {
		return nil, err
	}
	return &StockNewsAi{
		url:  u,
		page: *page,
		dir:  dir,
	}, err
}

func (s *StockNewsAi) GetEvents() ([]types.AiEvent, error) {
	events := make([]types.AiEvent, 0)

	if _, err := s.page.Goto(s.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return events, err
	}

	postsSelector := "div.mx-auto > div:nth-child(3) > div.relative"
	if _, err := s.page.WaitForSelector(fmt.Sprintf("%s:nth-child(10)", postsSelector)); err != nil {
		return events, err
	}
	posts, err := s.page.Locator(postsSelector).All()
	if err != nil {
		return events, err
	}
	if err = s.page.Locator("div[id='crisp-chatbox'] > div:nth-child(1) > a > span:first-child > span:first-child > span:first-child > span:first-child > span").Click(); err != nil {
		slog.Error(err.Error())
	}

	for _, post := range posts {
		internal := post.Locator("div:nth-child(2)")
		heading := internal.Locator("div:nth-child(1)")
		symbol, err := heading.Locator("a > span").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return events, err
		}
		sentiment, err := heading.Locator("label > span").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return events, err
		}
		a := internal.Locator("article")
		title, err := a.Locator("div:nth-child(1) > div:nth-child(1) > h1").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return events, err
		}
		content, err := a.Locator("div:nth-child(2) > div:nth-child(1) > p").TextContent(playwright.LocatorTextContentOptions{})

		p, err := s.screenshot(post, symbol)
		if err != nil {
			slog.Error(err.Error())
		}

		s.url.Path = fmt.Sprintf("us-stock/%s", symbol)
		event := types.AiEvent{
			Symbol:    symbol,
			Title:     title,
			Sentiment: sentiment,
			Content:   content,
			Link:      s.url.String(),
			Image:     p,
			TimeStamp: time.Now(),
		}
		events = append(events, event)
	}
	return events, nil
}

func (s *StockNewsAi) screenshot(post playwright.Locator, symbol string) (string, error) {
	p := path.Join(s.dir, symbol+".png")

	if _, err := post.Screenshot(playwright.LocatorScreenshotOptions{
		Path: playwright.String(p),
	}); err != nil {
		return "", err
	}
	return p, nil
}
