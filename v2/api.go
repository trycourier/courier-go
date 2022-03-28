package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// APIConfiguration represents the core data needed to communicate with the Courier API
type APIConfiguration struct {
	AuthToken  string
	BaseURL    string
	SDKVersion string
}

// HTTPError is a customer error returned when the API doesn't respond with a 2XX
type HTTPError struct {
	StatusCode   int
	ErrorMessage *string
}

func (err *HTTPError) Error() string {
	if err.ErrorMessage == nil {
		return fmt.Sprintf("HTTP Error %d: %s", err.StatusCode, "<no HTTP body>")
	}
	return fmt.Sprintf("HTTP Error %d: %s", err.StatusCode, *err.ErrorMessage)
}

// SendRequestWithJSON wraps HTTPSendBytes
func (api *APIConfiguration) SendRequestWithJSON(ctx context.Context, method string, relativePath string, body interface{}, response interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bytes, err := api.SendRequestWithBytes(ctx, method, relativePath, jsonBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return err
	}

	return nil
}

// SendRequestWithMaps wraps HTTPSendBytes
func (api *APIConfiguration) SendRequestWithMaps(ctx context.Context, method string, relativePath string, body map[string]interface{}) (map[string]json.RawMessage, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	bytes, err := api.SendRequestWithBytes(ctx, method, relativePath, jsonBody)
	if err != nil {
		return nil, err
	}

	var objmap map[string]json.RawMessage
	err = json.Unmarshal(bytes, &objmap)
	if err != nil {
		return nil, err
	}
	return objmap, nil
}

// SendRequestWithBytes wraps SendRequestWithReader
func (api *APIConfiguration) SendRequestWithBytes(ctx context.Context, method string, relativePath string, body []byte) ([]byte, error) {
	buf := bytes.NewReader(body)
	return api.SendRequestWithReader(ctx, method, relativePath, buf)
}

// SendRequestWithReader wraps HTTPRequest
func (api *APIConfiguration) SendRequestWithReader(ctx context.Context, method string, relativePath string, body io.Reader) ([]byte, error) {
	fullyQualifiedURL := api.BaseURL + relativePath
	if method == "GET" {
		body = nil
	}
	req, err := http.NewRequestWithContext(ctx, method, fullyQualifiedURL, body)
	if err != nil {
		return nil, err
	}
	return api.ExecuteRequest(req)
}

// ExecuteRequest issues an HTTP request and sets headers expected by the Courier API
func (api *APIConfiguration) ExecuteRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+api.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "courier-go/"+api.SDKVersion)

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

	// Success is indicated with 2xx status codes
	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		errMessage := string(body[:])
		return nil, &HTTPError{resp.StatusCode, &errMessage}
	}
	return body, nil
}
