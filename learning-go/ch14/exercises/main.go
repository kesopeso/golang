package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//exercise2()
	runServer()
}

func runServer() {
	timeoutMiddleware := exercise1_middleware(2000)
	mux := http.NewServeMux()
	mux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		iterations := 0
		sum := 0

		for {
			select {
			case <-ctx.Done():
				result := fmt.Sprintf("iterations: %d, sum: %d, err: %v", iterations, sum, ctx.Err())
				w.WriteHeader(http.StatusRequestTimeout)
				w.Write([]byte(result))
				return
			default:
			}

			num := rand.Intn(100000000)
			if num == 1234 {
				result := fmt.Sprintf("iterations: %d, sum: %d, err: 1234 generated", iterations, sum)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(result))
				return
			}

			sum += num
			iterations++
		}
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: exercise3_middleware(timeoutMiddleware(mux)),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server listen and serve error occured", err)
		return
	}
}

func exercise2() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	iterations := 0
	sum := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("iterations: %d, sum: %d, err: %v\n", iterations, sum, ctx.Err())
			return
		default:
		}

		num := rand.Intn(100000000)
		if num == 1234 {
			fmt.Printf("iterations: %d, sum: %d, err: 1234 generated\n", iterations, sum)
			return
		}

		sum += num
		iterations++
	}
}

func exercise1_middleware(timeoutMs int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeoutCause(ctx, time.Duration(timeoutMs)*time.Millisecond, context.DeadlineExceeded)
			defer cancel()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	None  Level = ""
)

type logLevel int

const logLevelKey logLevel = 0

func setLogLevel(ctx context.Context, l Level) context.Context {
	return context.WithValue(ctx, logLevelKey, l)
}

func getLogLevel(ctx context.Context) Level {
	logLevel, ok := ctx.Value(logLevelKey).(Level)
	if !ok {
		return None
	}
	return logLevel
}

func exercise3_middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		debugLevel := r.URL.Query().Get("log_level")
		ctx := setLogLevel(r.Context(), Level(debugLevel))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
		Log(r.Context(), Info, "logging info level message")
		Log(r.Context(), Debug, "logging debug level message")
	})
}

func Log(ctx context.Context, level Level, message string) {
	inLevel := getLogLevel(ctx)
	if level == Debug && inLevel == Debug {
		fmt.Println("lvl: debug", message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println("lvl: info", message)
	}
}
