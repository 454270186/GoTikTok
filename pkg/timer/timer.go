package timer

import "time"

// Timer for Cache and database scheduled synchronization

func SyncTimer(d time.Duration, task func()) {
	ticker := time.NewTicker(d)

	go func ()  {
		for range ticker.C {
			task()
		}
	}()
}