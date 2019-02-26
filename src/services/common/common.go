package common

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorMessage contains the error
type ErrorMessage struct {
	Message string `json:"error"`
}

// RenderStatusOK is used for rendering JSON response body with appropriate headers
func RenderStatusOK(w http.ResponseWriter, r *http.Request, response interface{}) {
	renderJSON(w, r, http.StatusOK, response)
}

// SendInternalServerError sends Internal Server Error Status and logs an error if it exists
func SendInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	renderError(w, r, http.StatusInternalServerError, "", err)
}

func renderJSON(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf(`"%s %s" err: %s`, r.Method, r.URL, err)
		return
	}

	render(w, status, data)
}

func renderError(w http.ResponseWriter, r *http.Request, status int, message string, errMessage error) {
	var data []byte
	var err error

	if message != "" {
		data, err = json.Marshal(ErrorMessage{Message: message})
		if err != nil {
			SendInternalServerError(w, r, err)
			return
		}
	}

	if errMessage != nil {
		log.Printf(`"%s %s" err: %s`, r.Method, r.URL, errMessage)
	}

	render(w, status, data)
}

func render(w http.ResponseWriter, status int, response []byte) {
	w.WriteHeader(status)
	if response == nil {
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(response)
	if err != nil {
		log.Printf("Cannot write response: %s", err)
	}
}
