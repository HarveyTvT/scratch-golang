package reflectutils

import (
	"fmt"
	"reflect"
	"time"
)

func MakeTimedFunc(f any) any {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)

	wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fv.Call(in)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})

	return wrapperF.Interface()
}
