package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/itsbth/aoc-2023/runner"
	"github.com/spf13/cobra"
)

var (
	year            int
	day             int
	part            int    = 0 // 0 = both, 1 = part 1, 2 = part 2
	variant         string = "default"
	copyToClipboard int    // 0 = none, 1 = part 1, 2 = part 2
	sample          int
	sampleFile      string
	listVariants    bool
	cacheDir        string
	token           string
	tokenFilePath   string
)

func printVariants(cmd *cobra.Command, year, day int) {
	cmd.Printf("Variants for %d-%d:\n", year, day)
	for _, v := range runner.GetVariants(year, day) {
		cmd.Printf("- %s\n", v)
	}
}

func getToken() (string, error) {
	if token != "" {
		return token, nil
	}

	if tokenFilePath != "" {
		f, err := os.Open(tokenFilePath)
		if err != nil {
			return "", err
		}
		defer f.Close()
		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, f)
		if err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	return "", fmt.Errorf("no token specified")
}

func copyStringToClipboard(content string) error {
	// only implement darwin pbcopy for now

	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(content)
	return cmd.Run()
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a solution",
	Long:  `Run a solution`,
	Run: func(cmd *cobra.Command, args []string) {
		if listVariants {
			printVariants(cmd, year, day)
			return
		}

		solver := runner.Get(year, day, variant)
		if solver == nil {
			cmd.PrintErrf("No solver found for %d-%d-%s\n", year, day, variant)
			printVariants(cmd, year, day)
			return
		}

		inputPath := path.Join(cacheDir, "input", fmt.Sprintf("%d-%d.txt", year, day))
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			token, err := getToken()
			if err != nil {
				cmd.PrintErrf("Error getting token: %v\n", err)
				return
			}

			content, err := DownloadInput(year, day, token)
			if err != nil {
				cmd.PrintErrf("Error downloading input: %v\n", err)
				return
			}

			err = os.MkdirAll(path.Dir(inputPath), 0o755)
			if err != nil {
				cmd.PrintErrf("Error creating cache directory: %v\n", err)
				return
			}

			f, err := os.Create(inputPath)
			if err != nil {
				cmd.PrintErrf("Error creating input file: %v\n", err)
				return
			}
			defer f.Close()

			_, err = f.WriteString(content)
			if err != nil {
				cmd.PrintErrf("Error writing input file: %v\n", err)
				return
			}
		}
		if sampleFile != "" {
			inputPath = sampleFile
		}
		f, err := os.Open(inputPath)
		if err != nil {
			cmd.PrintErrf("Error opening input file: %v\n", err)
			return
		}
		defer f.Close()

		part1, part2, err := solver.Solve(f)
		if err != nil {
			cmd.PrintErrf("Error solving: %v\n", err)
			return
		}

		if part == 0 || part == 1 {
			cmd.Printf("Part 1: %d\n", part1)
		}
		if part == 0 || part == 2 {
			cmd.Printf("Part 2: %d\n", part2)
		}

		if copyToClipboard == 1 {
			err = copyStringToClipboard(fmt.Sprintf("%d", part1))
			if err != nil {
				cmd.PrintErrf("Error copying to clipboard: %v\n", err)
				return
			}
		}

		if copyToClipboard == 2 {
			err = copyStringToClipboard(fmt.Sprintf("%d", part2))
			if err != nil {
				cmd.PrintErrf("Error copying to clipboard: %v\n", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	now := time.Now()
	year = now.Year()
	day = now.Day()

	runCmd.Flags().IntVarP(&year, "year", "y", year, "Year to run")
	runCmd.Flags().IntVarP(&day, "day", "d", day, "Day to run")
	runCmd.Flags().IntVarP(&part, "part", "p", 0, "Part to run")
	runCmd.Flags().IntVarP(&sample, "sample", "s", 0, "Run on sample input")
	runCmd.Flags().StringVarP(&sampleFile, "sample-file", "i", "", "Sample input file")
	runCmd.Flags().IntVarP(&copyToClipboard, "copy", "x", 0, "Copy result to clipboard (1 = part 1, 2 = part 2)")
	runCmd.Flags().StringVarP(&variant, "variant", "v", "default", "Variant to run")
	runCmd.Flags().BoolVarP(&listVariants, "list-variants", "l", false, "List variants")
	runCmd.Flags().StringVarP(&cacheDir, "cache-dir", "c", ".cache", "Cache directory")

	// token (or token file) is required for downloading inputs
	runCmd.Flags().StringVarP(&token, "token", "t", "", "Session token")
	runCmd.Flags().StringVarP(&tokenFilePath, "token-file", "f", ".sessionid", "Session token file")
}
