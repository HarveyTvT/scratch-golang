package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/harveytvt/scratch-golang/roadmapsh/urlshorten/internal"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func newBunDB() *bun.DB {
	sqldb, err := sql.Open("mysql", internal.DefaultConfig.MysqlDSN)
	if err != nil {
		panic(err)
	}

	return bun.NewDB(sqldb, mysqldialect.New())
}

func main() {
	bizManager := internal.NewManager(internal.NewRepo(newBunDB()))

	serveMux := http.NewServeMux()
	internal.RegisterMux(serveMux, bizManager)

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", internal.DefaultConfig.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      serveMux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
