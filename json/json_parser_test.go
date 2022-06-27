package json

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	jsonValue, err := ParseJson("{\"name\":\"\\\"John Gruber\\\"\",\"age\":30,\"city\":\"New York\"}")
	if err != nil {
		t.Error(err)
	}

	jsonObject := jsonValue.(JsonObject)
	fmt.Println(jsonObject)
}
