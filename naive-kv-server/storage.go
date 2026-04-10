package main

import (
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
