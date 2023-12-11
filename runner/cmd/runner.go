package main

import (
	"github.com/spf13/cobra"

	// START IMPORTS
	_ "github.com/itsbth/aoc-2023/d1"
	_ "github.com/itsbth/aoc-2023/d10"
	_ "github.com/itsbth/aoc-2023/d11"
	_ "github.com/itsbth/aoc-2023/d2"
	_ "github.com/itsbth/aoc-2023/d3"
	_ "github.com/itsbth/aoc-2023/d4"
	_ "github.com/itsbth/aoc-2023/d5"
	_ "github.com/itsbth/aoc-2023/d6"
	_ "github.com/itsbth/aoc-2023/d7"
	_ "github.com/itsbth/aoc-2023/d8"
	_ "github.com/itsbth/aoc-2023/d9"
	// END IMPORTS
)

var rootCmd = &cobra.Command{
	Use:   "advent",
	Short: "Advent of Code",
	Long:  `Advent of Code`,
}

func main() {
	rootCmd.Execute()
}
