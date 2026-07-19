package storage

import (
	"sync"
	"time"
)

type Value struct {
	Data      string
	ExpiresAt time.Time
	HasExpiry bool
}

type Store struct {
	entries map[string]Value
	mu      sync.RWMutex
}

func NewStore() *Store {
	var s Store
	s.entries = make(map[string]Value)
	return &s
}
