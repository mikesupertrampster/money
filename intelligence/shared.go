package intelligence

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
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
