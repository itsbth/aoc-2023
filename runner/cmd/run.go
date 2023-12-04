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

type RunConfig struct {
	Year            int
	Day             int
	Part            int    // 0 = both, 1 = part 1, 2 = part 2
	Variant         string // "default"
	CopyToClipboard int    // 0 = none, 1 = part 1, 2 = part 2
	Sample          int
	SampleFile      string
	ListVariants    bool
	CacheDir        string
	Token           string
	TokenFilePath   string
}

var config RunConfig

func printVariants(cmd *cobra.Command, year, day int) {
	cmd.Printf("Variants for %d-%d:\n", year, day)
	for _, v := range runner.GetVariants(year, day) {
		cmd.Printf("- %s\n", v)
	}
}

func getToken() (string, error) {
	if config.Token != "" {
		return config.Token, nil
	}

	if config.TokenFilePath != "" {
		f, err := os.Open(config.TokenFilePath)
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
		if config.ListVariants {
			printVariants(cmd, config.Year, config.Day)
			return
		}

		solver := runner.Get(config.Year, config.Day, config.Variant)
		if solver == nil {
			cmd.PrintErrf("No solver found for %d-%d-%s\n", config.Year, config.Day, config.Variant)
			printVariants(cmd, config.Year, config.Day)
			return
		}

		inputPath := path.Join(config.CacheDir, "input", fmt.Sprintf("%d-%d.txt", config.Year, config.Day))
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			token, err := getToken()
			if err != nil {
				cmd.PrintErrf("Error getting token: %v\n", err)
				return
			}

			content, err := DownloadInput(config.Year, config.Day, token)
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
		if config.SampleFile != "" {
			inputPath = config.SampleFile
		}
		f, err := os.Open(inputPath)
		if err != nil {
			cmd.PrintErrf("Error opening input file: %v\n", err)
			return
		}
		defer f.Close()

		start := time.Now()
		part1, part2, err := solver.Solve(f)
		duration := time.Since(start)
		if err != nil {
			cmd.PrintErrf("Error solving: %v\n", err)
			return
		}
		cmd.Printf("Solved in %s\n", duration)

		if config.Part == 0 || config.Part == 1 {
			cmd.Printf("Part 1: %d\n", part1)
		}
		if config.Part == 0 || config.Part == 2 {
			cmd.Printf("Part 2: %d\n", part2)
		}

		if config.CopyToClipboard == 1 {
			err = copyStringToClipboard(fmt.Sprintf("%d", part1))
			if err != nil {
				cmd.PrintErrf("Error copying to clipboard: %v\n", err)
				return
			}
		}

		if config.CopyToClipboard == 2 {
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
	year := now.Year()
	day := now.Day()

	runCmd.Flags().IntVarP(&config.Year, "year", "y", year, "Year to run")
	runCmd.Flags().IntVarP(&config.Day, "day", "d", day, "Day to run")
	runCmd.Flags().IntVarP(&config.Part, "part", "p", 0, "Part to run")
	runCmd.Flags().IntVarP(&config.Sample, "sample", "s", 0, "Run on sample input")
	runCmd.Flags().StringVarP(&config.SampleFile, "sample-file", "i", "", "Sample input file")
	runCmd.Flags().IntVarP(&config.CopyToClipboard, "copy", "x", 0, "Copy result to clipboard (1 = part 1, 2 = part 2)")
	runCmd.Flags().StringVarP(&config.Variant, "variant", "v", "default", "Variant to run")
	runCmd.Flags().BoolVarP(&config.ListVariants, "list-variants", "l", false, "List variants")
	runCmd.Flags().StringVarP(&config.CacheDir, "cache-dir", "c", ".cache", "Cache directory")

	// token (or token file) is required for downloading inputs
	runCmd.Flags().StringVarP(&config.Token, "token", "t", "", "Session token")
	runCmd.Flags().StringVarP(&config.TokenFilePath, "token-file", "f", ".sessionid", "Session token file")
}
