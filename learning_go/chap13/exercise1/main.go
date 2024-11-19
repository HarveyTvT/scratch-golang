package main

import (
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respData := time.Now().Format(time.RFC3339)
		w.Write([]byte(respData))
	})

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
