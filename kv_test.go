package main

import (
	"os"
	"testing"
)

func setup(t *testing.T) *Store {
	f, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatalf("could not create temporary file")
	}

	path := f.Name()

	err = f.Close()
	if err != nil {
		t.Fatalf("could not close temporary file")
	}

	kv, _ := NewStore(path)

	return kv
}

func TestGet(t *testing.T) {
	kv := setup(t)
	_, err := kv.Get("hello")
	if err == nil {
		t.Fatalf("error expected, nil given")
	}
}

func TestPut(t *testing.T) {
	kv := setup(t)

	err := kv.Put("hello", "world")
	if err != nil {
		t.Fatalf("failed to write value to key: %v", err)
	}

	val, _ := kv.Get("hello")
	if val != "world" {
		t.Fatalf("actual: %s does not match expected: %s", val, "world")
	}
}

func TestDelete(t *testing.T) {
	kv := setup(t)

	err := kv.Put("hello", "world")
	if err != nil {
		t.Fatalf("failed to write value to key: %v", err)
	}

	err = kv.Delete("hello")
	if err != nil {
		t.Fatalf("failed to delete key: %v", err)
	}

	_, err = kv.Get("hello")
	if err == nil {
		t.Fatalf("key still exists")
	}
}

func TestDeleteNotFound(t *testing.T) {
	kv := setup(t)

	err := kv.Delete("hello")
	if err == nil {
		t.Fatalf("expected error deleting non-existent key")
	}
}
