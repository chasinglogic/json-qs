package jsonqs

import "fmt"

func encodeLiteral(input interface{}) (string, error) {
	switch val := input.(type) {
	case string:
		return encodeString(val), nil
	case int16, int32, int64, int:
		return fmt.Sprintf("%d", val), nil
	case float32, float64:
		return fmt.Sprintf("%f", val), nil
	case bool:
		if val {
			return "true", nil
		}
		return "false", nil
	case nil:
		return "null", nil
	default:
		return "", fmt.Errorf("non-encodable value: %v %T", val, input)
	}
}

func decodeLiteral(input string) interface{} {
	return nil
}
