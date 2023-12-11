package d11

import (
	"bufio"
	"io"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

type star struct {
	x, y int
}

func (solver) Solve(input io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(input)
	var stars []star
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			if c == '#' {
				stars = append(stars, star{x, y})
			}
		}
		y++
	}

	starX := make([]bool, y)
	starY := make([]bool, y)

	for _, s := range stars {
		starX[s.x] = true
		starY[s.y] = true
	}

	translateX := make([]int, y)
	translateY := make([]int, y)
	translateX[0] = 0
	translateY[0] = 0

	for i := 1; i < y; i++ {
		if starX[i] {
			translateX[i] = translateX[i-1] + 1
		} else {
			translateX[i] = translateX[i-1] + 2
		}
		if starY[i] {
			translateY[i] = translateY[i-1] + 1
		} else {
			translateY[i] = translateY[i-1] + 2
		}
	}

	part1 := 0

	for idx, s1 := range stars {
		for _, s2 := range stars[idx+1:] {
			dist := abs(
				translateX[s1.x]-translateX[s2.x],
			) + abs(
				translateY[s1.y]-translateY[s2.y],
			)
			part1 += dist
		}
	}

	translateX[0] = 0
	translateY[0] = 0

	for i := 1; i < y; i++ {
		if starX[i] {
			translateX[i] = translateX[i-1] + 1
		} else {
			translateX[i] = translateX[i-1] + 1_000_000
		}
		if starY[i] {
			translateY[i] = translateY[i-1] + 1
		} else {
			translateY[i] = translateY[i-1] + 1_000_000
		}
	}

	part2 := 0

	for idx, s1 := range stars {
		for _, s2 := range stars[idx+1:] {
			dist := abs(
				translateX[s1.x]-translateX[s2.x],
			) + abs(
				translateY[s1.y]-translateY[s2.y],
			)
			part2 += dist
		}
	}

	return part1, part2, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func init() {
	runner.Register(2023, 11, solver{})
}
