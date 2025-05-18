package intelligence

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log/slog"
	"money/intelligence"
	"money/types"
	"net/url"
	"os"
	"reflect"
	"strings"
)

type Finviz struct {
	url *url.URL
}

func NewFinviz() (*Finviz, error) {
	u, err := url.Parse("https://finviz.com")
	return &Finviz{url: u}, err
}

func (f *Finviz) GetTicker(page playwright.Page) (types.TickerAlt, error) {
	var ticker types.TickerAlt
	rows, err := page.Locator("table.styled-table-new > tbody > tr.styled-row").All()
	if err != nil {
		return ticker, err
	}

	for idx, row := range rows {
		elements := map[string]int{
			"symbol": 2,
			"name":   3,
			"sector": 4,
			"mcap":   7,
			"price":  9,
			"change": 10,
			"volume": 11,
		}
		capture := make(map[string]string)
		for key, value := range elements {
			capture[key], err = row.Locator(fmt.Sprintf("td:nth-child(%d)", value)).TextContent(playwright.LocatorTextContentOptions{})
			if err != nil {
				return ticker, err
			}
		}

		chartLink := fmt.Sprintf(f.GetUrl(capture["symbol"]))
		ticker = types.TickerAlt{
			Rank:      idx + 1,
			Symbol:    capture["symbol"],
			Name:      capture["name"],
			Sector:    capture["sector"],
			MCap:      capture["mcap"],
			Price:     capture["price"],
			Change:    capture["change"],
			Volume:    capture["volume"],
			ChartLink: chartLink,
		}

		page, err = intelligence.GetPlaywrightPage(chartLink)
		if err != nil {
			return ticker, err
		}
		if _, err = f.screenshot(page, capture["symbol"], ""); err != nil {
			slog.Error(err.Error())
		}
	}
	return ticker, nil
}

func (f *Finviz) screenshot(page playwright.Page, symbol string, inPath string) (string, error) {
	if err := page.Locator("div[id='qc-cmp2-ui'] > div:nth-child(2) > div > button:nth-child(3)").Click(); err != nil {
		slog.Error(err.Error())
	}

	file, err := os.CreateTemp("", "*.png")
	if err != nil {
		return "", err
	}

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String(file.Name()),
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
	return file.Name(), nil
}

func (f *Finviz) GetGraph(symbol string) (string, error) {
	page, err := intelligence.GetPlaywrightPage(f.GetUrl(symbol))
	if err != nil {
		return "", err
	}

	file, err := os.CreateTemp("", "*.png")
	if err != nil {
		return "", err
	}

	path, err := f.screenshot(page, symbol, file.Name())
	if err != nil {
		return "", err
	}
	if err = file.Close(); err != nil {
		return "", err
	}

	return path, nil
}

func (f *Finviz) GetUrl(symbol string) string {
	f.url.Path = "quote.ashx"
	q := f.url.Query()
	q.Set("t", symbol)
	q.Set("ty", "c&p")
	q.Set("d&b", "1")
	f.url.RawQuery = q.Encode()
	return f.url.String()
}

func (f *Finviz) GetMetrics(symbol string) (types.Metrics, error) {
	m := &types.Metrics{}
	page, err := intelligence.GetPlaywrightPage(f.GetUrl(symbol))
	if err != nil {
		return *m, err
	}
	table := page.Locator("table.screener_snapshot-table-body > tbody")
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
