package storage

func (db *Store) Set(key string, val Value) string {
	db.entries[key] = val
	return "OK"
}

func (db *Store) Get(key string) (Value, bool) {
	value, ok := db.entries[key]
	return value, ok
}

func (db *Store) Delete(key string) bool {

	_, ok := db.entries[key] // since delete is safe even if no key exists
	if ok {
		delete(db.entries, key)
	}
	return ok
}
