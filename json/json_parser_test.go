package json

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	tokens := Tokenize("{\"name\":\"\\\"John Gruber\\\"\",\"age\":30,\"city\":\"New York\"}")
	for _, token := range tokens {
		fmt.Println(token)
	}
}
