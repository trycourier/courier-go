package courier

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiKey  string
	BaseUrl string
}

func CourierClient(apiKey string, baseUrl string) *Client {
	if baseUrl == "" {
		baseUrl = "https://api.trycourier.app"
	}
	return &Client{
		ApiKey:  apiKey,
		BaseUrl: baseUrl,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+s.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "courier-go/0.0.1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
