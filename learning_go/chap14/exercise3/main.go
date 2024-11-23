package main

import (
	"context"
	"fmt"
	"net/http"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, "level", level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	v, ok := ctx.Value("level").(Level)
	return v, ok
}

func LogLevelHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if logLevel := req.URL.Query().Get("log_level"); logLevel == string(Debug) || logLevel == string(Info) {
			ctx := ContextWithLevel(req.Context(), Level(logLevel))
			req = req.WithContext(ctx)
		}

		h.ServeHTTP(w, req)
	})
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	if v, ok := LevelFromContext(ctx); ok {
		inLevel = v
	}

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}

}
