package courier_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-go/v2"
)

func TestGetBrands(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"Paging": {
						"Cursor": null,
						"More": false
				},
				"Results": [
						{
								"ID": "VYXZGN7TY94NHDG7A7P9F534DFBK",
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

	t.Run("Get Brands", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetBrands(context.Background(), "")
		assert.Nil(t, err)
		assert.Equal(t, "VYXZGN7TY94NHDG7A7P9F534DFBK", rsp.Results[0].ID)
	})
}

func TestGetBrand(t *testing.T) {
	brandID := "VYXZGN7TY94NHDG7A7P9F534DFBK"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands/"+brandID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"ID": "%s",
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

	t.Run("Get Brand", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.GetBrand(context.Background(), brandID)
		assert.Nil(t, err)
		assert.Equal(t, brandID, rsp.ID)
	})
}

func TestPostBrand(t *testing.T) {
	brandName := "my-brand-new-brand"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands", req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"created": 1611967387647,
				"ID": "YVXV0VMM6Q4AK3NQ7CWEFGGVM6G4",
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

	t.Run("Post Brand", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.PostBrand(context.Background(), courier.PostBrandBody{
			Name: brandName,
		})
		assert.Nil(t, err)
		assert.Equal(t, brandName, rsp.Name)
	})
}

func TestPutBrand(t *testing.T) {
	brandID := "VYXZGN7TY94NHDG7A7P9F534DFBK"
	brandUpdatedName := "my-brand-updated"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands/"+brandID, req.URL.String())

			rw.Header().Add("Content-Type", "application/json")
			rsp := `
			{
				"ID": "%s",
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

	t.Run("Put Brand", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		rsp, err := client.PutBrand(context.Background(), brandID, courier.PutBrandBody{
			Name: brandUpdatedName,
		})
		assert.Nil(t, err)
		assert.Equal(t, brandUpdatedName, rsp.Name)
	})
}

func TestDeleteBrand(t *testing.T) {
	brandID := "VYXZGN7TY94NHDG7A7P9F534DFBK"
	server := httptest.NewServer(http.HandlerFunc(

		func(rw http.ResponseWriter, req *http.Request) {

			assert.Equal(t, "/brands/"+brandID, req.URL.String())

			rw.WriteHeader(http.StatusNoContent)

		}))
	defer server.Close()

	t.Run("Delete brand", func(t *testing.T) {
		client := courier.CreateClient("key", &server.URL)
		client.DeleteBrand(context.Background(), brandID)
	})
}
