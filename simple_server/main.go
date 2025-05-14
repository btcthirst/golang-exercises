package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	// Add any necessary fields here
}

func main() {
	h := &Handler{}
	server(h)
}

func server(h *Handler) {
	http.HandleFunc("/", h.startHandler)
	http.HandleFunc("/hello", h.helloHandler)
	http.HandleFunc("/time", h.timeHandler)
	http.ListenAndServe(":8080", nil)
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) startHandler(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{"message": "Server started successfully!"}
	JSONResponse(w, message)
}

func (h *Handler) helloHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := map[string]string{"message": "Hello, User!"}
	JSONResponse(w, jsonData)
}

func (h *Handler) timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	jsonData := map[string]string{"current_time": now.Format(time.RFC1123)}
	JSONResponse(w, jsonData)
}
