package intelligence

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"money/intelligence"
	"net/url"
	"time"
)

type QuiverQuant struct {
	url  *url.URL
	page playwright.Page
	dir  string
}

func NewQuiverQuant() (*QuiverQuant, error) {
	dir := "_cache/screenshots/quiverquant"
	if err := intelligence.MkDir(dir); err != nil {
		return nil, err
	}

	u, err := url.Parse("https://www.quiverquant.com")
	page, err := intelligence.GetPlaywrightPage()
	if err != nil {
		return nil, err
	}
	return &QuiverQuant{
		url:  u,
		page: *page,
		dir:  dir,
	}, err
}

func (q *QuiverQuant) GetTopOwnership(symbol string) (string, error) {
	q.url.Path = fmt.Sprintf("stock/%s/ownership", symbol)
	if _, err := q.page.Goto(q.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	_, err := q.page.WaitForSelector("div.item-ownership > * td")
	if err != nil {
		return "", err
	}

	box, err := q.page.Locator("div.item-ownership").BoundingBox()
	if err != nil {
		return "", err
	}

	return q.get(
		symbol,
		"top_ownership",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height,
		},
	)
}

func (q *QuiverQuant) GetOwnership(symbol string) (string, error) {
	q.url.Path = fmt.Sprintf("stock/%s/institutions", symbol)
	if _, err := q.page.Goto(q.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	_, err := q.page.WaitForSelector("div.item-institutions > * td")
	if err != nil {
		return "", err
	}

	box, err := q.page.Locator("div.item-institutions").BoundingBox()
	if err != nil {
		return "", err
	}

	return q.get(
		symbol,
		"ownership",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height,
		},
	)
}

func (q *QuiverQuant) GetWhaleActivity(symbol string) (string, error) {
	q.url.Path = fmt.Sprintf("stock/%s/institutions", symbol)
	if _, err := q.page.Goto(q.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	if _, err := q.page.WaitForSelector("div.whale-graph"); err != nil {
		return "", err
	}

	div := q.page.Locator("div.whale-graph")
	box, err := q.page.Locator("div.item-overview").Filter(playwright.LocatorFilterOptions{
		Has: div,
	}).BoundingBox()
	if err != nil {
		return "", err
	}

	return q.get(
		symbol,
		"whale_overview",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height + 20,
		},
	)
}

func (q *QuiverQuant) GetAnalystRating(symbol string) (string, error) {
	q.url.Path = fmt.Sprintf("stock/%s/forecast", symbol)
	if _, err := q.page.Goto(q.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	if _, err := q.page.WaitForSelector("div.buy-sell-hold-outer > * svg"); err != nil {
		return "", err
	}

	box, err := q.page.Locator("div.buy-sell-hold-outer").BoundingBox()
	if err != nil {
		return "", err
	}

	return q.get(
		symbol,
		"analyst_rating",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height + 20,
		},
	)
}

func (q *QuiverQuant) GetForecasts(symbol string) (string, error) {
	q.url.Path = fmt.Sprintf("stock/%s/forecast", symbol)
	if _, err := q.page.Goto(q.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	time.Sleep(time.Duration(4) * time.Second)

	box, err := q.page.Locator("div.price-target-outer").BoundingBox()
	if err != nil {
		return "", err
	}

	return q.get(
		symbol,
		"forecasts",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height + 20,
		},
	)
}

func (q *QuiverQuant) GetGeneral(symbol string) (string, error) {
	q.url.Path = fmt.Sprintf("stock/%s", symbol)
	if _, err := q.page.Goto(q.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	time.Sleep(time.Duration(4) * time.Second)

	box, err := q.page.Locator("div.ticker-content").BoundingBox()
	if err != nil {
		return "", err
	}

	return q.get(
		symbol,
		"general",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height + 20,
		},
	)
}

func (q *QuiverQuant) get(symbol string, section string, rect *playwright.Rect) (string, error) {
	p := fmt.Sprintf("%s/%s_%s.png", q.dir, symbol, section)
	if _, err := q.page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String(p),
		FullPage: playwright.Bool(true),
		Clip:     rect,
	}); err != nil {
		return "", err
	}
	return p, nil
}
