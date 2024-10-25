package jsonqs

func encodeArray(input []interface{}) (string, error) {
	output := "("
	for i, val := range input {
		if i > 0 {
			output += ","
		}

		encoded, err := encode(val)
		if err != nil {
			return output, err
		}

		output += encoded
	}

	return output + ")", nil
}

func decodeArray(input string) ([]interface{}, error) {
	return nil, nil
}
