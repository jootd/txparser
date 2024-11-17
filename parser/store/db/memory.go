package db

import "sync"

type MemoryStorage struct {
	mem map[string]string
	mu  sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		mem: make(map[string]string),
	}
}

func (ms *MemoryStorage) Get(key string) (string, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	val, ok := ms.mem[key]
	return val, ok
}

func (ms *MemoryStorage) Set(key, value string) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.mem[key] = value
}
