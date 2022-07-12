package message

import (
	"log"
	"testing"
)

func TestParseJson(t *testing.T) {
	messages, err := ParseJson("/home/sandbox/GolandProjects/github.com/mark-marushak/public-sonar-technical-assessment/storage/messages/messages.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(messages) <= 0 {
		t.FailNow()
	}

	for i := 0; i < len(messages); i++ {
		if messages[i].GetText() == "" {
			log.Fatal("Empty message")
		}
	}
}
