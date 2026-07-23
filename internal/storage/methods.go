package storage

import "time"

func (db *Store) Set(key string, val Value) string {

	db.mu.Lock()
	defer db.mu.Unlock()
	db.entries[key] = val
	return "OK"
}

func (db *Store) Get(key string) (Value, bool) {

	db.mu.Lock()
	defer db.mu.Unlock()

	value, ok := db.entries[key]

	if ok {
		if value.HasExpiry && !value.ExpiresAt.After(time.Now()) {

			ok = false
			delete(db.entries, key)
			return Value{}, ok
		}

		return value, ok
	}

	return Value{}, ok
}

func (db *Store) Delete(key string) bool {

	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.entries[key] // since delete is safe even if no key exists
	if ok {
		delete(db.entries, key)
	}
	return ok
}
