package store

import (
	"errors"
	"fmt"
	"sync"
)

type Entry struct {
	value string
}

type Store struct {
	data map[string]*Entry
	mu   sync.RWMutex
}

var store *Store

func Instance() *Store {
	if store == nil {
		store = NewStore()
	}
	return store
}

func NewStore() *Store {
	return &Store{data: make(map[string]*Entry)}
}

func (s *Store) Set(key string, value string) {
	s.mu.Lock()
	s.data[key] = &Entry{value}
	s.mu.Unlock()
}

func (s *Store) Get(key string) (string, error) {
	s.mu.RLock()
	if v, ok := s.data[key]; ok {
		s.mu.RUnlock()
		return v.value, nil
	}
	return "", errors.New(fmt.Sprintf("Key '%s' not found", key))
}
