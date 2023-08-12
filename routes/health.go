package routes

import (
	"encoding/json"
	"github.com/bijukunjummen/hello-go/types"
	"net/http"
)

type HealthHandler struct {
}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := types.Health{Status: "UP"}
	json.NewEncoder(w).Encode(m)
}
