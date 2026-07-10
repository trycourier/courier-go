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
	"github.com/trycourier/courier-go/v4/shared"
)

func TestUserPreferenceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Preferences.Get(
		context.TODO(),
		"user_id",
		courier.UserPreferenceGetParams{
			TenantID: courier.String("tenant_id"),
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

func TestUserPreferenceBulkReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Preferences.BulkReplace(
		context.TODO(),
		"user_id",
		courier.UserPreferenceBulkReplaceParams{
			Topics: []courier.UserPreferenceBulkReplaceParamsTopic{{
				Status:           "OPTED_IN",
				TopicID:          "pt_01kx4h2jdafq8bk996nn92357r",
				CustomRouting:    []shared.ChannelClassification{shared.ChannelClassificationInbox, shared.ChannelClassificationEmail},
				HasCustomRouting: courier.Bool(true),
			}},
			TenantID: courier.String("tenant_id"),
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

func TestUserPreferenceBulkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Preferences.BulkUpdate(
		context.TODO(),
		"user_id",
		courier.UserPreferenceBulkUpdateParams{
			Topics: []courier.UserPreferenceBulkUpdateParamsTopic{{
				Status:           "OPTED_IN",
				TopicID:          "pt_01kx4h2jdafq8bk996nn92357r",
				CustomRouting:    []shared.ChannelClassification{shared.ChannelClassificationInbox, shared.ChannelClassificationEmail},
				HasCustomRouting: courier.Bool(true),
			}, {
				Status:           "OPTED_OUT",
				TopicID:          "pt_01kx4h2jdafq8bk99eyt3dx43x",
				CustomRouting:    []shared.ChannelClassification{shared.ChannelClassificationDirectMessage},
				HasCustomRouting: courier.Bool(true),
			}},
			TenantID: courier.String("tenant_id"),
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

func TestUserPreferenceDeleteTopicWithOptionalParams(t *testing.T) {
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
	err := client.Users.Preferences.DeleteTopic(
		context.TODO(),
		"topic_id",
		courier.UserPreferenceDeleteTopicParams{
			UserID:   "user_id",
			TenantID: courier.String("tenant_id"),
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

func TestUserPreferenceGetTopicWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Preferences.GetTopic(
		context.TODO(),
		"topic_id",
		courier.UserPreferenceGetTopicParams{
			UserID:   "user_id",
			TenantID: courier.String("tenant_id"),
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

func TestUserPreferenceUpdateOrNewTopicWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Preferences.UpdateOrNewTopic(
		context.TODO(),
		"topic_id",
		courier.UserPreferenceUpdateOrNewTopicParams{
			UserID: "user_id",
			Topic: courier.UserPreferenceUpdateOrNewTopicParamsTopic{
				Status:           shared.PreferenceStatusOptedIn,
				CustomRouting:    []shared.ChannelClassification{shared.ChannelClassificationInbox, shared.ChannelClassificationEmail},
				HasCustomRouting: courier.Bool(true),
			},
			TenantID: courier.String("tenant_id"),
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
