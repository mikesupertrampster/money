package intelligence

import (
	"money/types"
	"net/url"
	"os"
	"strings"

	"github.com/playwright-community/playwright-go"
	"log/slog"
)

type StockTitans struct {
	url *url.URL
}

func NewStockTitans() (*StockTitans, error) {
	u, err := url.Parse("https://www.stocktitan.net/scanner/momentum")
	return &StockTitans{url: u}, err
}

func (s *StockTitans) Scrape(page playwright.Page) (types.Post, error) {
	p := types.Post{}

	if err := page.Locator("div[id='cmpbox'] > * a.cmpboxbtnyes").Click(); err != nil {
		return p, err
	}
	columns, err := page.Locator("div.scanner-momentum > div").All()
	if err != nil {
		return p, err
	}

	for _, col := range columns {
		heading, err := col.Locator("div:nth-child(2) > h2").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return p, err
		}

		posts, err := col.Locator("div.body > div").All()
		if err != nil {
			return p, err
		}
		for _, post := range posts {
			contents, err := post.Locator("div.content").All()
			if err != nil {
				return p, err
			}

			var symbol string
			for idx, content := range contents {
				symbol, err = content.Locator("div:first-child > div.symbol > a").TextContent(playwright.LocatorTextContentOptions{})
				if err != nil {
					return p, err
				}

				var paragraph string
				blocks, err := content.Locator("div.commentary").All()
				if err != nil {
					slog.Error(err.Error())
				}
				for _, block := range blocks {
					p, err := block.TextContent(playwright.LocatorTextContentOptions{Timeout: playwright.Float(1000)})
					if err != nil {
						slog.Error(err.Error())
					}
					paragraph = paragraph + p
				}

				news, err := content.Locator("div:nth-child(4) > ul > li").All()
				if err != nil {
					slog.Error(err.Error())
				}

				n := make([]types.News, 0)
				for _, entries := range news {
					time, err := entries.Locator("div:first-child > div:first-child > span:first-child").TextContent(playwright.LocatorTextContentOptions{})
					if err != nil {
						slog.Error(err.Error())
					}

					link := entries.Locator("div.news-content > a")
					headline, err := link.TextContent(playwright.LocatorTextContentOptions{})
					if err != nil {
						slog.Error(err.Error())
					}
					href, err := link.GetAttribute("href")
					if err != nil {
						slog.Error(err.Error())
					}

					s.url.Path = href

					n = append(n, types.News{
						Time:     strings.TrimSpace(time),
						Headline: strings.TrimSpace(headline),
						Link:     strings.TrimSpace(s.url.String()),
					})
				}

				if err = s.screenshot(post, symbol); err != nil {
					slog.Error(err.Error())
				}

				p = types.Post{
					Rank:    idx + 1,
					Symbol:  symbol,
					Heading: strings.TrimSpace(heading),
					Content: strings.TrimSpace(paragraph),
					News:    n,
				}
			}
		}
	}

	return p, nil
}

func (s *StockTitans) screenshot(post playwright.Locator, symbol string) error {
	file, err := os.CreateTemp("", "*.png")
	if err != nil {
		return err
	}

	if _, err := post.Screenshot(playwright.LocatorScreenshotOptions{
		Path: playwright.String(file.Name()),
	}); err != nil {
		return err
	}
	return nil
}
