package intelligence

import (
	"github.com/playwright-community/playwright-go"
	"os"
)

func GetPlaywrightPage() (*playwright.Page, error) {
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
	return &page, nil
}

func MkDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}
