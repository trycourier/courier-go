// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package courier_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/courier-go"
	"github.com/stainless-sdks/courier-go/internal/testutil"
	"github.com/stainless-sdks/courier-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Send.SendMessage(context.TODO(), courier.SendSendMessageParams{
		Message: courier.MessageUnionParam{
			OfContentMessage: &courier.MessageContentMessageParam{
				BaseMessageParam:       courier.BaseMessageParam{},
				BaseMessageSendToParam: courier.BaseMessageSendToParam{},
				Content: courier.ContentUnionParam{
					OfElementalContent: &courier.ContentElementalContentParam{
						Elements: []courier.ElementalNodeUnionParam{{
							OfElementalNodeObject: &courier.ElementalNodeObjectParam{},
						}},
						Version: "version",
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.RequestID)
}
