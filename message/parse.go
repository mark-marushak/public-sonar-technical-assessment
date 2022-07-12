package message

import "github.com/mark-marushak/public-sonar-technical-assessment/parser"

func ParseJson(filepath string) ([]Message, error) {
	var output = make([]string, 0, 1000)
	err := parser.ParseJson(filepath, &output)
	if err != nil {
		return nil, err
	}

	messages := make([]Message, 0, len(output))
	for i := 0; i < len(output); i++ {
		messages = append(messages, Message{output[i]})
	}

	return messages, nil
}
