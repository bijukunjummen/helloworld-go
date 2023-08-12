package routes

import (
	"encoding/json"
	"github.com/bijukunjummen/hello-go/types"
	"net/http"
)

type HelloWorldHandler struct {
}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nameValue := r.URL.Query().Get("name")
	if nameValue == "" {
		nameValue = "World"
	}
	m := types.Message{Payload: "Hello " + nameValue}
	json.NewEncoder(w).Encode(m)
}
