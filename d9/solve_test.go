package d9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredict(t *testing.T) {
	// 0 3 6 9 12 15
	inp := []int{0, 3, 6, 9, 12, 15}

	assert.Equal(t, 18, predict(inp))
}

func TestPredict2(t *testing.T) {
	// 1 3 6 10 15 2
	inp := []int{1, 3, 6, 10, 15, 2}

	assert.Equal(t, 18, predict(inp))
}

func TestPredict3(t *testing.T) {
	// 10 13 16 21 30 45
	inp := []int{0, 3, 6, 9, 12, 15}

	assert.Equal(t, 18, predict(inp))
}
