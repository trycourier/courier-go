package courier

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type AutomationStep struct {
	Action           string      `json:"action"`
	Brand            string      `json:"brand,omitempty"`
	CancelationToken string      `json:"cancelation_token,omitempty"`
	Data             interface{} `json:"data,omitempty"`
	Duration         string      `json:"duration,omitempty"`
	If               string      `json:"if,omitempty"`
	List             string      `json:"list,omitempty"`
	Override         interface{} `json:"override,omitempty"`
	Profile          interface{} `json:"profile,omitempty"`
	Recipient        string      `json:"recipient,omitempty"`
	Ref              string      `json:"ref,omitempty"`
	Template         string      `json:"template,omitempty"`
	Until            string      `json:"until,omitempty"`
}

type Automation struct {
	CancelationToken string           `json:"cancelation_token,omitempty"`
	Steps            []AutomationStep `json:"steps"`
}

type AutomationInvokeBody struct {
	Automation Automation  `json:"automation"`
	Brand      string      `json:"brand,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Profile    interface{} `json:"profile,omitempty"`
	Recipient  string      `json:"recipient,omitempty"`
	Template   string      `json:"template,omitempty"`
}

type AutomationTemplateInvokeBody struct {
	Brand     string      `json:"brand,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Profile   interface{} `json:"profile,omitempty"`
	Recipient string      `json:"recipient,omitempty"`
	Template  string      `json:"template,omitempty"`
}

// Calls the POST /automations/invoke endpoint of the Courier API
func (c *Client) InvokeAutomation(ctx context.Context, body interface{}) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}

	response, err := c.API.SendRequestWithMaps(ctx, "POST", "/automations/invoke", bodyMap)
	if err != nil {
		return "", err
	}

	var runID string
	err = json.Unmarshal(response["runId"], &runID)
	if err != nil {
		return "", err
	}
	return runID, nil
}

// Calls the POST /automations/:templateId/invoke endpoint of the Courier API
func (c *Client) InvokeAutomationTemplate(ctx context.Context, templateId string, body interface{}, opts ...Option) (string, error) {
	bodyMap, err := toJSONMap(body)
	if err != nil {
		return "", err
	}

	req, err := c.prepareRequest(ctx, "POST", "/automations/"+templateId+"/invoke", bodyMap)
	if err != nil {
		return nil, err
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

	return payload["runId"], nil
}

// #FIXME: This code overlaps with newSendRequest in send.go and should be refactored
func (c *Client) prepareRequest(ctx context.Context, relativeUrl string, body io.Reader, method string, opts ...Option) (*http.Request, error) {
	if !relativeUrl.startsWith("/") {
		relativeUrl = "/" + relativeUrl
	}

	fullyQualifiedURL := c.API.BaseURL + relativeUrl
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
