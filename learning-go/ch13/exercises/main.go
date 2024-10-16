package main

import (
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	logRequestReceivedMsg       string = "request received"
	logRequestReceivedIp        string = "ip"
	logRemoteIpNotAvailable     string = "not available"
	logStructToJsonFailedMsg    string = "unable to create json from struct"
	logStructToJsonFailedStruct string = "struct"
	logRequestDurationMsg       string = "request duration"
	logRequestDurationKey       string = "elapsed"
)

type DateResponse struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Year       int    `json:"year"`
	Month      string `json:"month"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

type ServerConfig struct {
	logger *slog.Logger
}

type RouteConfig struct {
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

func main() {
	logger := NewLogger()
	server := NewServer(logger)
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func NewServer(logger *slog.Logger) *http.Server {
	serverConfig := ServerConfig{
		logger: logger,
	}
	return serverConfig.createServer()

}

func NewLogger() *slog.Logger {
	loggerOptions := &slog.HandlerOptions{Level: slog.LevelDebug}
	loggerHandler := slog.NewJSONHandler(os.Stdout, loggerOptions)
	logger := slog.New(loggerHandler)
	return logger
}

func getDateResponseFromTime(t time.Time) DateResponse {
	dayOfWeek := t.Weekday()
	year, month, dayOfMonth := t.Date()
	hour, minute, second := t.Clock()

	return DateResponse{
		DayOfWeek:  dayOfWeek.String(),
		DayOfMonth: dayOfMonth,
		Year:       year,
		Month:      month.String(),
		Hour:       hour,
		Minute:     minute,
		Second:     second,
	}
}

func (s *ServerConfig) nowRoute(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now()

	var responseContentType string
	var responseBody []byte

	switch r.Header.Get("Accept") {
	case "application/json":
		responseContentType = "application/json"
		dateResponse := getDateResponseFromTime(timeNow)
		dateResponseJson, err := json.Marshal(dateResponse)
		if err != nil {
			dateResponseJson = []byte("")
			s.logger.Error(logStructToJsonFailedMsg, logStructToJsonFailedStruct, dateResponse)
		}
		responseBody = dateResponseJson
	default:
		responseContentType = "text/plain"
		responseBody = []byte(timeNow.Format(time.RFC3339))
	}

	w.Header().Set("Content-Type", responseContentType)
	w.Write(responseBody)
}

func (s *ServerConfig) getRoutes() []RouteConfig {
	routes := []RouteConfig{
		{
			path:    "GET /now",
			handler: s.nowRoute,
		},
	}
	return routes
}

func (s *ServerConfig) getMiddlerware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				s.logger.Debug(logRequestReceivedMsg, logRequestReceivedIp, logRemoteIpNotAvailable)
			} else {
				s.logger.Debug(logRequestReceivedMsg, logRequestReceivedIp, ip)
			}

			startTime := time.Now()

			h.ServeHTTP(w, r)

			duration := time.Since(startTime)
			s.logger.Debug(logRequestDurationMsg, logRequestDurationKey, duration)
		})
	}
}

func (s *ServerConfig) getRouter() *http.ServeMux {
	router := http.NewServeMux()
	for _, route := range s.getRoutes() {
		router.HandleFunc(route.path, route.handler)
	}
	return router
}

func (s *ServerConfig) getHandler() http.Handler {
	middleware := s.getMiddlerware()
	router := s.getRouter()
	handler := middleware(router)
	return handler
}

func (s *ServerConfig) createServer() *http.Server {
	return &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Minute,
		Handler:      s.getHandler(),
	}
}
