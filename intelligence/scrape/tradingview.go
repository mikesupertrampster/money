package intelligence

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"money/intelligence"
	"net/url"
)

type TradingView struct {
	url  *url.URL
	page playwright.Page
	dir  string
}

func NewTradingView() (*TradingView, error) {
	dir := "_cache/screenshots/tradingview"
	if err := intelligence.MkDir(dir); err != nil {
		return nil, err
	}

	u, err := url.Parse("https://www.tradingview.com")
	page, err := intelligence.GetPlaywrightPage()
	if err != nil {
		return nil, err
	}
	return &TradingView{
		url:  u,
		page: *page,
		dir:  dir,
	}, err
}

func (t *TradingView) GetFinancials(symbol string) (string, error) {
	t.url.Path = fmt.Sprintf("symbols/NASDAQ-%s/", symbol)
	if _, err := t.page.Goto(t.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	box, err := t.page.Locator("div[data-an-widget-id='financials-overview-id']").BoundingBox()
	if err != nil {
		return "", err
	}

	return t.get(
		symbol,
		"financials",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height + 20,
		},
	)
}

func (t *TradingView) GetTechnicals(symbol string) (string, error) {
	t.url.Path = fmt.Sprintf("symbols/NASDAQ-%s/", symbol)
	if _, err := t.page.Goto(t.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	box, err := t.page.Locator("div[data-an-widget-id='technicals-analyst-curve-layout']").BoundingBox()
	if err != nil {
		return "", err
	}

	return t.get(
		symbol,
		"technicals",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height + 20,
		},
	)
}

func (t *TradingView) GetSeasonals(symbol string) (string, error) {
	t.url.Path = fmt.Sprintf("symbols/NASDAQ-%s/", symbol)
	if _, err := t.page.Goto(t.url.String(), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return "", err
	}

	box, err := t.page.Locator("div[data-an-widget-id='seasonals']").BoundingBox()
	if err != nil {
		return "", err
	}

	return t.get(
		symbol,
		"seasonals",
		&playwright.Rect{
			X:      box.X - 10,
			Y:      box.Y - 10,
			Width:  box.Width + 20,
			Height: box.Height + 20,
		},
	)
}

func (t *TradingView) get(symbol string, section string, rect *playwright.Rect) (string, error) {
	p := fmt.Sprintf("%s/%s_%s.png", t.dir, symbol, section)
	if _, err := t.page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String(p),
		FullPage: playwright.Bool(true),
		Clip:     rect,
	}); err != nil {
		return "", err
	}
	return p, nil
}
