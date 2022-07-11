package query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"public-sonar-technical-assessment/repository"
	"public-sonar-technical-assessment/repository/regex"
	"reflect"
	"strings"
	"testing"
	"unicode"
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

	result := queue[0].Search(repository.NewServiceCompare(regex.RegexCompare{"Hello mark one"}).Compare)

	assert.True(t, result)
}

func TestDefineType(t *testing.T) {
	queries := "(juventus OR (real madrid OR realmadrid) OR barcelona) AND ((messi OR ronaldo) AND (goal OR goals))"

	const (
		OpenGroup  = rune('(')
		CloseGroup = rune(')')
		OR         = rune('|')
		AND        = rune('&')
	)

	queries = strings.ReplaceAll(queries, "((", "( (")
	queries = strings.ReplaceAll(queries, "))", ") )")
	stringWithOutOR := strings.ReplaceAll(queries, "OR", "|")
	stringWithOutAND := strings.ReplaceAll(stringWithOutOR, "AND", "&")

	var formated = stringWithOutAND

	var queue = make([]InterfaceNode, 0, 10)
	queue = append(queue, &Node{})
	for i := 0; i < len(formated); i++ {
		if len(strings.Split(formated, " ")) <= 1 {
			AndGroupCondFunc(queue)
			SetPhraseFunc(queue, formated)
			break
		}

		r := rune(formated[i])

		switch r {
		case OpenGroup:
			OpenGroupFunc(queue)
		case CloseGroup:
			queue = CloseGroupFunc(queue)
		case OR:
			OrGroupCondFunc(queue)
		case AND:
			AndGroupCondFunc(queue)
		default:
			if unicode.IsLetter(r) {
				for end, r2 := range formated[i:] {
					if unicode.IsLetter(r2) == false && unicode.IsSpace(r2) == false {
						SetPhraseFunc(queue, formated[i:i+end])
						formated = formated[end:]
						break
					}
				}
			}
		}
	}

	fmt.Printf("%#v", queue[0])
}
