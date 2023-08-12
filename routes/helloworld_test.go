package routes

import (
	"encoding/json"
	"github.com/bijukunjummen/hello-go/types"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldEndpoint(t *testing.T) {
	svr := httptest.NewServer(HelloWorldHandler{})

	defer svr.Close()

	resp, _ := http.Get(svr.URL)

	b, _ := io.ReadAll(resp.Body)
	var content types.Message
	json.Unmarshal(b, content)

	if !(content.Payload != "Hello World") {
		t.Errorf("Bad Response content: %v", content)
	}
}
