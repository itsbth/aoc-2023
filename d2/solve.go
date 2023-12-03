package d2

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/itsbth/aoc-2023/runner"
)

type game struct {
	id       int
	revealed []map[string]int
}

func parseGame(line string) game {
	g := game{}
	parts := strings.SplitN(line, ": ", 2)
	// "Game NN"
	id, err := strconv.ParseInt(parts[0][5:], 10, 64)
	if err != nil {
		log.Fatalf("failed to parse game id: %v", err)
	}
	g.id = int(id)
	sets := strings.Split(parts[1], "; ")
	for _, set := range sets {
		m := make(map[string]int)
		for _, color := range strings.Split(set, ", ") {
			parts := strings.Split(color, " ")
			count, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				log.Fatalf("failed to parse color count: %v", err)
			}
			m[parts[1]] = int(count)
		}
		g.revealed = append(g.revealed, m)
	}
	return g
}

// faster parser(?)
func parseGame2(line string) (game, error) {
	game := game{}
	idx := strings.IndexByte(line, ':')
	if idx == -1 {
		return game, errors.New("invalid game")
	}
	id, err := strconv.ParseInt(line[5:idx], 10, 64)
	if err != nil {
		return game, err
	}
	game.id = int(id)

	for idx < len(line) {
		m := make(map[string]int)
		for idx < len(line) {
			idx += 2
			space := strings.IndexByte(line[idx:], ' ')
			if space == -1 {
				return game, errors.New("invalid color count")
			}
			count, err := strconv.ParseInt(line[idx:idx+space], 10, 64)
			if err != nil {
				return game, err
			}
			idx += space + 1
			start := idx
			for idx < len(line) {
				if line[idx] == ',' || line[idx] == ';' {
					break
				}
				idx++
			}

			color := line[start:idx]
			m[color] = int(count)
			if idx == len(line) || line[idx] == ';' {
				break
			}
		}
		game.revealed = append(game.revealed, m)
	}
	return game, nil
}

func (g *game) isPossible(limits map[string]int) bool {
	for _, m := range g.revealed {
		for color, count := range m {
			if count > limits[color] {
				return false
			}
		}
	}
	return true
}

func (g *game) power() int {
	power := 1
	max := map[string]int{}
	for _, m := range g.revealed {
		for color, count := range m {
			if count > max[color] {
				max[color] = count
			}
		}
	}
	for _, count := range max {
		power *= count
	}
	return power
}

type solver struct{}

var _ runner.Solver = solver{}

func (solver) Solve(input io.Reader) (int, int, error) {
	games := make([]game, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		g, err := parseGame2(scanner.Text())
		if err != nil {
			return 0, 0, err
		}
		games = append(games, g)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 12 red cubes, 13 green cubes, and 14 blue cubes
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	power := 0
	for _, g := range games {
		if g.isPossible(limits) {
			sum += g.id
		}
		power += g.power()
	}

	return sum, power, nil
}

func init() {
	runner.Register(2023, 2, solver{})
}
