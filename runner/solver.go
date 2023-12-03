package runner

import (
	"fmt"
	"io"
)

type Solver interface {
	// Assuming that a 0 return is unimplemented *should* be safe
	Solve(input io.Reader) (int, int, error)
}

var (
	registry = make(map[string]Solver)
	variants = make(map[string][]string)
)

func Register(year int, day int, solver Solver, name ...string) {
	registerName := "default"
	if len(name) > 0 {
		registerName = name[0]
	}

	registry[fmt.Sprintf("%d-%d-%s", year, day, registerName)] = solver
	key := fmt.Sprintf("%d-%d", year, day)
	if _, ok := variants[key]; !ok {
		variants[key] = make([]string, 0)
	}
	variants[key] = append(variants[key], registerName)
}

func Get(year int, day int, name ...string) Solver {
	getName := "default"
	if len(name) > 0 {
		getName = name[0]
	}

	return registry[fmt.Sprintf("%d-%d-%s", year, day, getName)]
}

func GetVariants(year int, day int) []string {
	return variants[fmt.Sprintf("%d-%d", year, day)]
}
