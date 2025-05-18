package intelligence

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log/slog"
	"money/types"
	"net/url"
	"os"
)

type StockNewsAI struct {
	url *url.URL
}

func NewStockNewsAI() (*StockNewsAI, error) {
	u, err := url.Parse("https://stocknews.ai/ai-events")
	return &StockNewsAI{url: u}, err
}

func (s *StockNewsAI) GetArticles(page playwright.Page) (types.Article, error) {
	var article types.Article

	postsSelector := "div.mx-auto > div:nth-child(3) > div.relative"
	if _, err := page.WaitForSelector(fmt.Sprintf("%s:nth-child(10)", postsSelector)); err != nil {
		return article, err
	}
	posts, err := page.Locator(postsSelector).All()
	if err != nil {
		return article, err
	}
	if err = page.Locator("div[id='crisp-chatbox'] > div:nth-child(1) > a > span:first-child > span:first-child > span:first-child > span:first-child > span").Click(); err != nil {
		slog.Error(err.Error())
	}

	for _, post := range posts {
		internal := post.Locator("div:nth-child(2)")
		heading := internal.Locator("div:nth-child(1)")
		symbol, err := heading.Locator("a > span").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return article, err
		}
		sentiment, err := heading.Locator("label > span").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return article, err
		}
		a := internal.Locator("article")
		title, err := a.Locator("div:nth-child(1) > div:nth-child(1) > h1").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return article, err
		}
		content, err := a.Locator("div:nth-child(2) > div:nth-child(1) > p").TextContent(playwright.LocatorTextContentOptions{})

		s.url.Path = fmt.Sprintf("us-stock/%s", symbol)
		article = types.Article{
			Symbol:    symbol,
			Title:     title,
			Sentiment: sentiment,
			Content:   content,
			Link:      s.url.String(),
		}
	}
	return article, nil
}

func (s *StockNewsAI) screenshot(post playwright.Locator, symbol string) error {
	file, err := os.CreateTemp("", "*.png")
	if err != nil {
		return err
	}

	if _, err = post.Screenshot(playwright.LocatorScreenshotOptions{
		Path: playwright.String(fmt.Sprintf(file.Name(), symbol)),
	}); err != nil {
		return err
	}
	return nil
}
