package internal

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

type router struct {
	m   Manager
	mux *http.ServeMux
}

func RegisterMux(mux *http.ServeMux, m Manager) {
	s := &router{
		m:   m,
		mux: mux,
	}

	sv := reflect.ValueOf(s)
	st := reflect.TypeOf(s)

	// register all methods with reflect
	for i := 0; i < st.NumMethod(); i++ {
		method := st.Method(i)
		method.Func.Call([]reflect.Value{sv})
	}
}

func (s *router) Create() {
	s.mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var (
			req CreateReq
			err error
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		record, err := s.m.Create(r.Context(), &req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(record)
	})
}

func (s *router) Get() {
	s.mux.HandleFunc("GET /shorten/{shortCode}", func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.PathValue("shortCode")

		record, err := s.m.Get(r.Context(), shortCode)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Not found"))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(record)

	})
}

func (s *router) Update() {
	s.mux.HandleFunc("PUT /shorten/{shortCode}", func(w http.ResponseWriter, r *http.Request) {
		var req UpdateReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		req.ShortCode = r.PathValue("shortCode")

		record, err := s.m.Update(r.Context(), &req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(record)

	})
}

func (s *router) Delete() {
	s.mux.HandleFunc("DELETE /shorten/{shortCode}", func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.PathValue("shortCode")

		if err := s.m.Delete(r.Context(), shortCode); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Not found"))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)

	})
}

func (s *router) GetStats() {
	s.mux.HandleFunc("GET /shorten/{shortCode}/stats", func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.PathValue("shortCode")

		record, err := s.m.GetStats(r.Context(), shortCode)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Not found"))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(record)
	})
}
