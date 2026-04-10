package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	storage *Storage
}

func NewServer(storage *Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) handleGetKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	val, ok := s.storage.Get(key)

	if !ok {
		http.Error(w, "Key not in storage", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]int{key: val}); err != nil {
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

	ok := s.storage.Put(req.Key, req.Value)
	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{"status": "updated"}

	if !ok {
		w.WriteHeader(http.StatusCreated)
		resp = map[string]string{"status": "created"}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) handleListKeys(w http.ResponseWriter, r *http.Request) {
	list := s.storage.List()

	w.Header().Set("Content-Type", "application/json")
	resp := map[string][]string{"keys": list}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) handleDeleteKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	if ok := s.storage.Delete(key); !ok {
		http.Error(w, "Key does not exist", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	storage := NewStorage()
	server := NewServer(storage)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /storage/{key}", server.handleGetKey)
	mux.HandleFunc("GET /storage", server.handleListKeys)
	mux.HandleFunc("DELETE /storage/{key}", server.handleDeleteKey)
	mux.HandleFunc("POST /storage", server.handlePutKey)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
