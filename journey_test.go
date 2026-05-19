// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/internal/testutil"
	"github.com/trycourier/courier-go/v4/option"
)

func TestJourneyNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Journeys.New(context.TODO(), courier.JourneyNewParams{
		CreateJourneyRequest: courier.CreateJourneyRequestParam{
			Name: "Welcome Journey",
			Nodes: []courier.JourneyNodeUnionParam{{
				OfAPIInvokeTrigger: &courier.JourneyAPIInvokeTriggerNodeParam{
					TriggerType: courier.JourneyAPIInvokeTriggerNodeTriggerTypeAPIInvoke,
					Type:        courier.JourneyAPIInvokeTriggerNodeTypeTrigger,
					ID:          courier.String("trigger-1"),
					Conditions: courier.JourneyConditionsFieldUnionParam{
						OfSingleCondition: courier.JourneyConditionAtom{"string", "string"},
					},
					Schema: map[string]any{
						"foo": "bar",
					},
				},
			}, {
				OfAPIInvokeTrigger: &courier.JourneyAPIInvokeTriggerNodeParam{
					TriggerType: courier.JourneyAPIInvokeTriggerNodeTriggerTypeAPIInvoke,
					Type:        courier.JourneyAPIInvokeTriggerNodeTypeTrigger,
					ID:          courier.String("send-1"),
					Conditions: courier.JourneyConditionsFieldUnionParam{
						OfSingleCondition: courier.JourneyConditionAtom{"string", "string"},
					},
					Schema: map[string]any{
						"foo": "bar",
					},
				},
			}},
			Enabled: courier.Bool(true),
			State:   courier.JourneyStateDraft,
		},
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJourneyGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Journeys.Get(
		context.TODO(),
		"x",
		courier.JourneyGetParams{
			Version: courier.String("published"),
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

func TestJourneyListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Journeys.List(context.TODO(), courier.JourneyListParams{
		Cursor:  courier.String("cursor"),
		Version: courier.JourneyListParamsVersionPublished,
	})
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJourneyArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	err := client.Journeys.Archive(context.TODO(), "x")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJourneyInvokeWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Journeys.Invoke(
		context.TODO(),
		"templateId",
		courier.JourneyInvokeParams{
			JourneysInvokeRequest: courier.JourneysInvokeRequestParam{
				Data: map[string]any{
					"order_id": "bar",
					"amount":   "bar",
				},
				Profile: map[string]any{
					"foo": "bar",
				},
				UserID: courier.String("user-123"),
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

func TestJourneyListVersions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Journeys.ListVersions(context.TODO(), "x")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJourneyPublishWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Journeys.Publish(
		context.TODO(),
		"x",
		courier.JourneyPublishParams{
			JourneyPublishRequest: courier.JourneyPublishRequestParam{
				Version: courier.String("v321669910225"),
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

func TestJourneyReplaceWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Journeys.Replace(
		context.TODO(),
		"x",
		courier.JourneyReplaceParams{
			CreateJourneyRequest: courier.CreateJourneyRequestParam{
				Name: "Welcome Journey v2",
				Nodes: []courier.JourneyNodeUnionParam{{
					OfAPIInvokeTrigger: &courier.JourneyAPIInvokeTriggerNodeParam{
						TriggerType: courier.JourneyAPIInvokeTriggerNodeTriggerTypeAPIInvoke,
						Type:        courier.JourneyAPIInvokeTriggerNodeTypeTrigger,
						ID:          courier.String("x"),
						Conditions: courier.JourneyConditionsFieldUnionParam{
							OfSingleCondition: courier.JourneyConditionAtom{"string", "string"},
						},
						Schema: map[string]any{
							"foo": "bar",
						},
					},
				}},
				Enabled: courier.Bool(true),
				State:   courier.JourneyStateDraft,
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
