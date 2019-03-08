package services

import (
	"net/http"

	"./welcome"
	"github.com/gorilla/mux"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiV1 := router.PathPrefix("/v1").Subrouter()

	apiV1.HandleFunc("/hello-world", welcome.GetWelcomeHandler).Methods(http.MethodGet)

	return router
}
