package storage

import "time"

type Value struct {
	Data      string
	ExpiresAt time.Time
}

type Store struct {
	entries map[string]Value
}

func NewStore() *Store {
	var s Store
	s.entries = make(map[string]Value)
	return &s
}
