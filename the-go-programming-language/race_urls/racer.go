package race_urls

import (
	"net/http"
	"time"
)

func RaceUrlsFast(url1, url2 string) string {
	select {
	case <-pingUrl(url1):
		return url1
	case <-pingUrl(url2):
		return url2
	}
}

func RaceUrlsSlow(url1, url2 string) string {
	start := time.Now()
	http.Get(url1)
	duration1 := time.Since(start)

	start = time.Now()
	http.Get(url2)
	duration2 := time.Since(start)

	if duration1 < duration2 {
		return url1
	}
	return url2
}

func pingUrl(url string) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
