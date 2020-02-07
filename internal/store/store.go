package store

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Entry struct {
	value   string
	expires int64
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
		if e.expires == -1 || time.Now().Before(time.Unix(0, e.expires)) {
			return e, nil
		}

		s.mu.Lock()
		delete(s.data, key)
		s.mu.Unlock()
	}
	return nil, errors.New(fmt.Sprintf("Key '%s' not found", key))
}
