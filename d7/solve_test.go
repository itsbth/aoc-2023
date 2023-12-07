package d7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandType(t *testing.T) {
	tests := []struct {
		hand string
		best handType
	}{
		{"AAAAA", HAND_FIVE_OF_A_KIND},
		{"AAAA2", HAND_FOUR_OF_A_KIND},
		{"AAA22", HAND_FULL_HOUSE},
		{"AK222", HAND_THREE_OF_A_KIND},
		{"AA234", HAND_ONE_PAIR},
		{"A2345", HAND_HIGH_CARD},
	}
	for _, test := range tests {
		h := NewHand(test.hand, 0)
		assert.Equal(t, test.best, h.best, "expected %s to be %s", test.hand, test.best)
	}
}
