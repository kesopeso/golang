package race

import (
	"sync"
)

func Race(useLock bool) int {
	counter := 0
	var l sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				if useLock {
					l.Lock()
				}
				counter++
				if useLock {
					l.Unlock()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}
