package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Storage struct {
	data map[string]int
	mu   sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]int),
	}
}

func (s *Storage) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	return val, ok
}

func (s *Storage) List() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res := make([]string, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}

	return res
}

func (s *Storage) Put(key string, value int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]

	s.data[key] = value
	return ok
}

func (s *Storage) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]

	if !ok {
		return false
	}
	delete(s.data, key)
	return true
}

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
