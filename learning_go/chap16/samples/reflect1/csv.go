package reflectutils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Marshal(t any) ([][]string, error) {
	sliceVal := reflect.ValueOf(t)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}

	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)

	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}

		out = append(out, row)
	}

	return out, nil
}

func marshalHeader(t reflect.Type) []string {
	var row []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

func marshalOne(v reflect.Value) ([]string, error) {
	var row []string
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		fieldVal := v.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.Int:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
		}
	}

	return row, nil
}

func Unmarshal(data [][]string, t any) error {
	sliceValPointer := reflect.ValueOf(t)
	if sliceValPointer.Kind() != reflect.Pointer {
		return errors.New("must be a pointer to a slice of structs")
	}
	sliceVal := sliceValPointer.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("must be a pointer to a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("must be a pointer to a slice of struct")
	}

	header := data[0]
	numPos := make(map[string]int, len(header))
	for i, name := range header {
		numPos[name] = i
	}

	for _, row := range data[1:] {
		newVal := reflect.New(structType).Elem()
		err := unmarshalOne(row, numPos, newVal)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}

	return nil
}

func unmarshalOne(row []string, numPos map[string]int, v reflect.Value) error {
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		typeField := vt.Field(i)
		pos, ok := numPos[typeField.Tag.Get("csv")]
		if !ok {
			continue
		}
		val := row[pos]
		field := v.Field(i)

		switch field.Kind() {
		case reflect.Int:
			i, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(i)
		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(b)
		case reflect.String:
			field.SetString(val)
		default:
			return fmt.Errorf("cannot handle field of kind %v", field.Kind())
		}
	}
	return nil
}
