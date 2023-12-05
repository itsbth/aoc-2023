package d5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48
*/

func TestTranslateRange(t *testing.T) {
	mapping := mapping{
		entries: []entry{
			{50, 52, 98 - 50},
			{52, 100, 50 - 52},
		},
	}
	ranges := []span{
		{79, 79 + 14 + 1},
		{55, 55 + 13 + 1},
	}
	for r := range ranges {
		for v := ranges[r].start; v < ranges[r].end; v++ {
			// compare a known good m.translate(v) with m.translateRange(v..v+1)
			good := mapping.translate(v)
			got := mapping.translateRange(span{v, v + 1})
			assert.Len(t, got, 1, "translateRange(v, v+1) should have 1 returned value, got %d", len(got))
			assert.Equal(t, 1, got[0].end-got[0].start, "translateRange(v, v+1) should have a return of size 1")
			assert.Equal(t, good, got[0].start, "translateRange(v, v+1) should return the same value as translate(v)")
		}
	}
}
