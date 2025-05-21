package intelligence

import (
	"github.com/playwright-community/playwright-go"
	"log/slog"
	"money/intelligence"
	"money/types"
	"net/url"
	"path"
	"strings"
	"time"
)

type StockTitans struct {
	url  *url.URL
	page playwright.Page
	dir  string
}

func NewStockTitans() (*StockTitans, error) {
	dir := "_cache/screenshots/stocktitans"
	if err := intelligence.MkDir(dir); err != nil {
		return nil, err
	}

	u, err := url.Parse("https://www.stocktitan.net/scanner/momentum")
	page, err := intelligence.GetPlaywrightPage()
	if err != nil {
		return nil, err
	}

	return &StockTitans{
		url:  u,
		page: *page,
		dir:  dir,
	}, err
}

func (s *StockTitans) GetNews() ([]types.TitanPost, error) {
	t := make([]types.TitanPost, 0)

	if _, err := s.page.Goto(s.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return t, err
	}
	if err := s.page.Locator("div[id='cmpbox'] > * a.cmpboxbtnyes").Click(); err != nil {

		slog.Error(err.Error())
	}
	if err := s.page.Locator("div.adthrive-sticky-outstream > * path").Click(); err != nil {
		slog.Error(err.Error())
	}

	columns, err := s.page.Locator("div.scanner-momentum > div").All()
	if err != nil {
		return t, err
	}

	for _, col := range columns {
		heading, err := col.Locator("div:nth-child(2) > h2").TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return t, err
		}

		posts, err := col.Locator("div.body > div").All()
		if err != nil {
			return t, err
		}
		for _, post := range posts {

			contents, err := post.Locator("div.content").All()
			if err != nil {
				return t, err
			}

			var symbol string
			for idx, content := range contents {
				symbol, err = content.Locator("div:first-child > div.symbol > a").TextContent(playwright.LocatorTextContentOptions{})
				if err != nil {
					return t, err
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

				n := make([]types.TitanNews, 0)
				for _, entries := range news {
					date, err := entries.Locator("div:first-child > div:first-child > span:first-child").TextContent(playwright.LocatorTextContentOptions{})
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

					n = append(n, types.TitanNews{
						Time:     strings.TrimSpace(date),
						Headline: strings.TrimSpace(headline),
						Link:     strings.TrimSpace(strings.TrimSpace("https://www.stocktitan.net" + href)),
					})
				}

				p, err := s.screenshot(post, symbol)
				if err != nil {
					slog.Error(err.Error())
				}

				t = append(t, types.TitanPost{
					Rank:      idx + 1,
					Symbol:    symbol,
					Heading:   strings.TrimSpace(heading),
					Content:   strings.TrimSpace(paragraph),
					News:      n,
					TimeStamp: time.Now(),
					Image:     p,
				})
			}
		}
	}
	return t, nil
}

func (s *StockTitans) screenshot(post playwright.Locator, symbol string) (string, error) {
	p := path.Join(s.dir, symbol+".png")
	if _, err := post.Screenshot(playwright.LocatorScreenshotOptions{
		Path: playwright.String(p),
	}); err != nil {
		return "", err
	}
	return p, nil
}
