package d10

import (
	"bufio"
	"io"
	"slices"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

// bitmask, UDLR. Only two bits can be set at a time (ie only straight or corners, not junctions)
var connectivity = map[byte]int{
	'|': 0b1100,
	'-': 0b0011,
	'L': 0b1001,
	'J': 0b1010,
	'7': 0b0110,
	'F': 0b0101,
	'.': 0b0000,
}

var (
	DIR_UP    = 0b1000
	DIR_DOWN  = 0b0100
	DIR_LEFT  = 0b0010
	DIR_RIGHT = 0b0001
)

func dirName(dir int) string {
	switch dir {
	case DIR_UP:
		return "up"
	case DIR_DOWN:
		return "down"
	case DIR_LEFT:
		return "left"
	case DIR_RIGHT:
		return "right"
	}
	return "unknown"
}

func flip(dir int) int {
	switch dir {
	case DIR_UP:
		return DIR_DOWN
	case DIR_DOWN:
		return DIR_UP
	case DIR_LEFT:
		return DIR_RIGHT
	case DIR_RIGHT:
		return DIR_LEFT
	}
	return 0
}

type pos struct {
	x, y int
}

func traverse(sx, sy int, dir int, pipes [][]byte) (int, []pos) {
	var loop []pos
	dist := 0
	x, y := sx, sy
	for {
		loop = append(loop, pos{x, y})
		switch dir {
		case DIR_UP:
			y--
		case DIR_DOWN:
			y++
		case DIR_LEFT:
			x--
		case DIR_RIGHT:
			x++
		}
		dist++
		if x < 0 || y < 0 || x >= len(pipes[0]) || y >= len(pipes) {
			return -1, loop
		}
		// log.Printf("x: %d, y: %d, dir: %s [%d], pipe: %c", x, y, dirName(dir), dir, rune(pipes[y][x]))
		if x == sx && y == sy {
			// log.Printf("Found start again")
			return dist, loop
		}
		np := connectivity[pipes[y][x]]
		if np == 0 {
			// log.Printf("No connection")
			return -1, loop
		}
		from := flip(dir)
		if np&from == 0 {
			// log.Printf("No connection from %s", dirName(from))
			// print 3x3 around current for debugging
			// for dy := -2; dy <= 2; dy++ {
			// 	line := ""
			// 	for dx := -2; dx <= 2; dx++ {
			// 		if x+dx < 0 || y+dy < 0 || x+dx >= len(pipes[0]) || y+dy >= len(pipes) {
			// 			line += " "
			// 			continue
			// 		}
			// 		if dx == 0 && dy == 0 {
			// 			line += color.HiRedString("%c", pipes[y+dy][x+dx])
			// 		} else if dx == 0 && dy == -1 && from == DIR_UP {
			// 			line += color.BlueString("%c", pipes[y+dy][x+dx])
			// 		} else if dx == 0 && dy == 1 && from == DIR_DOWN {
			// 			line += color.BlueString("%c", pipes[y+dy][x+dx])
			// 		} else if dx == -1 && dy == 0 && from == DIR_LEFT {
			// 			line += color.BlueString("%c", pipes[y+dy][x+dx])
			// 		} else if dx == 1 && dy == 0 && from == DIR_RIGHT {
			// 			line += color.BlueString("%c", pipes[y+dy][x+dx])
			// 		} else {
			// 			line += string(pipes[y+dy][x+dx])
			// 		}
			// 	}
			// 	// log.Printf("%s", line)
			// }
			return -1, loop
		}
		dir = from ^ np
	}
}

func (solver) Solve(input io.Reader) (int, int, error) {
	var pipes [][]byte
	scanner := bufio.NewScanner(input)
	var start struct {
		x, y int
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		pipes = append(pipes, slices.Clone(line))
		for i, c := range line {
			if c == 'S' {
				start.x = i
				start.y = len(pipes) - 1
			}
		}
	}

	part1 := 0
	var loop []pos
	for dir := 1; dir <= 8; dir <<= 1 {
		// log.Printf("Trying dir %s", dirName(dir))
		dist, l2 := traverse(start.x, start.y, dir, pipes)
		// log.Printf("Dist: %d", dist)
		if dist > 0 {
			part1 += dist
			loop = l2
			break
		}
	}

	next := make(map[pos]int)
	for idx, p := range loop {
		next[p] = idx
	}
	area := 0
	for y := 0; y < len(pipes); y++ {
		inside := false
		for x := 0; x < len(pipes[0]); x++ {
			// if !inside and next[x,y] is set, we're inside
			// set entered to 1 or 2 depending on direction
			// if inside and next[x,y] is not set, determine inside by direction

			if _, ok := next[pos{x, y}]; ok {
				curr := pipes[y][x]
				if curr == '|' || curr == 'L' || curr == 'J' {
					// crosses a vertical pipe; toggle inside
					inside = !inside
				}
			} else if inside {
				area++
				// if _, ok := next[pos{x, y}]; !ok {
				// 	log.Printf("Enclosed area at %d,%d", x, y)
				// }
			}
		}
	}

	return part1 / 2, area, nil
}

func init() {
	runner.Register(2023, 10, solver{})
}
