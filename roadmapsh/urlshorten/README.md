# how to run

1. init sql using the sql in ./sql/urlshorten.sql
2. change the mysql config in ./internal/config.go
3. build and run

   ```bash
    go build cmd/url-shorten.go
    ./url-shorten
   ```
