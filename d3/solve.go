package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func isSymbol(b byte) bool {
	return b != '.'
}

const MAX_LINE = 140

func main() {
	inputFile := "input"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	adj := make([][]bool, MAX_LINE)
	for i := 0; i < MAX_LINE; i++ {
		adj[i] = make([]bool, MAX_LINE)
	}
	type num struct {
		x, y int
		len  int
		val  int
	}
	type gear struct {
		x, y int
	}

	nums := make([]num, 0)
	gears := make([]gear, 0)

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		bytes := scanner.Bytes()
		for i := 0; i < len(bytes); i++ {
			if bytes[i] >= '0' && bytes[i] <= '9' {
				// find number
				start := i
				for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
					i++
				}
				val, err := strconv.Atoi(string(bytes[start:i]))
				if err != nil {
					log.Fatalf("failed to parse number: %v", err)
				}
				nums = append(nums, num{
					x:   start,
					y:   line,
					len: i - start,
					val: val,
				})
				i--
			} else if isSymbol(bytes[i]) {
				// mark 3x3 as adjecent
				for j := -1; j <= 1; j++ {
					for k := -1; k <= 1; k++ {
						// make sure it's inside 140x140 (MAX_LINE)
						if line+j >= 0 && line+j < MAX_LINE && i+k >= 0 && i+k < MAX_LINE {
							adj[line+j][i+k] = true
						}
					}
				}
				if bytes[i] == '*' {
					gears = append(gears, gear{
						x: i,
						y: line,
					})
				}
			}
		}
		line++
	}
	sum := 0
	// sum all numbers overlapping adj
	for _, n := range nums {
		for j := n.x; j < n.x+n.len; j++ {
			if adj[n.y][j] {
				sum += n.val
				break
			}
		}
	}
	sumRatio := 0
	// sum all gear ratios (for ever gear with exactly 2 overlapping numbers, multiply them)
	for _, g := range gears {
		count := 0
		val := 1
		for _, n := range nums {
			ld := g.y - n.y
			if (ld >= -1 && ld <= 1) && (g.x >= n.x-1 && g.x < n.x+n.len+1) {
				count++
				val *= n.val
			}
		}
		if count == 2 {
			sumRatio += val
			log.Printf("DBG: found %d at %d,%d", val, g.x, g.y)
		}
	}
	log.Printf("Part 1: %d", sum)
	log.Printf("Part 2: %d", sumRatio)
}
