package json_schema

import (
	"net/http"
	"testing"
)

var client *APIClient

func TestMain(m *testing.M) {
	baseURL := "https://httpbin.org"
	client = NewAPIClient(baseURL)

	m.Run()
}

func TestGetMyIP(t *testing.T) {
	client.TestEndpoint(t, "/ip", http.StatusOK, "ip.json")
}
