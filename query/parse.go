package query

import (
	"public-sonar-technical-assessment/parser"
	"strings"
	"unicode"
)

func ParseJson(filepath string) ([]Query, error) {
	var output = make([]QueryDTO, 0, 100)
	err := parser.ParseJson(filepath, &output)
	if err != nil {
		return nil, err
	}

	queries := make([]Query, 0, len(output))
	for i := 0; i < len(output); i++ {
		dtoQuery := output[i]
		if dtoQuery.Queries != "" {
			queries = append(queries, Query{
				CaseID: dtoQuery.CaseID,
				Query:  dtoQuery.Queries,
			})
		}

		if dtoQuery.Query != "" {
			queries = append(queries, Query{
				CaseID: dtoQuery.CaseID,
				Query:  dtoQuery.Query,
			})
		}
	}

	return queries, nil
}

func ParseString(queries string) InterfaceNode {
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
