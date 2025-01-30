package main

import (
	"os"
	"testing"
)

func setup(t *testing.T) (*Store, string) {
	f, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatalf("could not create temporary file: %v", err)
	}

	path := f.Name()

	if err = f.Close(); err != nil {
		t.Fatalf("could not close temporary file: %v", err)
	}

	kv, err := NewStore(path)
	if err != nil {
		t.Fatalf("could not create kv store: %v", err)
	}

	return kv, path
}

func TestGet(t *testing.T) {
	kv, _ := setup(t)

	_, err := kv.Get("hello")
	if err == nil {
		t.Fatalf("error expected, nil given")
	}
}

func TestPut(t *testing.T) {
	kv, _ := setup(t)

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
	kv, _ := setup(t)

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
	kv, _ := setup(t)

	err := kv.Delete("hello")
	if err == nil {
		t.Fatalf("expected error deleting non-existent key")
	}
}

func TestRecovery(t *testing.T) {
	kv1, path := setup(t)
	kv1.Put("hello", "world")

	kv2, _ := NewStore(path)

	val, err := kv2.Get("hello")
	if err != nil {
		t.Fatalf("could not get key: %v", err)
	}

	if val != "world" {
		t.Fatalf("unexpected value %s for key hello, expected world", val)
	}
}

func TestConcurrentAccess(t *testing.T) {
	kv, _ := setup(t)

}
