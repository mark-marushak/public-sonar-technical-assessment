package query

import (
	"strings"
	"unicode"
)

type Parser interface {
	Decode(interface{}) error
}

func Parse(parser Parser, v interface{}) error {
	return parser.Decode(v)
}

func ParseQuery(queries string) InterfaceNode {
	const (
		OpenGroup  = '('
		CloseGroup = ')'
		OR         = '|'
		AND        = '&'
	)

	queries = strings.ReplaceAll(queries, "((", "( (")
	queries = strings.ReplaceAll(queries, "))", ") )")
	stringWithOutOR := strings.ReplaceAll(queries, "OR", "|")
	stringWithOutAND := strings.ReplaceAll(stringWithOutOR, "AND", "&")

	var formated = stringWithOutAND

	var queue []InterfaceNode
	for start, r := range formated {
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
				for end, r2 := range formated[start:] {
					if unicode.IsLetter(r2) == false && unicode.IsSpace(r2) == false {
						SetPhraseFunc(queue, formated[start:end])
					}
				}
			}
		}
	}

	return queue[0]
}

func OpenGroupFunc(queue []InterfaceNode) []InterfaceNode {
	queue = append(queue, &Node{})
	return queue
}

func CloseGroupFunc(queue []InterfaceNode) []InterfaceNode {
	if len(queue)-1 == 0 {
		return queue
	}

	queue[len(queue)-2].Add(queue[len(queue)-1])
	queue = queue[:len(queue)-1]
	return queue
}

func OrGroupCondFunc(queue []InterfaceNode) {
	queue[len(queue)-1].SetCondType(OR)
}

func AndGroupCondFunc(queue []InterfaceNode) {
	queue[len(queue)-1].SetCondType(AND)
}

func SetPhraseFunc(queue []InterfaceNode, phrase string) {
	queue[len(queue)-1].Add(&Node{Phrase: phrase})
}
