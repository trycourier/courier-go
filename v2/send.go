package courier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// SendBody is the JSON body expected by Courier's /send endpoint
type SendBody struct {
	Profile  interface{} `json:"profile"`
	Data     interface{} `json:"data"`
	Override interface{} `json:"override,omitempty"`
	Brand    string      `json:"brand,omitempty"`
}

// SendMessageRequestBody is the JSON body expected by
// Courier's /send endpoint that accepts a message object
type SendMessageRequestBody struct {
	Message interface{} `json:"message"`
}

// Send calls the /send endpoint of the Courier API (passing a struct)
func (c *Client) Send(ctx context.Context, eventID, recipientID string, body interface{}) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}
	return c.SendMap(ctx, eventID, recipientID, bodyMap)
}

// SendMap calls the /send endpoint of the Courier API (passing maps)
func (c *Client) SendMap(ctx context.Context, eventID, recipientID string, body map[string]interface{}) (string, error) {
	// these are required, so we accept them as separate params
	body["event"] = eventID
	body["recipient"] = recipientID

	response, err := c.API.SendRequestWithMaps(ctx, "POST", "/send", body)
	if err != nil {
		return "", err
	}

	var messageID string
	err = json.Unmarshal(response["messageId"], &messageID)
	if err != nil {
		return "", err
	}
	return messageID, nil
}

// SendMessage calls the enhanced /send endpoint of the Courier API with a message struct
func (c *Client) SendMessage(ctx context.Context, body SendMessageRequestBody) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}
	return c.SendMessageMap(ctx, bodyMap)
}

// SendMessageMap calls the enhanced /send endpoint of the Courier API with maps
func (c *Client) SendMessageMap(ctx context.Context, body map[string]interface{}) (string, error) {
	response, err := c.API.SendRequestWithMaps(ctx, "POST", "/send", body)
	if err != nil {
		return "", err
	}

	var requestID string
	err = json.Unmarshal(response["requestId"], &requestID)
	if err != nil {
		return "", err
	}
	return requestID, nil
}

type Option interface {
	apply(r *http.Request) error
}

type reqFunc func(r *http.Request) error

func (fn reqFunc) apply(r *http.Request) error {
	return fn(r)
}

func (c *Client) newSendRequest(ctx context.Context, body io.Reader, method string, opts ...Option) (*http.Request, error) {
	fullyQualifiedURL := fmt.Sprintf("%s/send", c.API.BaseURL)
	req, err := http.NewRequestWithContext(ctx, method, fullyQualifiedURL, body)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		if err := opt.apply(req); err != nil {
			return req, err
		}
	}
	return req, err
}

func withHeader(header, value string) Option {
	return reqFunc(func(r *http.Request) error {
		if header != "" && value != "" {
			r.Header.Add(header, value)
		}
		return nil
	})
}

func WithIdempotencyKey(value string) Option {
	return withHeader("Idempotency-Key", value)
}

func WithIdempotencyKeyExpiration(expiration time.Time) Option {
	return withHeader("x-idempotency-expiration", strconv.FormatInt(expiration.UnixMilli(), 10))
}

func (c *Client) SendMessageWithOptions(ctx context.Context, body map[string]interface{}, method string, opts ...Option) (string, error) {
	if method == "GET" {
		body = nil
	}
	bytesBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	buf := bytes.NewReader(bytesBody)

	req, err := c.newSendRequest(ctx, buf, method, opts...)
	if err != nil {
		return "", err
	}

	res, err := c.API.ExecuteRequest(req)
	if err != nil {
		return "", err
	}

	var payload map[string]string
	err = json.Unmarshal(res, &payload)
	if err != nil {
		return "", err
	}

	return payload["requestId"], nil
}
