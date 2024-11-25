package main

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/harveyTvT/scrath-golang/roadmapsh/cachingproxy/internal"
)

var (
	port   uint64
	origin string
)

type CacheValue struct {
	Method string
	URL    string

	StatusCode int
	Header     http.Header
	Body       []byte
}

func getCacheKey(req *http.Request) string {
	hash := sha256.New()
	hash.Write([]byte(req.Method + "|"))
	hash.Write([]byte(req.URL.String() + "|"))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func proxyHandler(lruCache *internal.LRUCache[string, CacheValue]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// get result from cache
		cacheKey := getCacheKey(req)
		cachedValue, ok := lruCache.Get(cacheKey)
		if ok {
			w.Header().Add("X-Cache", "HIT")
			for k, vs := range cachedValue.Header {
				for _, v := range vs {
					w.Header().Add(k, v)
				}
			}
			w.WriteHeader(cachedValue.StatusCode)
			w.Write([]byte(cachedValue.Body))
			return
		}

		// cache miss
		w.Header().Add("X-Cache", "MISS")

		originURL := fmt.Sprintf("%s%s?%s", origin, req.URL.RawPath, req.URL.RawQuery)
		newReq, err := http.NewRequestWithContext(req.Context(), req.Method, originURL, req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newReq.Header = req.Header

		resp, err := http.DefaultClient.Do(newReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// add to cache
		cacheValue := &CacheValue{
			StatusCode: resp.StatusCode,
			Header:     resp.Header,
			Body:       respBody,
		}

		lruCache.Put(cacheKey, cacheValue)

		for k, vs := range resp.Header {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}

		w.WriteHeader(resp.StatusCode)
		w.Write(respBody)
	})
}

func main() {
	flag.Uint64Var(&port, "port", 0, "the port on which the caching proxy server will run")
	flag.StringVar(&origin, "origin", "", "the URL of the server to which the requests will be forwarded")
	flag.Parse()

	if port == 0 {
		panic(errors.New("must set port"))
	}

	_, err := url.Parse(origin)
	if err != nil {
		panic(err)
	}

	lruCache := internal.NewLRUCache[string, CacheValue](100000)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      proxyHandler(lruCache),
	}

	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
