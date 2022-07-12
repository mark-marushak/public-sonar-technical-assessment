package query

import (
	"github.com/mark-marushak/public-sonar-technical-assessment/repository"
	"github.com/mark-marushak/public-sonar-technical-assessment/repository/regex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchOR(t *testing.T) {
	conds := []Node{
		{Phrase: "manchester united"},
		{Phrase: "manchester"},
		{Phrase: "united"},
		{Phrase: "man u"},
		{Phrase: "man"},
		{Phrase: "man utd"},
		{Phrase: "mufc"},
	}

	group := Node{CondType: OR}
	for i := 0; i < len(conds); i++ {
		group.Add(&conds[i])
	}

	assert.True(t, group.Search(repository.NewServiceCompare(regex.RegexCompare{Text: "man u united man utd"}).Compare), "Case didn't match")
}

func TestSearchAND(t *testing.T) {

	group1 := Node{CondType: OR}
	group1.Add(&Node{Phrase: "mbappe"})
	group1.Add(&Node{Phrase: "lukaku"})

	group2 := Node{CondType: AND}
	group2.Add(&group1)
	group2.Add(&Node{Phrase: "scored"})

	assert.True(t, group2.Search(repository.NewServiceCompare(regex.RegexCompare{Text: "lukaku scored"}).Compare), "Case didn't match")
}

func TestSearchStrongQuery(t *testing.T) {

	g11 := &Node{Phrase: "juventus"}

	g12 := Node{CondType: OR}
	g12.Add(&Node{Phrase: "real madrid"})
	g12.Add(&Node{Phrase: "realmadrid"})

	g13 := &Node{Phrase: "barcelona"}

	g1 := Node{CondType: OR}
	g1.Add(g11)
	g1.Add(&g12)
	g1.Add(g13)

	g21 := Node{CondType: OR}
	g21.Add(&Node{Phrase: "messi"})
	g21.Add(&Node{Phrase: "ronaldo"})
	g22 := Node{CondType: OR}
	g22.Add(&Node{Phrase: "goals"})
	g22.Add(&Node{Phrase: "goal"})

	g2 := Node{CondType: AND}
	g2.Add(&g21)
	g2.Add(&g22)

	group := Node{CondType: AND}
	group.Add(&g1)
	group.Add(&g2)

	assert.True(t, group.Search(repository.NewServiceCompare(regex.RegexCompare{Text: "messi goal juventus"}).Compare), "Case didn't match")
}

//func TestTableSearch(t *testing.T) {
//	tests := []struct {
//		name    string
//		expect  bool
//		message map[string]int
//		query   Node
//	}{
//		{
//			"Grouped query",
//			true,
//			map[string]int{
//				"juventus": 1,
//				"ronaldo":  1,
//				"goals":    1,
//			},
//			Node{
//				CondType: AND,
//				Conditions: []InterfaceNode{
//					&Node{
//						CondType: OR,
//						Conditions: []InterfaceNode{
//							&Node{Phrase: "juventus"},
//							&Node{CondType: OR, Conditions: []InterfaceNode{
//								&Node{Phrase: "real madrid"},
//								&Node{Phrase: "realmadrid"},
//							}},
//							&Node{Phrase: "barcelona"},
//						},
//					},
//					&Node{
//						CondType: AND,
//						Conditions: []InterfaceNode{
//							&Node{CondType: OR, Conditions: []InterfaceNode{
//								&Node{Phrase: "messi"},
//								&Node{Phrase: "ronaldo"},
//							}},
//							&Node{CondType: OR, Conditions: []InterfaceNode{
//								&Node{Phrase: "goal"},
//								&Node{Phrase: "goals"},
//							}},
//						},
//					},
//				},
//			},
//		},
//		{
//			"Long OR query",
//			true,
//			map[string]int{
//				"man utd": 1,
//			},
//			Node{
//				CondType: OR,
//				Conditions: []InterfaceNode{
//					&Node{Phrase: "manchester united"},
//					&Node{Phrase: "manchester"},
//					&Node{Phrase: "united"},
//					&Node{Phrase: "man u"},
//					&Node{Phrase: "man"},
//					&Node{Phrase: "man utd"},
//					&Node{Phrase: "mufc"},
//				},
//			},
//		},
//		{
//			"Shor OR with AND query",
//			true,
//			map[string]int{
//				"scored": 1,
//				"lukaku": 1,
//			},
//			Node{
//				CondType: AND,
//				Conditions: []InterfaceNode{
//					&Node{
//						CondType: OR,
//						Conditions: []InterfaceNode{
//							&Node{Phrase: "mbappe"},
//							&Node{Phrase: "lukaku"},
//						},
//					},
//					&Node{Phrase: "scored"},
//				},
//			},
//		},
//		{
//			"Query with one word",
//			true,
//			map[string]int{
//				"ajax": 1,
//			},
//			Node{
//				CondType: AND,
//				Conditions: []InterfaceNode{
//					&Node{Phrase: "ajax"},
//				},
//			},
//		},
//	}
//
//	var result bool
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			result = test.query.Search(test.message)
//			if result != test.expect {
//				t.Fail()
//			}
//		})
//	}
//}
