package query

/**
(\((.*?)\)(*ACCEPT))
(mbappe OR lukaku) AND scored
*/

const (
	AND = "AND"
	OR  = "OR"
)

type InterfaceNode interface {
	search(map[string]int) bool
}

type Group struct {
	Conditions []InterfaceNode
	CondType   string
	Limit      int
}

func (g *Group) Add(node InterfaceNode) {
	g.Conditions = append(g.Conditions, node)
}

func (g Group) search(message map[string]int) bool {
	var flag = false
	if g.CondType == OR {
		for i := 0; i < len(g.Conditions); i++ {
			if g.Conditions[i].search(message) == true {
				flag = true
			}
		}

		return flag
	}

	if g.CondType == AND {
		flag = true
		for i := 0; i < len(g.Conditions); i++ {
			if g.Conditions[i].search(message) == false {
				flag = false
				break
			}
		}

	}

	return flag
}

type Condition struct {
	Phrase string
}

func (cond Condition) search(message map[string]int) bool {
	_, ok := message[cond.Phrase]
	return ok
}
