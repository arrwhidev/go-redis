package store

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Entry struct {
	Value   string
	Expires int64
}

type Store struct {
	data map[string]*Entry // TODO: consider sync.Map
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

func NewEntry(value string, expires int64) *Entry {
	return &Entry{value, expires}
}

func (s *Store) Set(key string, value *Entry) {
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()
}

func (s *Store) Get(key string) (*Entry, error) {
	s.mu.RLock()
	if e, ok := s.data[key]; ok {
		s.mu.RUnlock()
		if e.Expires == -1 || time.Now().Before(time.Unix(0, e.Expires)) {
			return e, nil
		}

		s.mu.Lock()
		delete(s.data, key)
		s.mu.Unlock()
	}
	return nil, errors.New(fmt.Sprintf("Key '%s' not found", key))
}

func (s *Store) Del(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; ok {
		delete(s.data, key)
		return true
	}

	return false
}

func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0, len(s.data))
	for k, v := range s.data {
		if v.Expires == -1 || time.Now().Before(time.Unix(0, v.Expires)) {
			keys = append(keys, k)
		}
	}
	return keys
}
