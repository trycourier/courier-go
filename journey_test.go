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
				OfSend: &courier.JourneySendNodeParam{
					Message: courier.JourneySendNodeMessageParam{
						Context: courier.JourneySendNodeMessageContextParam{
							TenantID: "x",
						},
						Data: map[string]any{
							"foo": "bar",
						},
						Delay: courier.JourneySendNodeMessageDelayParam{
							Until:    "x",
							Timezone: courier.String("x"),
						},
						Template: courier.String("nt_01kx4h2jdafq8bk9aftxak4b40"),
						To: courier.JourneySendNodeMessageToParam{
							EmailOverride: courier.String("x"),
							MsTeams: courier.JourneySendNodeToMsTeamsParam{
								ChannelID:   courier.String("x"),
								ChannelName: courier.String("x"),
								Email:       courier.String("x"),
								ServiceURL:  courier.String("x"),
								TeamID:      courier.String("x"),
								TenantID:    courier.String("x"),
								UserID:      courier.String("x"),
							},
							PhoneNumberOverride: courier.String("x"),
							Slack: courier.JourneySendNodeToSlackUnionParam{
								OfJourneySendNodeToSlackChannel: &courier.JourneySendNodeToSlackChannelParam{
									Channel:     "x",
									AccessToken: courier.String("x"),
								},
							},
							UserIDOverride: courier.String("x"),
						},
					},
					Type: courier.JourneySendNodeTypeSend,
					ID:   courier.String("send-1"),
					Conditions: courier.JourneyConditionsFieldUnionParam{
						OfSingleCondition: courier.JourneyConditionAtom{"string", "string"},
					},
					Experiment: courier.JourneyExperimentParam{
						BucketingKey: "x",
						Variants: []courier.JourneyExperimentVariantParam{{
							ID:         "x",
							TemplateID: "x",
							Weight:     0,
							Name:       courier.String("name"),
						}, {
							ID:         "x",
							TemplateID: "x",
							Weight:     0,
							Name:       courier.String("name"),
						}},
						ID:   courier.String("x"),
						Name: courier.String("name"),
					},
				},
			}, {
				OfExit: &courier.JourneyExitNodeParam{
					Type: courier.JourneyExitNodeTypeExit,
					ID:   courier.String("exit-1"),
				},
			}},
			Enabled: courier.Bool(true),
			State:   courier.JourneyStateDraft,
		},
		IdempotencyKey:         courier.String("order-ORD-456-user-123"),
		XIdempotencyExpiration: courier.String("1785312000"),
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

func TestJourneyCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Journeys.Cancel(context.TODO(), courier.JourneyCancelParams{
		CancelJourneyRequest: courier.CancelJourneyRequestUnionParam{
			OfByCancelationToken: &courier.CancelJourneyRequestByCancelationTokenParam{
				CancelationToken: "order-1234",
			},
		},
		IdempotencyKey:         courier.String("order-ORD-456-user-123"),
		XIdempotencyExpiration: courier.String("1785312000"),
	})
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
			IdempotencyKey:         courier.String("order-ORD-456-user-123"),
			XIdempotencyExpiration: courier.String("1785312000"),
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
			IdempotencyKey:         courier.String("order-ORD-456-user-123"),
			XIdempotencyExpiration: courier.String("1785312000"),
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
