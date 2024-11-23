package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ValidateStringLength(v any) error {
	vv := reflect.ValueOf(v)
	vt := reflect.TypeOf(v)
	if vt.Kind() != reflect.Struct {
		return errors.New("must be a struct")

	}

	var err error
	for i := 0; i < vt.NumField(); i++ {
		fieldVal := vv.Field(i)
		fieldType := vt.Field(i)
		if fieldVal.Kind() != reflect.String {
			continue
		}

		minStrLenTag, ok := fieldType.Tag.Lookup("minStrLen")
		if !ok {
			continue
		}

		minStrLen, pErr := strconv.ParseInt(minStrLenTag, 10, 64)
		if pErr != nil {
			err = errors.Join(err, pErr)
			continue
		}
		if fieldVal.Len() > int(minStrLen) {
			err = errors.Join(err, fmt.Errorf("string field %s exceed max len: %d/%d", fieldType.Name, fieldVal.Len(), minStrLen))
			continue
		}
	}

	return err
}

func main() {
	foo := struct {
		a string `minStrLen:"1"`
		b string `minStrLen:"2"`
		c string
		d bool
	}{
		a: "123",
		b: "1234",
		c: "123321321",
		d: false,
	}

	if err := ValidateStringLength(foo); err != nil {
		fmt.Println(err)
	}

}
