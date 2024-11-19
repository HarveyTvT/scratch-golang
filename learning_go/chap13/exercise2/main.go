package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respData := time.Now().Format(time.RFC3339)
		w.Write([]byte(respData))
	})

	wrapperedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mySlog := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{}))
		mySlog.Info("Request received", "ip", r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      wrapperedHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
