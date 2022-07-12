package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/mark-marushak/public-sonar-technical-assessment/message"
	"github.com/mark-marushak/public-sonar-technical-assessment/query"
	"github.com/mark-marushak/public-sonar-technical-assessment/repository"
	"github.com/mark-marushak/public-sonar-technical-assessment/repository/regex"
	"io"
	"os"
)

func main() {
	pathToQueries := flag.String("cases", "./storage/queries/cases.json", "Path to cases.")
	pathToMessages := flag.String("messages", "./storage/messages/messages.json", "Path to messages.")
	pathToOutput := flag.String("outputs", "./storage/output/outputs.json", "Path to output.")

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
