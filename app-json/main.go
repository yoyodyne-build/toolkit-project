package main

import (
	"github.com/yoyodyne-build/toolkit"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ResponsePayload struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
}

func main() {
	mux := routes()

	log.Println("starting server at http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/receive-post", receivePost)
	mux.HandleFunc("/remote-service", remoteService)
	mux.HandleFunc("/simulated-service", simulatedService)

	return mux
}

func simulatedService(w http.ResponseWriter, r *http.Request) {
	payload := ResponsePayload{
		Message: "simulated service",
	}

	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, payload)
}

func receivePost(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message: "hanlder received post",
	}

	err = t.WriteJSON(w, http.StatusOK, responsePayload)
	if err != nil {
		log.Println(err)
	}
}

func remoteService(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}

	_, statusCode, err := t.PostJSONToRemote("http://localhost:8080/simulated-service", requestPayload)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message:    "hanlder received post",
		StatusCode: statusCode,
	}

	err = t.WriteJSON(w, http.StatusOK, responsePayload)
	if err != nil {
		log.Println(err)
	}
}
