package d9

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

func predict(series []int) (int, int) {
	// recursively find the difference between all numbers in the series
	// until it's all zeroes, then build the next number in the series
	// by adding the difference to the last number in the series
	// return the next number in the series
	deltas := make([]int, len(series)-1)
	allSame := true
	for i := 1; i < len(series); i++ {
		delta := series[i] - series[i-1]
		deltas[i-1] = delta
		allSame = allSame && delta == deltas[0]
	}
	if allSame {
		return series[0] - deltas[0], series[len(series)-1] + deltas[0]
	}
	pl, pu := predict(deltas)
	return series[0] - pl, series[len(series)-1] + pu
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
	part2 := 0
	for _, s := range series {
		p2, p1 := predict(s)
		part1 += p1
		part2 += p2
	}

	return part1, part2, nil
}

func init() {
	runner.Register(2023, 9, solver{})
}
