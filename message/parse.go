package message

import "public-sonar-technical-assessment/parser"

func ParseJson(filepath string) ([]Message, error) {
	var output = make([]Message, 0, 1000)
	err := parser.ParseJson(filepath, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
