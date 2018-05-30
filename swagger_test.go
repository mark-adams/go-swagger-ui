package swaggerui_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	swaggerui "github.com/mark-adams/go-swagger-ui"
)

func TestSwaggerUIHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/docs", swaggerui.Handler())

	server := httptest.NewServer(mux)
	resp, err := http.Get(server.URL + "/docs")
	if err != nil {
		t.Fatalf("error loading docs URL: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected response code: %d", resp.StatusCode)
	}

}
