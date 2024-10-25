package jsonqs

import (
	"fmt"
	"reflect"
)

func encode(input interface{}) (string, error) {
	t := reflect.TypeOf(input)
	switch t.Kind() {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.String:
		return encodeLiteral(input)

	case reflect.Array, reflect.Slice:
		return encodeArray(input.([]interface{}))

	case reflect.Struct:
		return encodeObject(input, true)
		// reflect.Interface,

		// reflect.Map,

	case reflect.Pointer:
		// Here be dragons!
		return encode(*input.(*interface{}))
	default:
		return "", fmt.Errorf("non-encodeable type: %s", t.Name())
	}
}

func Encode(input interface{}) (string, error) {
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("%s is not an object", t.Name())
	}

	return encodeObject(input, false)
}
