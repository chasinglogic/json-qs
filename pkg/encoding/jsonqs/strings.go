package jsonqs

import (
	"io"
	"net/url"
	"unicode"
)

var alwaysEscaped = "{}(),:"

func encodeString(input string) string {
	switch input {
	case "null", "true", "false":
		return "\\" + input
	}

	output := make([]byte, 0, len(input))
	for i, r := range input {
		if i == 0 && (unicode.IsDigit(r) || r == '-' || r == '\\') {
			output = append(output, '\\', byte(r))
			continue
		}

		switch r {
		case '{', '}', '(', ')', ',', ':':
			output = append(output, '\\', byte(r))
		case ' ':
			output = append(output, '+')
		case '#':
			output = append(output, '%', '2', '3')
		case '&':
			output = append(output, '%', '2', '6')
		case '"':
			output = append(output, '%', '2', '2')
		case '%':
			output = append(output, '%', '2', '5')
		case '+':
			output = append(output, '%', '2', 'B')
		default:
			output = append(output, byte(r))
		}
	}

	return string(output)
}

func decodeString(input string) (string, error) {
	pos := 0
	char := byte(0)
	peekChar := byte(0)

	advance := func() error {
		char = peekChar
		if pos >= len(input) {
			peekChar = byte(0)
		} else {
			peekChar = input[pos]
			pos++
		}

		if char == byte(0) {
			return io.EOF
		}

		return nil
	}

	advance()
	output := make([]byte, 0, len(input))
	for {
		if err := advance(); err != nil {
			break
		}

		switch char {
		case '\\':
			if peekChar == '\\' {
				output = append(output, '\\')
				advance()
			}

			continue
		case '+':
			output = append(output, ' ')
		default:
			output = append(output, char)
		}
	}

	return url.QueryUnescape(string(output))
}
