package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type CustomTime struct {
	time.Time
}

func (c *CustomTime) MarshalJSON() ([]byte, error) {
	a := struct {
		DayOfWeek  string `json:"day_of_week"`
		DayOfMonth string `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}{
		DayOfWeek:  c.Weekday().String(),
		DayOfMonth: c.Format("02"),
		Month:      c.Month().String(),
		Year:       c.Year(),
		Hour:       c.Hour(),
		Minute:     c.Minute(),
		Second:     c.Second(),
	}

	return json.Marshal(a)
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respData := time.Now().Format(time.RFC3339)
		accept := r.Header.Get("Accept")
		if accept == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			b, err := json.Marshal(&CustomTime{time.Now()})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		} else {
			w.Write([]byte(respData))

		}

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
