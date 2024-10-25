package jsonqs

import (
	"fmt"
	"reflect"
)

func encodeObject(input interface{}, subObj bool) (string, error) {
	var output string
	if subObj {
		output = "{"
	}

	var keySep string
	if subObj {
		keySep = ":"
	} else {
		keySep = "="
	}

	var fieldSep string
	if subObj {
		fieldSep = ","
	} else {
		fieldSep = "&"
	}

	t := reflect.TypeOf(input)
	val := reflect.ValueOf(input)
	fields := reflect.VisibleFields(t)
	for i, field := range fields {
		// TODO something with the tags
		if !field.IsExported() {
			continue
		}

		fieldVal := val.FieldByName(field.Name)
		encodedValue, err := encode(fieldVal.Interface())
		if err != nil {
			return output, err
		}

		output += fmt.Sprintf("%s%s%s", field.Name, keySep, encodedValue)
		if i > 0 {
			output += fieldSep
		}
	}

	if subObj {
		output += "}"
	}

	return output, nil
}
