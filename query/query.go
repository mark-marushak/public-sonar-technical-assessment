package query

var (
	AND CondType = condType("AND")
	OR  CondType = condType("OR")
)

type CondType *string

func condType(s string) *string {
	return &s
}

type InterfaceNode interface {
	Search(map[string]int) bool
	Add(node InterfaceNode)
	SetCondType(condType CondType)
	SetPhrase(string)
}

type Node struct {
	Phrase     string
	Conditions []InterfaceNode
	CondType   CondType
}

func (g *Node) Add(node InterfaceNode) {
	g.Conditions = append(g.Conditions, node)
}

func (g *Node) SetCondType(condType CondType) {
	g.CondType = condType
}

func (g *Node) SetPhrase(s string) {
	g.Phrase = s
}

func (g Node) Search(message map[string]int) bool {
	var flag = false
	if g.CondType == OR {
		for i := 0; i < len(g.Conditions); i++ {
			if g.Conditions[i].Search(message) == true {
				flag = true
			}
		}

		return flag
	}

	if g.CondType == AND {
		flag = true
		for i := 0; i < len(g.Conditions); i++ {
			if g.Conditions[i].Search(message) == false {
				flag = false
				break
			}
		}

		return flag
	}

	_, ok := message[g.Phrase]
	return ok

}
