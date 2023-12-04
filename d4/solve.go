package d4

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

func (solver) Solve(input io.Reader) (int, int, error) {
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	scanner := bufio.NewScanner(input)
	part1 := 0
	part2 := 0
	mul := []int{1}
	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		match := make([]int, 100)
		parts := strings.Split(line, " ")
		for _, part := range parts[2:] {
			if part == "|" {
				continue
			}
			val, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			match[val]++
		}
		both := 0
		for _, v := range match {
			if v == 2 {
				both++
			}
		}
		// make sure mul is at least as long as idx+both
		for len(mul) < idx+both+1 {
			mul = append(mul, 1)
		}
		// first is 1, then doubles every new
		part1 += int(math.Pow(2, float64(both)-1))
		part2 += mul[idx]
		for i := 0; i < both; i++ {
			mul[idx+1+i] += mul[idx]
		}
		idx++
	}
	return part1, part2, scanner.Err()
}

func init() {
	runner.Register(2023, 4, solver{})
}
