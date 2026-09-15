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

func TestWorkspacePreferenceTopicNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkspacePreferences.Topics.New(
		context.TODO(),
		"section_id",
		courier.WorkspacePreferenceTopicNewParams{
			WorkspacePreferenceTopicCreateRequest: courier.WorkspacePreferenceTopicCreateRequestParam{
				DefaultStatus:      courier.WorkspacePreferenceTopicCreateRequestDefaultStatusOptedOut,
				Name:               "Marketing",
				AllowedPreferences: []string{"snooze"},
				Description:        courier.String("description"),
				Digest: courier.TopicDigestRequestParam{
					Schedules: []courier.TopicDigestScheduleRequestParam{{
						Frequency:  courier.DigestFrequencyInstant,
						DayOfMonth: courier.Int(1),
						DayOfWeek:  courier.DigestDayOfWeekSunday,
						DaysOfWeek: []courier.DigestDayOfWeek{courier.DigestDayOfWeekSunday},
						Disabled:   courier.Bool(true),
						IsDefault:  courier.Bool(true),
						ScheduleID: courier.String("schedule_id"),
						Time:       courier.String("time"),
						Timezone:   courier.String("timezone"),
					}},
					TemplateID: "template_id",
					AudienceID: courier.String("audience_id"),
					Categories: []courier.TopicDigestCategoryParam{{
						CategoryKey: "category_key",
						Limit:       courier.Int(1),
						Retain:      courier.TopicDigestCategoryRetainFirst,
						SortKey:     courier.String("sort_key"),
					}},
					TriggerEmpty: courier.Bool(true),
				},
				IncludeUnsubscribeHeader: courier.Bool(true),
				RoutingOptions:           []shared.ChannelClassification{shared.ChannelClassificationDirectMessage},
				TopicData: map[string]any{
					"foo": "bar",
				},
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

func TestWorkspacePreferenceTopicGet(t *testing.T) {
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
	_, err := client.WorkspacePreferences.Topics.Get(
		context.TODO(),
		"topic_id",
		courier.WorkspacePreferenceTopicGetParams{
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

func TestWorkspacePreferenceTopicList(t *testing.T) {
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
	_, err := client.WorkspacePreferences.Topics.List(context.TODO(), "section_id")
	if err != nil {
		var apierr *courier.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkspacePreferenceTopicArchive(t *testing.T) {
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
	err := client.WorkspacePreferences.Topics.Archive(
		context.TODO(),
		"topic_id",
		courier.WorkspacePreferenceTopicArchiveParams{
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

func TestWorkspacePreferenceTopicDeleteDigest(t *testing.T) {
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
	err := client.WorkspacePreferences.Topics.DeleteDigest(
		context.TODO(),
		"topic_id",
		courier.WorkspacePreferenceTopicDeleteDigestParams{
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

func TestWorkspacePreferenceTopicReleaseDigestWithOptionalParams(t *testing.T) {
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
	err := client.WorkspacePreferences.Topics.ReleaseDigest(
		context.TODO(),
		"topic_id",
		courier.WorkspacePreferenceTopicReleaseDigestParams{
			SectionID: "section_id",
			TopicDigestReleaseRequest: courier.TopicDigestReleaseRequestParam{
				UserID:   "user_01h1p2c3d4e5f6g7h8",
				TenantID: courier.String("x"),
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

func TestWorkspacePreferenceTopicReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkspacePreferences.Topics.Replace(
		context.TODO(),
		"topic_id",
		courier.WorkspacePreferenceTopicReplaceParams{
			SectionID: "section_id",
			WorkspacePreferenceTopicReplaceRequest: courier.WorkspacePreferenceTopicReplaceRequestParam{
				DefaultStatus:      courier.WorkspacePreferenceTopicReplaceRequestDefaultStatusOptedIn,
				Name:               "Product Updates",
				AllowedPreferences: []string{"channel_preferences"},
				Description:        courier.String("description"),
				Digest: courier.TopicDigestRequestParam{
					Schedules: []courier.TopicDigestScheduleRequestParam{{
						Frequency:  courier.DigestFrequencyInstant,
						DayOfMonth: courier.Int(1),
						DayOfWeek:  courier.DigestDayOfWeekSunday,
						DaysOfWeek: []courier.DigestDayOfWeek{courier.DigestDayOfWeekSunday},
						Disabled:   courier.Bool(true),
						IsDefault:  courier.Bool(true),
						ScheduleID: courier.String("schedule_id"),
						Time:       courier.String("time"),
						Timezone:   courier.String("timezone"),
					}},
					TemplateID: "template_id",
					AudienceID: courier.String("audience_id"),
					Categories: []courier.TopicDigestCategoryParam{{
						CategoryKey: "category_key",
						Limit:       courier.Int(1),
						Retain:      courier.TopicDigestCategoryRetainFirst,
						SortKey:     courier.String("sort_key"),
					}},
					TriggerEmpty: courier.Bool(true),
				},
				IncludeUnsubscribeHeader: courier.Bool(true),
				RoutingOptions:           []shared.ChannelClassification{shared.ChannelClassificationEmail, shared.ChannelClassificationInbox},
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
