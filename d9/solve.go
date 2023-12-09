package d9

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

func predict(series []int) int {
	// recursively find the difference between all numbers in the series
	// until it's all zeroes, then build the next number in the series
	// by adding the difference to the last number in the series
	// return the next number in the series
	var deltas []int
	for i := 1; i < len(series); i++ {
		deltas = append(deltas, series[i]-series[i-1])
	}
	dv := deltas[0]
	allSame := true
	for _, d := range deltas {
		if d != dv {
			allSame = false
			break
		}
	}
	if allSame {
		return series[len(series)-1] + dv
	}
	return series[len(series)-1] + predict(deltas)
}

func (solver) Solve(input io.Reader) (int, int, error) {
	var series [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		var s []int
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				return 0, 0, err
			}
			s = append(s, n)
		}
		series = append(series, s)
	}

	part1 := 0
	for _, s := range series {
		part1 += predict(s)
	}

	part2 := 0
	for _, s := range series {
		slices.Reverse(s)
		part2 += predict(s)
	}

	return part1, part2, nil
}

func init() {
	runner.Register(2023, 9, solver{})
}
