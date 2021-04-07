package courier_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go/v2"
)

func TestInvokeAutomation(t *testing.T) {
	type AutomationStep struct {
		Action string
	}
	type Automation struct {
		Steps []AutomationStep
	}
	type Data struct {
		Foo string
	}
	type Profile struct {
		Email string
	}

	type RequestBody struct {
		Automation Automation
		Brand      string
		Data       Data
		Profile    Profile
		Recipient  string
		Template   string
	}

	var requestBody RequestBody

	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/automations/invoke", req.URL.String())

			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"runId\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}

		}))
	defer server.Close()

	t.Run("invokes automation", func(t *testing.T) {

		client := courier.CreateClient("key", &server.URL)

		brand := "brand"
		data := &Data{
			Foo: "bar",
		}
		profile := &Profile{
			Email: "foo@bar.com",
		}
		recipient := "recpient"
		template := "template"

		runID, err := client.InvokeAutomation(context.Background(), courier.AutomationInvokeBody{
			Automation: courier.Automation{
				Steps: []courier.AutomationStep{
					courier.AutomationStep{
						Action: "send",
					},
				},
			},
			Brand:     brand,
			Data:      data,
			Profile:   profile,
			Recipient: recipient,
			Template:  template,
		})

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, runID)
		assert.Equal(t, 1, len(requestBody.Automation.Steps))
		assert.Equal(t, "send", requestBody.Automation.Steps[0].Action)
		assert.Equal(t, brand, requestBody.Brand)
		assert.Equal(t, data.Foo, requestBody.Data.Foo)
		assert.Equal(t, profile.Email, requestBody.Profile.Email)
		assert.Equal(t, recipient, requestBody.Recipient)
		assert.Equal(t, template, requestBody.Template)
	})
}

func TestInvokeAutomationTemplate(t *testing.T) {
	type Data struct {
		Foo string
	}
	type Profile struct {
		Email string
	}

	type RequestBody struct {
		Brand     string
		Data      Data
		Profile   Profile
		Recipient string
		Template  string
	}
	templateId := "template001"

	var requestBody RequestBody

	expectedResponseID := "123456789"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/automations/"+templateId+"/invoke", req.URL.String())

			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&requestBody)
			if err != nil {
				t.Error(err)
			}

			rw.Header().Add("Content-Type", "application/json")
			_, writeErr := rw.Write([]byte(fmt.Sprintf("{ \"runId\" : \"%s\" }", expectedResponseID)))
			if writeErr != nil {
				t.Error(writeErr)
			}

		}))
	defer server.Close()

	t.Run("invokes automation template", func(t *testing.T) {

		client := courier.CreateClient("key", &server.URL)

		brand := "brand"
		data := &Data{
			Foo: "bar",
		}
		profile := &Profile{
			Email: "foo@bar.com",
		}
		recipient := "recpient"
		template := "template"

		runID, err := client.InvokeAutomationTemplate(context.Background(), templateId, courier.AutomationTemplateInvokeBody{
			Brand:     brand,
			Data:      data,
			Profile:   profile,
			Recipient: recipient,
			Template:  template,
		})

		assert.Nil(t, err)
		assert.Equal(t, expectedResponseID, runID)
		assert.Equal(t, brand, requestBody.Brand)
		assert.Equal(t, data.Foo, requestBody.Data.Foo)
		assert.Equal(t, profile.Email, requestBody.Profile.Email)
		assert.Equal(t, recipient, requestBody.Recipient)
		assert.Equal(t, template, requestBody.Template)
	})
}
