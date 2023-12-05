package d5

import (
	"bufio"
	"io"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

func (solver) Solve(input io.Reader) (int, int, error) {
	var seeds []int
	var mappings []mapping
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// seeds: 79 14 55 13
		if strings.HasPrefix(line, "seeds: ") {
			for _, seed := range strings.Split(line[7:], " ") {
				s, err := strconv.Atoi(seed)
				if err != nil {
					return 0, 0, err
				}
				seeds = append(seeds, s)
			}
			continue
		}
		// seed-to-soil map:
		if strings.HasSuffix(line, "map:") {
			mappings = append(mappings, mapping{})
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			continue
		}

		dest, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, err
		}
		source, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0, err
		}
		length, err := strconv.Atoi(parts[2])
		if err != nil {
			return 0, 0, err
		}
		mappings[len(mappings)-1].addEntry(entry{
			start:  source,
			end:    source + length,
			offset: dest - source,
		})
	}
	for i := range mappings {
		mappings[i].sort()
	}
	var locations []int
	for _, seed := range seeds {
		for _, mapping := range mappings {
			seed = mapping.translate(seed)
		}
		locations = append(locations, seed)
	}
	slices.Sort(locations)

	var ranges []span
	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, span{
			start: seeds[i],
			end:   seeds[i] + seeds[i+1] + 1,
		})
	}
	tasks := make(chan span)
	results := make(chan []span)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for task := range tasks {
				ranges := []span{task}
				for _, mapping := range mappings {
					newRanges := []span{}
					for _, r := range ranges {
						newRanges = append(newRanges, mapping.translateRange(r)...)
					}
					ranges = newRanges
				}
				results <- ranges
			}
		}()
	}
	for _, r := range ranges {
		tasks <- r
	}
	close(tasks)
	var solved []span
	for i := 0; i < len(ranges); i++ {
		solved = append(solved, <-results...)
	}
	slices.SortFunc(solved, func(a, b span) int {
		return a.start - b.start
	})
	return locations[0], solved[0].start, nil
}

func init() {
	runner.Register(2023, 5, solver{})
}

type entry struct {
	start, end int
	offset     int
}

type mapping struct {
	entries []entry
}

func (m *mapping) addEntry(e entry) {
	m.entries = append(m.entries, e)
}

func (m *mapping) sort() {
	slices.SortFunc(m.entries, func(a, b entry) int {
		return a.start - b.start
	})
}

// map source to dest
// if it matches an entry (ie source >= entry.source && source < entry.source + entry.length)
// return dest + (source - entry.source)
// else return source
// assumes entries are sorted, so bail early if source > entry.source + entry.length
func (m *mapping) translate(source int) int {
	for _, e := range m.entries {
		if source < e.start {
			break
		}
		if source >= e.end {
			continue
		}
		return e.offset + source
	}
	return source
}

func (m *mapping) translateRange(from span) []span {
	var out []span
	start := from.start
	end := from.end

	for _, e := range m.entries {
		if start >= end {
			break
		}
		if start >= e.end || end <= e.start {
			continue
		}
		// some overlap. split into before, overlap, after (the first and last of which may be empty)
		if start < e.start {
			out = append(out, span{
				start: start,
				end:   e.start,
			})
		}
		out = append(out, span{
			start: max(start, e.start) + e.offset,
			end:   min(end, e.end) + e.offset,
		})
		start = max(start, e.end)
	}
	if start <= end {
		out = append(out, span{
			start: start,
			end:   end,
		})
	}
	return out
}

type span struct {
	start, end int
}

func (s span) String() string {
	return strconv.Itoa(s.start) + "-" + strconv.Itoa(s.end)
}
