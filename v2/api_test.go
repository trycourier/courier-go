package courier

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCourier_ExecuteRequest(t *testing.T) {
	t.Parallel()
	want := []byte("Ok")
	server := mockServer(want)
	defer server.Close()
	api := &APIConfiguration{
		AuthToken: "",
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", server.URL, http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	got, err := api.ExecuteRequest(req)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(string(want), string(got)))
	}
}

func mockServer(want []byte) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(want)
	}))
	return server
}
