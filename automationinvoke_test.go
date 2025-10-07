// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/courier-go"
	"github.com/stainless-sdks/courier-go/internal/testutil"
	"github.com/stainless-sdks/courier-go/option"
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
				OfAutomationDelayStep: &courier.AutomationInvokeInvokeAdHocParamsAutomationStepAutomationDelayStep{
					Action:   "delay",
					Duration: courier.String("duration"),
					Until:    courier.String("20240408T080910.123"),
				},
			}, {
				OfAutomationSendStep: &courier.AutomationInvokeInvokeAdHocParamsAutomationStepAutomationSendStep{
					Action: "send",
					Brand:  courier.String("brand"),
					Data: map[string]any{
						"foo": "bar",
					},
					Profile: map[string]any{
						"foo": "bar",
					},
					Recipient: courier.String("recipient"),
					Template:  courier.String("64TP5HKPFTM8VTK1Y75SJDQX9JK0"),
				},
			}},
			CancelationToken: courier.String("delay-send--user-yes--abc-123"),
		},
		Brand: courier.String("brand"),
		Data: map[string]any{
			"name": "bar",
		},
		Profile: map[string]any{
			"tenant_id": "bar",
		},
		Recipient: courier.String("user-yes"),
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
			Recipient: courier.String("recipient"),
			Brand:     courier.String("brand"),
			Data: map[string]any{
				"foo": "bar",
			},
			Profile: map[string]any{
				"foo": "bar",
			},
			Template: courier.String("template"),
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
