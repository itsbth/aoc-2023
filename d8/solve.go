package d8

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

var lazyRe = regexp.MustCompile(`([A-Z0-9]{3})`)

type xr struct {
	left, right string
}

type camel struct {
	xrs        *map[string]xr
	directions string
	pos        string
	ptr        int // where in directions we are
}

func (c *camel) step() {
	dir := c.directions[c.ptr]
	c.ptr = (c.ptr + 1) % len(c.directions)
	xr := (*c.xrs)[c.pos]
	if dir == 'L' {
		c.pos = xr.left
	} else {
		c.pos = xr.right
	}
}

func (c *camel) detectLoop() (int, int, int) {
	seen := make(map[string]int)
	steps := 0
	lastZ := 0
	for {
		key := fmt.Sprintf("%s:%d", c.pos, c.ptr)
		// key := fmt.Sprintf("%s", c.pos)
		if c.pos[2] == 'Z' {
			lastZ = steps
		}
		if cycle, ok := seen[key]; ok {
			for k, v := range seen {
				if v == 0 {
					log.Printf("first: %s, current: %s", k, key)
				}
			}
			return (steps - cycle), cycle - 1, lastZ
		}
		seen[key] = steps
		c.step()
		steps += 1
	}
}

func (solver) Solve(input io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	// LR string
	directions := scanner.Text()

	xrs := make(map[string]xr)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := lazyRe.FindAllString(line, -1)
		xrs[parts[0]] = xr{
			left:  parts[1],
			right: parts[2],
		}
	}
	pos := "AAA"
	cur := 0
	steps := 0
	for pos != "ZZZ" {
		dir := directions[cur]
		cur = (cur + 1) % len(directions)
		steps += 1
		xr := xrs[pos]
		if dir == 'L' {
			pos = xr.left
		} else {
			pos = xr.right
		}
	}
	var ghosts []camel
	for n := range xrs {
		if n[2] == 'A' {
			ghosts = append(ghosts, camel{
				xrs:        &xrs,
				directions: directions,
				pos:        n,
			})
		}
	}
	part2 := 0
	var cycles []int
	var starts []int
	stepSize := 1
	for _, ghost := range ghosts {
		cycle, _, lastZ := ghost.detectLoop()
		cycles = append(cycles, cycle)
		starts = append(starts, lastZ-cycle)
		stepSize = lcm(stepSize, cycle)
	}
	log.Printf("cycles: %+v", cycles)
	log.Printf("starts: %+v", starts)
	log.Printf("stepSize: %d", stepSize)
	cpos := 0
	for {
		cpos += stepSize
		allValid := true
		for idx, cycle := range cycles {
			if (cpos-starts[idx])%cycle != 0 {
				allValid = false
			}
		}
		if allValid {
			part2 = cpos
			break
		}
	}
	return steps, part2, nil
}

func init() {
	runner.Register(2023, 8, solver{})
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
