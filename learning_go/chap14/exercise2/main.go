package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"
)

func monkey() {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 2*time.Second, errors.New("timeout"))
	defer cancel()

	var cnt int
	for {
		if err := context.Cause(ctx); err != nil {
			fmt.Printf("timeout err: %s, cnt: %d", err.Error(), cnt)
			return
		}

		randomNumber, err := rand.Int(rand.Reader, big.NewInt(100000000))
		if err != nil {
			return
		}

		cnt++

		if randomNumber.Int64() == 1234 {
			fmt.Printf("number reached, cnt: %d", cnt)
			return
		}

	}
}

func main() {
	monkey()
}
