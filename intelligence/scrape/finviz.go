package intelligence

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"money/intelligence"
	"money/types"
	"net/url"
	path "path"
	"reflect"
	"strings"
	"time"
)

type Finviz struct {
	url  *url.URL
	page playwright.Page
	dir  string
}

func NewFinviz() (*Finviz, error) {
	dir := "_cache/screenshots/finviz"
	if err := intelligence.MkDir(dir); err != nil {
		return nil, err
	}

	u, err := url.Parse("https://finviz.com")
	page, err := intelligence.GetPlaywrightPage()
	if err != nil {
		return nil, err
	}
	return &Finviz{
		url:  u,
		page: *page,
		dir:  dir,
	}, err
}

func (f *Finviz) GetMetrics(symbol string) (types.Metrics, error) {
	m := types.Metrics{}

	if _, err := f.page.Goto(f.ChartPage(symbol), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return m, err
	}
	if err := f.page.Locator("div[id='qc-cmp2-ui'] > div:nth-child(2) > div > button:nth-child(3)").Click(); err != nil {
		return m, err
	}

	m, err := f.scrapeMetrics(symbol)
	if err != nil {
		return m, err
	}

	p, err := f.getGraph(symbol)
	if err != nil {
		return m, err
	}
	m.Image = p

	if err = f.page.Close(); err != nil {
		return m, err
	}

	return m, nil
}

func (f *Finviz) scrapeMetrics(symbol string) (types.Metrics, error) {
	m := &types.Metrics{
		Symbol:    symbol,
		TimeStamp: time.Now(),
	}

	table := f.page.Locator("table.screener_snapshot-table-body > tbody")
	properties := map[string]struct {
		x int
		y int
	}{
		"MarketCap":           {x: 2, y: 2},
		"FloatOutstanding":    {x: 10, y: 1},
		"FloatTotal":          {x: 10, y: 2},
		"FiftyTwoWeekRange":   {x: 10, y: 6},
		"FiftyTwoWeekHighPct": {x: 10, y: 7},
		"FiftyTwoWeekLowPct":  {x: 10, y: 8},
		"RSI14":               {x: 10, y: 9},
		"Volume":              {x: 10, y: 13},
		"Price":               {x: 12, y: 12},
		"Change":              {x: 12, y: 13},
	}
	for name, coordinate := range properties {
		value, err := table.Locator(fmt.Sprintf("tr:nth-child(%d) > td:nth-child(%d)", coordinate.y, coordinate.x)).TextContent(playwright.LocatorTextContentOptions{})
		if err != nil {
			return *m, err
		}
		v := reflect.ValueOf(m)
		if v.Kind() != reflect.Ptr {
			return *m, fmt.Errorf("expected a pointer")
		}

		field := reflect.Indirect(reflect.ValueOf(m)).FieldByName(name)
		if field.Kind() != reflect.Invalid {
			if name == "FiftyTwoWeekRange" {
				s := strings.Split(value, " - ")
				m.FiftyTwoWeekLow = s[0]
				m.FiftyTwoWeekHigh = s[1]
			}
			field.SetString(value)
		}
	}
	return *m, nil
}

func (f *Finviz) getGraph(symbol string) (string, error) {
	p := path.Join(f.dir, symbol+".png")
	if _, err := f.page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String(p),
		FullPage: playwright.Bool(true),
		Clip: &playwright.Rect{
			X:      0,
			Y:      310,
			Width:  1280,
			Height: 550,
		},
	}); err != nil {
		return "", err
	}
	return p, nil
}

func (f *Finviz) ChartPage(symbol string) string {
	f.url.Path = "quote.ashx"
	q := f.url.Query()
	q.Set("t", symbol)
	q.Set("p", "d")
	f.url.RawQuery = q.Encode()
	return f.url.String()
}
