package courier_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go"
)

func TestClient_GetBrand(t *testing.T) {
	brandID := "VYXZGN7TY94NHDG7A7P9F534DFBK"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands/"+brandID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Id": "%s",
				"Name": "tejas-test",
				"Created": 1606098535741,
				"Published": 1606098535741,
				"Updated": 1606098535741,
				"Settings": {
						"Colors": {
								"Primary": "blue",
								"Secondary": "",
								"Tertiary": ""
						},
						"Email": {
								"Header": null,
								"Footer": {}
						}
				},
				"Snippets": {},
				"Version": "2020-11-23T02:28:55.741Z"
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, brandID)))

		}))
	defer server.Close()

	t.Run("makes request to get a brand", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		rsp, err := client.GetBrand(context.Background(), brandID)
		assert.Nil(t, err)
		assert.Equal(t, brandID, rsp.Id)
	})
}

func TestClient_GetBrands(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Paging": {
						"Cursor": "",
						"More": false
				},
				"Results": [
						{
								"Id": "VYXZGN7TY94NHDG7A7P9F534DFBK",
								"Name": "tejas-test",
								"Created": 1606098535741,
								"Published": 1606098535741,
								"Updated": 1606098535741,
								"Settings": {
										"Colors": {
												"Primary": "blue",
												"Secondary": "",
												"Tertiary": ""
										},
										"Email": {
												"Header": null,
												"Footer": {}
										}
								},
								"Snippets": {},
								"Version": "2020-11-23T02:28:55.741Z"
						}
				]
			}`
			_, _ = rw.Write([]byte(rsp))

		}))
	defer server.Close()

	t.Run("makes request to get all brands", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		rsp, err := client.GetBrands(context.Background(), "")
		assert.Nil(t, err)
		assert.Equal(t, "VYXZGN7TY94NHDG7A7P9F534DFBK", rsp.Results[0].Id)
	})
}

func TestClient_PostBrand(t *testing.T) {
	brandName := "my-brand-new-brand"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"created": 1611967387647,
				"id": "YVXV0VMM6Q4AK3NQ7CWEFGGVM6G4",
				"name": "%s",
				"published": 1611967387647,
				"settings": {
						"colors": {},
						"email": {
								"footer": {
										"markdown": null
								}
						}
				},
				"snippets": {},
				"updated": 1611967387647,
				"version": "2021-01-30T00:43:07.647Z"
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, brandName)))

		}))
	defer server.Close()

	t.Run("makes request to create a brand", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		rsp, err := client.PostBrand(context.Background(), "", brandName, courier.BrandSettings{}, courier.BrandSnippets{}, "")
		assert.Nil(t, err)
		assert.Equal(t, brandName, rsp.Name)
	})
}

func TestClient_PutBrand(t *testing.T) {
	brandID := "VYXZGN7TY94NHDG7A7P9F534DFBK"
	brandUpdatedName := "my-brand-updated"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands/"+brandID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Id": "%s",
				"Name": "%s",
				"Created": 1611970081707,
				"Published": 1611970081707,
				"Updated": 1611970081707,
				"Settings": {
						"colors": {},
						"email": {
								"footer": {
										"markdown": null
								}
						}
				},
				"Snippets": {},
				"Version": "2021-01-30T01:28:01.707Z"
			}`
			_, _ = rw.Write([]byte(fmt.Sprintf(rsp, brandID, brandUpdatedName)))

		}))
	defer server.Close()

	t.Run("makes request to update a brand", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		rsp, err := client.PutBrand(context.Background(), brandID, brandUpdatedName, courier.BrandSettings{}, courier.BrandSnippets{})
		assert.Nil(t, err)
		assert.Equal(t, brandUpdatedName, rsp.Name)
	})
}

func TestClient_DeleteBrand(t *testing.T) {
	brandID := "VYXZGN7TY94NHDG7A7P9F534DFBK"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands/"+brandID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("makes request to delete a brand", func(t *testing.T) {
		client := courier.CourierClient("key", server.URL, "", "")
		client.DeleteBrand(context.Background(), brandID)
	})
}
