package d10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample1(t *testing.T) {
	input := [][]byte{
		[]byte("....."),
		[]byte(".S-7."),
		[]byte(".|.|."),
		[]byte(".L-J."),
		[]byte("....."),
	}
	res := traverse(1, 1, DIR_DOWN, input)
	assert.Equal(t, 8, res)
}

func TestSample2(t *testing.T) {
	input := [][]byte{
		[]byte("..F7."),
		[]byte(".FJ|."),
		[]byte("SJ.L7"),
		[]byte("|F--J"),
		[]byte("LJ..."),
	}
	res := traverse(0, 2, DIR_RIGHT, input)
	assert.Equal(t, 16, res)
}
