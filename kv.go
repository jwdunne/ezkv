package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type Operation struct {
	Type      string
	Key       string
	Value     string
	Timestamp int64
}

type Store struct {
	data map[string]string
	mu   sync.RWMutex
	log  *os.File
}

func NewStore(path string) (*Store, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, fmt.Errorf("could not open path: %s", path)
	}

	store := &Store{
		data: make(map[string]string),
		log:  f,
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

func (s *Store) Put(key string, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.writeEntry("put", key, value); err != nil {
		return fmt.Errorf("could not write log entry: %v", err)
	}

	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) error {
	val, err := s.Get(key)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.writeEntry("delete", key, val); err != nil {
		return fmt.Errorf("could not write log entry: %v", err)
	}

	delete(s.data, key)

	return nil
}

func (s *Store) writeEntry(cmd string, key string, value string) error {
	op := Operation{
		Type:      cmd,
		Key:       key,
		Value:     value,
		Timestamp: time.Now().UnixNano(),
	}

	entry, err := json.Marshal(op)
	if err != nil {
		return fmt.Errorf("could not create log entry for %s:%s: %v", key, value, err)
	}

	if _, err := s.log.Write(append(entry, '\n')); err != nil {
		return fmt.Errorf("could not write log entry (%s): %v", entry, err)
	}

	if err := s.log.Sync(); err != nil {
		return fmt.Errorf("could not force write to disk: %v", err)
	}

	return nil
}
