package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	err := ValidatePerson(Person{"", "", -1})
	fmt.Println(err)
}
