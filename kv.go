package main

import (
	"errors"
	"sync"
)

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStore() (*Store, error) {
	store := &Store{
		data: make(map[string]string),
	}

	return store, nil
}

func (s *Store) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if value, exists := s.data[key]; exists {
		return value, nil
	}

	return "", errors.New("key: %s not found")
}

func (s *Store) Put(key string, value string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return value, nil
}

func (s *Store) Delete(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, err := s.Get(key)
	if err != nil {
		return "", err
	}

	delete(s.data, key)

	return value, nil
}
