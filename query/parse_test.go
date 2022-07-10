package query

import (
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	queries, err := ParseJson("/home/sandbox/GolandProjects/public-sonar-technical-assessment/storage/queries/cases.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(queries) <= 0 {
		t.FailNow()
	}

	for i := 0; i < len(queries); i++ {
		if queries[i].CaseID == 0 {
			t.FailNow()
		}

		if queries[i].Query == "" {
			t.FailNow()
		}
	}
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
