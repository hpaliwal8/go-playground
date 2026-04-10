package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type result struct {
	intResult  int
	boolResult bool
	listResult []string
	ok         bool
}

type command struct {
	key       string
	value     int
	operation string
	res       chan result
}

type Server struct {
	ch chan command
}

func NewServer() *Server {
	return &Server{
		ch: make(chan command),
	}
}

func (s *Server) handleGetKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	result := make(chan result)
	s.ch <- command{key: key, operation: "get", res: result}
	resp := <-result

	if !resp.ok {
		http.Error(w, "Key not in storage", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]int{key: resp.intResult}); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

type createKeyValue struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func (s *Server) handlePutKey(w http.ResponseWriter, r *http.Request) {
	var req createKeyValue
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	result := make(chan result)
	s.ch <- command{key: req.Key, value: req.Value, operation: "post", res: result}
	res := <-result

	if !res.ok {
		http.Error(w, "Error occured", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{"status": "updated"}

	if !res.boolResult {
		w.WriteHeader(http.StatusCreated)
		resp = map[string]string{"status": "created"}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) handleListKeys(w http.ResponseWriter, r *http.Request) {
	result := make(chan result)
	s.ch <- command{operation: "list", res: result}
	res := <-result

	w.Header().Set("Content-Type", "application/json")
	resp := map[string][]string{"keys": res.listResult}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) handleDeleteKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	result := make(chan result)
	s.ch <- command{key: key, operation: "delete", res: result}
	res := <-result

	if ok := res.boolResult; !ok {
		http.Error(w, "Key does not exist", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	server := NewServer()
	go storageManager(server.ch)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /storage/{key}", server.handleGetKey)
	mux.HandleFunc("GET /storage", server.handleListKeys)
	mux.HandleFunc("DELETE /storage/{key}", server.handleDeleteKey)
	mux.HandleFunc("POST /storage", server.handlePutKey)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
