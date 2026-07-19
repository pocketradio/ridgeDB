package storage

import "time"

func (db *Store) StartCleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {

		<-ticker.C // blocks until a tick is sent thru the chan
		db.mu.Lock()

		for i := range db.entries {

			if db.entries[i].HasExpiry && db.entries[i].ExpiresAt.Before(time.Now()) {
				delete(db.entries, i)
			}
		}
		db.mu.Unlock()

	}
}
