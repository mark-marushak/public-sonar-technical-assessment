package message

type InterfaceMessage interface {
	SetText(string)
	GetText() string
}

type Message struct {
	Text string
}

func (m *Message) SetText(message string) {
	m.Text = message
}

func (m Message) GetText() string {
	return m.Text
}

//
//func (m *Message) CollectWords() {
//	m.Words = make(map[string]map[string]bool)
//	words := strings.Split(m.Text, " ")
//	var word string
//	for i := 0; i < len(words); i++ {
//		word = words[i]
//
//		if _, ok := m.Words[word]; !ok {
//			m.Words[word] = make(map[string]bool)
//		}
//
//		if i-1 >= 0 {
//			m.Words[word][words[i-1]] = true
//		}
//
//		if i+1 <= len(words)-1 {
//			m.Words[word][words[i+1]] = true
//		}
//	}
//}
