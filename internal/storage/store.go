package storage

type Value struct {
	Data string
	TTL  int64
}

type Store struct {
	entries map[string]Value
}

func NewStore() *Store {
	var s Store
	s.entries = make(map[string]Value)
	return &s
}
