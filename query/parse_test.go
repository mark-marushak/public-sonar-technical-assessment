package query

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
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

//(juventus OR (real madrid OR realmadrid) OR barcelona) AND ((messi OR ronaldo) AND (goal OR goals))
func TestParseQueryBuilder(t *testing.T) {
	//s := "(juventus OR (real madrid OR realmadrid) OR barcelona) AND ((messi OR ronaldo) AND (goal OR goals))"
	s := "manchester united OR manchester OR united OR man u OR man OR man utd OR mufc"
	ParseQuery(s)
}

func TestOpenGroupFunc(t *testing.T) {
	var queue = make([]InterfaceNode, 0, 10)

	queue = OpenGroupFunc(queue)
	queue = OpenGroupFunc(queue)
	queue = OpenGroupFunc(queue)
	queue = OpenGroupFunc(queue)

	assert.Equal(t, 4, len(queue))
}

func TestCloseGroupFunc(t *testing.T) {
	var queue = make([]InterfaceNode, 0, 10)

	queue = OpenGroupFunc(queue)
	queue = OpenGroupFunc(queue)
	queue = OpenGroupFunc(queue)
	queue = OpenGroupFunc(queue)

	queue = CloseGroupFunc(queue)
	queue = CloseGroupFunc(queue)
	queue = CloseGroupFunc(queue)

	assert.Equal(t, 1, len(queue))
}

func TestGroupCondFunc(t *testing.T) {
	var queue = make([]InterfaceNode, 0, 10)

	queue = OpenGroupFunc(queue)
	OrGroupCondFunc(queue)
	queue = OpenGroupFunc(queue)
	OrGroupCondFunc(queue)
	queue = OpenGroupFunc(queue)
	AndGroupCondFunc(queue)

	queue = CloseGroupFunc(queue)
	queue = CloseGroupFunc(queue)

	expect := []InterfaceNode{
		&Node{
			CondType: OR,
			Conditions: []InterfaceNode{
				&Node{
					CondType: OR,
					Conditions: []InterfaceNode{
						&Node{
							CondType: AND,
						},
					},
				},
			},
		},
	}

	if reflect.DeepEqual(expect, queue) == false {
		t.Fail()
	}
}

func TestSetPhraseFunc(t *testing.T) {
	var queue = make([]InterfaceNode, 0, 10)

	queue = OpenGroupFunc(queue)
	SetPhraseFunc(queue, "Hello")
	OrGroupCondFunc(queue)
	queue = OpenGroupFunc(queue)
	SetPhraseFunc(queue, "mark")
	OrGroupCondFunc(queue)
	queue = OpenGroupFunc(queue)
	SetPhraseFunc(queue, "one")
	AndGroupCondFunc(queue)

	queue = CloseGroupFunc(queue)
	queue = CloseGroupFunc(queue)

	expect := []InterfaceNode{
		&Node{
			CondType: OR,
			Conditions: []InterfaceNode{
				&Node{Phrase: "Hello"},
				&Node{
					CondType: OR,
					Conditions: []InterfaceNode{
						&Node{Phrase: "mark"},
						&Node{
							CondType: AND,
							Conditions: []InterfaceNode{
								&Node{Phrase: "one"},
							},
						},
					},
				},
			},
		},
	}

	if reflect.DeepEqual(expect, queue) == false {
		t.Fail()
	}

	result := queue[0].Search(map[string]int{
		"Hello": 1,
		"mark":  1,
		"one":   1,
	})

	assert.True(t, result)
}
