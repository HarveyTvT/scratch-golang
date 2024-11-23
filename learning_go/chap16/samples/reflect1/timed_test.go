package reflectutils

import (
	"testing"
	"time"
)

func timeMe(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}

func TestTimedMe(t *testing.T) {
	timed := MakeTimedFunc(timeMe)
	result := timed.(func(int) int)(2)
	if result != 4 {
		t.Error("Expected 4, got ", result)
	}
}
