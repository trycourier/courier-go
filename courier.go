package courier

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Authorization struct {
	Type  string
	Token string
}

type Client struct {
	ApiKey        string
	BaseUrl       string
	Username      string
	Password      string
	Authorization Authorization
}

func CourierClient(apiKey string, baseUrl string, username string, password string) *Client {
	_baseUrl := ""
	_apiKey := ""
	_username := ""
	_password := ""
	var _authorization Authorization

	// Override baseUrl if passed as a param or set as an environment variable
	if baseUrl != "" {
		_baseUrl = baseUrl
	} else if os.Getenv("COURIER_BASE_URL") != "" {
		_baseUrl = os.Getenv("COURIER_BASE_URL")
	} else {
		_baseUrl = "https://api.courier.com"
	}

	// Token Auth takes precedence; If no token auth, then Basic Auth
	if apiKey != "" {
		_apiKey = apiKey
		_authorization.Type = "Bearer"
		_authorization.Token = apiKey
	} else if os.Getenv("COURIER_AUTH_TOKEN") != "" {
		_apiKey = os.Getenv("COURIER_AUTH_TOKEN")
		_authorization.Type = "Bearer"
		_authorization.Token = os.Getenv("COURIER_AUTH_TOKEN")
	} else if username != "" && password != "" {
		_username = username
		_password = password
		_authorization.Type = "Basic"
		_authorization.Token = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	} else if os.Getenv("COURIER_AUTH_USERNAME") != "" && os.Getenv("COURIER_AUTH_PASSWORD") != "" {
		_username = os.Getenv("COURIER_AUTH_USERNAME")
		_password = os.Getenv("COURIER_AUTH_PASSWORD")
		_authorization.Type = "Basic"
		_authorization.Token = base64.StdEncoding.EncodeToString([]byte(_username + ":" + _password))
	}

	return &Client{
		ApiKey:        _apiKey,
		BaseUrl:       _baseUrl,
		Username:      _username,
		Password:      _password,
		Authorization: _authorization,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	authorization := s.Authorization.Type + " " + s.Authorization.Token
	req.Header.Set("Authorization", authorization)
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
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
