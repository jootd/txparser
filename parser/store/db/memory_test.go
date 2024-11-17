package db

import (
	"sync"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	db := NewMemoryStorage()
	key := "x"
	expectedValue := "y"
	db.Set(key, expectedValue)
	val, exists := db.Get(key)
	if !exists {
		t.Fatalf("key %s not found", key)
	}

	if val != expectedValue {
		t.Errorf("key %s: got %s, expected %s", key, val, expectedValue)
	}

	// non existing key
	nonExistingKey := "empty"
	val, exists = db.Get(nonExistingKey)
	if exists {
		t.Fatalf("key %s should not exist", nonExistingKey)
	}
	// overwrite
	newValue := "z"
	db.Set(key, newValue)
	val, exists = db.Get(key)
	if !exists {
		t.Fatalf("key %s not found", key)
	}
	if val != newValue {
		t.Errorf("key %s: got %s, expected %s", key, val, newValue)
	}
}

func TestMutex(t *testing.T) {
	db := NewMemoryStorage()
	const key = "key"
	const goroutines = 500
	const value = "value"

	//  write
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			db.Set(key, value)
		}()
	}
	wg.Wait()

	//  read
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			val, exists := db.Get(key)
			if !exists || val != value {
				t.Errorf("exists=%v, value=%s", exists, val)
			}
		}()
	}
	wg.Wait()
}
