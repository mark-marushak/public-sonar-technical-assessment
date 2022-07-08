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
