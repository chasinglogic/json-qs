package jsonqs_test

import (
	"testing"

	"github.com/chasinglogic/json-qs/pkg/encoding/jsonqs"
)

type kitchenSink struct {
	Array       []int
	String      string
	Fraction    float64
	True        bool
	False       bool
	Null        *int
	NestedArray [][]int
	EmptyArray  []int
	EmptyObject struct{}
	Object      struct {
		A int
		B int
	}
	ObjectInArray []struct {
		A int
		B int
	}
}

func TestEncodeObject(t *testing.T) {
	input := kitchenSink{
		Array:       []int{1, 2, 3},
		String:      "Tom",
		Fraction:    float64(1.23),
		True:        true,
		False:       false,
		Null:        nil,
		NestedArray: [][]int{{1, 2, 3}, {4, 5, 6}},
		EmptyArray:  []int{},
		EmptyObject: struct{}{},
		Object: struct {
			A int
			B int
		}{A: 1, B: 2},
		ObjectInArray: []struct {
			A int
			B int
		}{{A: 1, B: 2}},
	}

	encoded, err := jsonqs.Encode(input)
	t.Error(err)
	t.Error(encoded)
}
