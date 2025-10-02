// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/trycourier/courier-go"
	"github.com/trycourier/courier-go/internal/testutil"
	"github.com/trycourier/courier-go/option"
)

func TestAutomationInvokeInvokeAdHocWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := courier.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Automations.Invoke.InvokeAdHoc(context.TODO(), courier.AutomationInvokeInvokeAdHocParams{
		Automation: courier.AutomationInvokeInvokeAdHocParamsAutomation{
			Steps: []courier.AutomationInvokeInvokeAdHocParamsAutomationStepUnion{{
				OfAutomationAddToDigestStep: &courier.AutomationInvokeInvokeAdHocParamsAutomationStepAutomationAddToDigestStep{
					AutomationStepParam: courier.AutomationStepParam{
						If:  courier.String("if"),
						Ref: courier.String("ref"),
					},
					Action:              "add-to-digest",
					SubscriptionTopicID: "RAJE97CMT04KDJJ88ZDS2TP1690S",
				},
			}},
			CancelationToken: courier.String("cancelation_token"),
		},
		Brand: courier.String("brand"),
		Data: map[string]any{
			"foo": "bar",
		},
		Profile:   map[string]interface{}{},
		Recipient: courier.String("recipient"),
		Template:  courier.String("template"),
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAutomationInvokeInvokeByTemplateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := courier.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Automations.Invoke.InvokeByTemplate(
		context.TODO(),
		"templateId",
		courier.AutomationInvokeInvokeByTemplateParams{
			AutomationInvokeParams: courier.AutomationInvokeParams{
				Brand: courier.String("brand"),
				Data: map[string]any{
					"foo": "bar",
				},
				Profile:   map[string]interface{}{},
				Recipient: courier.String("recipient"),
				Template:  courier.String("template"),
			},
		},
	)
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
