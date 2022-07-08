package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("/home/sandbox/GolandProjects/public-sonar-technical-assessment/storage/queries/cases.json")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer file.Close()

	type (
		Query struct {
			CaseID  int
			Query   string
			Queries string
		}
		Queries []Query
	)

	queries := Queries{}

	err = Parse(json.NewDecoder(file), &queries)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(queries)
}
