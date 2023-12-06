package d6

import (
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

type race struct {
	duration int
	record   int
}

func (solver) Solve(input io.Reader) (int, int, error) {
	numRe := regexp.MustCompile(`(\d+)`)
	buf, err := io.ReadAll(input)
	if err != nil {
		return 0, 0, err
	}
	lines := strings.Split(string(buf), "\n")
	var races []race
	durs := numRe.FindAllString(lines[0], -1)
	records := numRe.FindAllString(lines[1], -1)
	for i := range durs {
		dn, err := strconv.Atoi(durs[i])
		if err != nil {
			return 0, 0, err
		}
		rn, err := strconv.Atoi(records[i])
		if err != nil {
			return 0, 0, err
		}
		races = append(races, race{
			duration: dn,
			record:   rn,
		})
	}

	part1 := 1

	for _, race := range races {
		wins := 0
		for i := 0; i < race.duration; i++ {
			inv := race.duration - i
			dist := i * inv
			if dist > race.record {
				wins += 1
			}
		}
		part1 *= wins
	}

	p2Dur := strings.Join(durs, "")
	p2Rec := strings.Join(records, "")
	part2 := 0
	dur, err := strconv.Atoi(p2Dur)
	if err != nil {
		return 0, 0, err
	}
	rec, err := strconv.Atoi(p2Rec)
	if err != nil {
		return 0, 0, err
	}
	for i := 0; i < dur; i++ {
		inv := dur - i
		dist := i * inv
		if dist > rec {
			part2 += 1
		}
	}

	return part1, part2, nil
}

type solverAnalytic struct{}

var _ runner.Solver = solverAnalytic{}

func solveEquation(duration, record int) int {
	// quadratic equation
	// n**2 + n * duration - record = 0
	// a: 1, b: duration, c: -record
	a := 1.0
	b := -float64(duration)
	c := float64(record)

	cp := math.Sqrt(b*b - 4*a*c)

	c1 := (-b + cp) / (2 * a)
	c2 := (-b - cp) / (2 * a)
	return int(math.Ceil(c1)-math.Floor(c2)) - 1
}

func (solverAnalytic) Solve(input io.Reader) (int, int, error) {
	numRe := regexp.MustCompile(`(\d+)`)
	buf, err := io.ReadAll(input)
	if err != nil {
		return 0, 0, err
	}
	lines := strings.Split(string(buf), "\n")
	var races []race
	durs := numRe.FindAllString(lines[0], -1)
	records := numRe.FindAllString(lines[1], -1)
	for i := range durs {
		dn, err := strconv.Atoi(durs[i])
		if err != nil {
			return 0, 0, err
		}
		rn, err := strconv.Atoi(records[i])
		if err != nil {
			return 0, 0, err
		}
		races = append(races, race{
			duration: dn,
			record:   rn,
		})
	}
	part1 := 1
	for _, race := range races {
		part1 *= solveEquation(race.duration, race.record)
	}

	duration, err := strconv.Atoi(strings.Join(durs, ""))
	if err != nil {
		return 0, 0, err
	}
	record, err := strconv.Atoi(strings.Join(records, ""))
	if err != nil {
		return 0, 0, err
	}
	part2 := solveEquation(duration, record)

	return part1, part2, nil
}

func init() {
	runner.Register(2023, 6, solver{})
	runner.Register(2023, 6, solverAnalytic{}, "analytic")
}
