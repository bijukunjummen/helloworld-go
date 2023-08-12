package main

import (
	"github.com/bijukunjummen/hello-go/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func appRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", routes.HelloWorldHandler{}).Methods("GET")
	r.Handle("/health", routes.HealthHandler{}).Methods("GET")
	return r
}

func main() {
	r := appRouter()
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	http.ListenAndServe(":"+port, r)
}
