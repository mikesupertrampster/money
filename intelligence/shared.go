package intelligence

import (
	"encoding/json"
	"errors"
	"github.com/playwright-community/playwright-go"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func MapToStruct(data interface{}, result interface{}) error {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonStr, &result); err != nil {
		return err
	}
	return nil
}

func HttpGet(url *url.URL, result interface{}) error {
	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err = resp.Body.Close(); err != nil {
		return err
	}
	if err = json.Unmarshal(body, result); err != nil {
		return err
	}
	return nil
}

func MkDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

func GetPlaywrightPage(site string) (playwright.Page, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	browser, err := pw.Firefox.Launch()
	if err != nil {
		return nil, err
	}
	page, err := browser.NewPage()
	if err != nil {
		return nil, err
	}
	if _, err = page.Goto(site, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
	}); err != nil {
		return nil, err
	}

	return page, nil
}

func getMedalEmoji(idx int) string {
	emoji := ":"
	switch idx {
	case 0:
		emoji = " :first_place:"
	case 1:
		emoji = " :second_place:"
	case 2:
		emoji = " :third_place:"
	}
	return emoji
}

func getUpDownEmoji(number string) (string, error) {
	change, err := strconv.ParseFloat(strings.Replace(number, "%", "", -1), 32)
	if err != nil {
		return "", err
	}

	e := ":up:"
	if change < 0 {
		e = ":small_red_triangle_down:"
	}
	return e, nil
}
