package d1

import (
	"bufio"
	"io"

	"github.com/itsbth/aoc-2023/runner"
)

var MAP = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var m = matcher{}

func sum(line string) (uint64, uint64) {
	var sum1, sum2 uint64

	found := false

	m.reset()

	// forward scan, find first ASCII digit and first spelled-out digit
outer:
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			val := uint64(line[i]-'0') * 10
			sum1 += val
			if !found {
				sum2 += val
			}
			break outer
		} else if !found && line[i] >= 'a' && line[i] <= 'z' {
			if v, ok := m.match(line[i]); ok {
				val := uint64(v) * 10
				sum2 += val
				found = true
			}
		}
	}

	// reverse scan, find last ASCII digit and last spelled-out digit
	found = false
	m.resetBW()
outer2:
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			val := uint64(line[i] - '0')
			sum1 += val
			if !found {
				sum2 += val
			}
			break outer2
		} else if !found && line[i] >= 'a' && line[i] <= 'z' {
			if v, ok := m.matchBW(line[i]); ok {
				val := uint64(v)
				sum2 += val
				found = true
			}
		}
	}

	return sum1, sum2
}

type solver struct{}

var _ runner.Solver = solver{}

func (solver) Solve(input io.Reader) (int, int, error) {
	total1, total2 := uint64(0), uint64(0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		c1, c2 := sum(line)
		total1 += c1
		total2 += c2
	}

	return int(total1), int(total2), nil
}

func init() {
	runner.Register(2023, 1, solver{}, "default")
}
