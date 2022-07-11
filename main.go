package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-errors/errors"
	"io"
	"os"
	"public-sonar-technical-assessment/message"
	"public-sonar-technical-assessment/query"
	"public-sonar-technical-assessment/repository"
	"public-sonar-technical-assessment/repository/regex"
)

func main() {
	pathToQueries := flag.String("cases", "/home/sandbox/GolandProjects/public-sonar-technical-assessment/storage/queries/cases.json", "Path to cases.")
	pathToMessages := flag.String("messages", "/home/sandbox/GolandProjects/public-sonar-technical-assessment/storage/messages/messages.json", "Path to messages.")
	pathToOutput := flag.String("outputs", "/home/sandbox/GolandProjects/public-sonar-technical-assessment/storage/output/outputs.json", "Path to output.")

	queries, err := query.ParseJson(*pathToQueries)
	if err != nil {
		fmt.Println(errors.Errorf("error while parsing queries: %v", err).ErrorStack())
		return
	}

	messages, err := message.ParseJson(*pathToMessages)
	if err != nil {
		fmt.Println(errors.Errorf("error while parsing messages: %v", err).ErrorStack())
		return
	}

	var builded query.InterfaceNode
	var repo repository.InterfaceCompareFunc
	var service repository.ServiceCompare
	type OutputFormat struct {
		QueryID  int
		Messages []string
	}
	found := make([]OutputFormat, 0, 10)

	for i := 0; i < len(queries); i++ {
		builded = query.ParseString(queries[i].Query)
		outputFormat := OutputFormat{
			queries[i].CaseID,
			make([]string, 0, 10),
		}
		for j := 0; j < len(messages); j++ {
			repo = regex.RegexCompare{messages[j].GetText()}
			service = repository.NewServiceCompare(repo)
			if builded.Search(service.Compare) {
				outputFormat.Messages = append(outputFormat.Messages, messages[j].GetText())
			}
		}
		found = append(found, outputFormat)
	}

	outJson, err := json.Marshal(found)
	if err != nil {
		return
	}
	outputs, err := os.Create(*pathToOutput)
	if err != nil {
		return
	}

	_, err = io.WriteString(outputs, string(outJson))
	if err != nil {
		return
	}

}
