package courier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	ApiKey string
	BaseUrl string
}

type Profile struct {
	Profile interface{} `json:"profile"`
}

func CourierClient(apiKey string, baseUrl string) *Client {
	if baseUrl == "" {
		baseUrl = "https://api.trycourier.app"
	}
	return &Client{
		ApiKey: apiKey,
		BaseUrl: baseUrl,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer " + s.ApiKey)
	req.Header.Set("Content-Type", "application/json")
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
	// log.Println(string(body))
	return body, nil
}

func (s *Client) GetProfile(id string) (*Profile, error) {
	url := fmt.Sprintf(s.BaseUrl + "/profiles/%s", id)
	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Profile
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) MergeProfile(id string, profile []byte) error {
	log.Println(string(profile))
	url := fmt.Sprintf(s.BaseUrl + "/profiles/%s", id)
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(profile))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}
