package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchOR(t *testing.T) {
	conds := []Condition{
		{"manchester united"},
		{"manchester"},
		{"united"},
		{"man u"},
		{"man"},
		{"man utd"},
		{"mufc"},
	}

	group := Group{CondType: OR}
	for i := 0; i < len(conds); i++ {
		group.Add(conds[i])
	}

	message := map[string]int{
		"man u":   1,
		"united":  1,
		"man utd": 1,
	}

	assert.True(t, group.search(message), "Case didn't match")
}

func TestSearchAND(t *testing.T) {

	group1 := Group{CondType: OR}
	group1.Add(Condition{"mbappe"})
	group1.Add(Condition{"lukaku"})

	group2 := Group{CondType: AND}
	group2.Add(group1)
	group2.Add(Condition{"scored"})

	message := map[string]int{
		"lukaku": 1,
		"scored": 1,
	}

	assert.True(t, group2.search(message), "Case didn't match")
}

func TestSearchStrongQuery(t *testing.T) {

	g11 := Condition{"juventus"}

	g12 := Group{CondType: OR}
	g12.Add(Condition{"real madrid"})
	g12.Add(Condition{"realmadrid"})

	g13 := Condition{"barcelona"}

	g1 := Group{CondType: OR}
	g1.Add(g11)
	g1.Add(g12)
	g1.Add(g13)

	g21 := Group{CondType: OR}
	g21.Add(Condition{"messi"})
	g21.Add(Condition{"ronaldo"})
	g22 := Group{CondType: OR}
	g22.Add(Condition{"goals"})
	g22.Add(Condition{"goal"})

	g2 := Group{CondType: AND}
	g2.Add(g21)
	g2.Add(g22)

	group := Group{CondType: AND}
	group.Add(g1)
	group.Add(g2)

	message := map[string]int{
		"messi":    1,
		"goal":     1,
		"juventus": 1,
	}

	assert.True(t, group.search(message), "Case didn't match")
}
