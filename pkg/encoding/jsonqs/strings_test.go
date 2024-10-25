package jsonqs

import (
	"strings"
	"testing"
)

var stringEncodingTestCases = map[string]string{
	"123":                                    "\\123",
	"foo":                                    "foo",
	"{b:0}":                                  "\\{b\\:0\\}",
	"(b,0)":                                  "\\(b\\,0\\)",
	"1{b:0}(b,0)":                            "\\1\\{b\\:0\\}\\(b\\,0\\)",
	"-1":                                     "\\-1",
	"\\":                                     "\\\\",
	"\\stuff":                                "\\\\stuff",
	"a123":                                   "a123",
	"null":                                   "\\null",
	"true":                                   "\\true",
	"false":                                  "\\false",
	" ":                                      "+",
	"the man in the hat has a very nice cat": "the+man+in+the+hat+has+a+very+nice+cat",
	"\"a quoted string\"":                    "%22a+quoted+string%22",
	"sausage&pickle":                         "sausage%26pickle",
	"#number":                                "%23number",
	"pickle+mustard":                         "pickle%2Bmustard",
	"100%miserable":                          "\\100%25miserable",
}

func TestStringEncoding(t *testing.T) {
	for input, expected := range stringEncodingTestCases {
		t.Run(input, func(subt *testing.T) {
			encoded := encodeString(input)
			if encoded != expected {
				subt.Errorf("\n\texpected: %s\n\t     got: %s", expected, encoded)
			}
		})
	}
}

func TestStringDecoding(t *testing.T) {
	for expected, input := range stringEncodingTestCases {
		t.Run(input, func(subt *testing.T) {
			decoded, _ := decodeString(input)
			if decoded != expected {
				subt.Errorf("\n\texpected: %s\n\t     got: %s", expected, decoded)
			}
		})
	}
}

func BenchmarkEncodeString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		encodeString("1{b:0}#(b,0)")
	}
}

var bigString = strings.Repeat("1{b:0}#(b,0)", 500)

func BenchmarkEncodeBigString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		encodeString(bigString)
	}
}

func BenchmarkDecodeString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		decodeString("\\1\\{b\\:0\\}%23\\(b\\,0\\)")
	}
}

var bigEncodedString = strings.Repeat(encodeString("1{b:0}#(b,0)"), 500)

func BenchmarkDecodeBigString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		decodeString(bigEncodedString)
	}
}
