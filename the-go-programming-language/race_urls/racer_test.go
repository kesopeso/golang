package race_urls_test

import (
	"net/http"
	"net/http/httptest"
	"race_urls"
	"testing"
	"time"
)

func BenchmarkRaceUrls(b *testing.B) {
	slowServer, fastServer := getSlowAndFastServer()
	defer slowServer.Close()
	defer fastServer.Close()

	b.Run("RaceUrlsFast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			race_urls.RaceUrlsFast(slowServer.URL, fastServer.URL)
		}
	})

	b.Run("RaceUrlsSlow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			race_urls.RaceUrlsSlow(slowServer.URL, fastServer.URL)
		}
	})
}

func TestRaceUrls(t *testing.T) {
	slowServer, fastServer := getSlowAndFastServer()
	defer slowServer.Close()
	defer fastServer.Close()

	want := fastServer.URL
	got := race_urls.RaceUrls(slowServer.URL, fastServer.URL)

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func makeServer(responseDuration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseDuration)
		w.WriteHeader(http.StatusOK)
	}))
}

func getSlowAndFastServer() (*httptest.Server, *httptest.Server) {
	return makeServer(5 * time.Millisecond), makeServer(0 * time.Millisecond)
}