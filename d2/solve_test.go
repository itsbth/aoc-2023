package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"Game 10: 2 blue, 11 red, 3 green; 4 blue, 11 red, 13 green; 4 blue, 15 green, 4 red; 1 blue, 3 green, 17 red",
	"Game 11: 9 green, 3 blue, 2 red; 11 blue, 16 green, 5 red; 7 blue, 3 red, 5 green; 7 red, 8 green, 10 blue",
	"Game 12: 13 green, 2 red, 2 blue; 1 red, 6 green; 5 green, 3 red, 8 blue",
	"Game 13: 2 blue, 5 green; 2 blue, 2 green; 2 blue, 2 red, 4 green",
}

var GOOD = []game{
	{
		id: 10,
		revealed: []map[string]int{
			{
				"blue":  2,
				"red":   11,
				"green": 3,
			},
			{
				"blue":  4,
				"red":   11,
				"green": 13,
			},
			{
				"blue":  4,
				"red":   4,
				"green": 15,
			},
			{
				"blue":  1,
				"red":   17,
				"green": 3,
			},
		},
	},
	{
		id: 11,
		revealed: []map[string]int{
			{
				"blue":  3,
				"red":   2,
				"green": 9,
			},
			{
				"blue":  11,
				"red":   5,
				"green": 16,
			},
			{
				"blue":  7,
				"red":   3,
				"green": 5,
			},
			{
				"blue":  10,
				"red":   7,
				"green": 8,
			},
		},
	},
	{
		id: 12,
		revealed: []map[string]int{
			{
				"blue":  2,
				"red":   2,
				"green": 13,
			},
			{
				"red":   1,
				"green": 6,
			},
			{
				"blue":  8,
				"red":   3,
				"green": 5,
			},
		},
	},
	// "Game 13: 2 blue, 5 green; 2 blue, 2 green; 2 blue, 2 red, 4 green",
	{
		id: 13,
		revealed: []map[string]int{
			{
				"blue":  2,
				"green": 5,
			},
			{
				"blue":  2,
				"green": 2,
			},
			{
				"blue":  2,
				"red":   2,
				"green": 4,
			},
		},
	},
}

func TestParse(t *testing.T) {
	for i := 0; i < len(INPUT); i++ {
		g := parseGame(INPUT[i])
		assert.Equal(t, GOOD[i], g)
	}
}

func TestParse2(t *testing.T) {
	for i := 0; i < len(INPUT); i++ {
		g, err := parseGame2(INPUT[i])
		assert.NoError(t, err)
		assert.Equal(t, GOOD[i], g)
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, line := range INPUT {
			parseGame(line)
		}
	}
}

func BenchmarkParse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, line := range INPUT {
			parseGame2(line)
		}
	}
}
