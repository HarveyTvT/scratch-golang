package internal

import (
	"strconv"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRUCache[string, string](10)

	for i := range 100 {
		v := ""
		lru.Put(strconv.FormatInt(int64(i), 10), &v)
	}
}
