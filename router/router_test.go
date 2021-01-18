package router

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRouter(t *testing.T) {
	// Use default config file
	os.Setenv("CONFIG_PATH", "../config.yaml")

	ts := httptest.NewServer(InitRouter())
	defer ts.Close()

	// Test /healthz endpoint
	r, _ := http.Get(ts.URL + "/healthz")
	if r.StatusCode != http.StatusOK {
		t.Errorf("Failed to call /healthz endpoint. Status code: %d)", r.StatusCode)
	}

	// Test /v1/fizzbuzz endpoint
	r, _ = http.Get(ts.URL + "/v1/fizzbuzz?int1=3&int2=5&limit=16&str1=fizz&str2=buzz")
	if r.StatusCode != http.StatusOK {
		t.Errorf("Failed to call /v1/fizzbuzz endpoint. Status code: %d)", r.StatusCode)
	}
}
