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

func TestPreferenceSectionTopicNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PreferenceSections.Topics.New(
		context.TODO(),
		"section_id",
		courier.PreferenceSectionTopicNewParams{
			PreferenceTopicCreateRequest: courier.PreferenceTopicCreateRequestParam{
				DefaultStatus:            courier.PreferenceTopicCreateRequestDefaultStatusOptedOut,
				Name:                     "Marketing",
				AllowedPreferences:       []string{"snooze"},
				IncludeUnsubscribeHeader: courier.Bool(true),
				RoutingOptions:           []shared.ChannelClassification{shared.ChannelClassificationDirectMessage},
				TopicData: map[string]any{
					"foo": "bar",
				},
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

func TestPreferenceSectionTopicGet(t *testing.T) {
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
	_, err := client.PreferenceSections.Topics.Get(
		context.TODO(),
		"topic_id",
		courier.PreferenceSectionTopicGetParams{
			SectionID: "section_id",
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

func TestPreferenceSectionTopicList(t *testing.T) {
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
	_, err := client.PreferenceSections.Topics.List(context.TODO(), "section_id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPreferenceSectionTopicArchive(t *testing.T) {
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
	err := client.PreferenceSections.Topics.Archive(
		context.TODO(),
		"topic_id",
		courier.PreferenceSectionTopicArchiveParams{
			SectionID: "section_id",
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

func TestPreferenceSectionTopicReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.PreferenceSections.Topics.Replace(
		context.TODO(),
		"topic_id",
		courier.PreferenceSectionTopicReplaceParams{
			SectionID: "section_id",
			PreferenceTopicReplaceRequest: courier.PreferenceTopicReplaceRequestParam{
				DefaultStatus:            courier.PreferenceTopicReplaceRequestDefaultStatusOptedOut,
				Name:                     "name",
				AllowedPreferences:       []string{"snooze"},
				IncludeUnsubscribeHeader: courier.Bool(true),
				RoutingOptions:           []shared.ChannelClassification{shared.ChannelClassificationDirectMessage},
				TopicData: map[string]any{
					"foo": "bar",
				},
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
