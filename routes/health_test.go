package routes_test

import (
	"encoding/json"
	"github.com/bijukunjummen/hello-go/routes"
	"github.com/bijukunjummen/hello-go/types"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	svr := httptest.NewServer(routes.HealthHandler{})
	defer svr.Close()

	resp, err := http.Get(svr.URL)

	if err != nil {
		t.Fatalf("Error in request: %v", err)
	}

	b, _ := io.ReadAll(resp.Body)
	var healthContent types.Health
	json.Unmarshal(b, &healthContent)
	if healthContent.Status != "UP" {
		t.Errorf("Status is not %v", healthContent.Status)
	}
}
