package message

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCollectWords(t *testing.T) {
	message := Message{
		Text: "@ManCityStand Who knows??",
	}

	message.CollectWords()

	result := reflect.DeepEqual(message.Words, map[string]map[string]bool{
		"@ManCityStand": {
			"Who": true,
		},
		"Who": {
			"@ManCityStand": true,
			"knows??":       true,
		},
		"knows??": {
			"Who": true,
		},
	})

	assert.True(t, result, "Collected words hasn't right consistency")
}
