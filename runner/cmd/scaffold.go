package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var (
	scaffoldDay   int
	scaffoldForce bool
)

//go:embed solve.go.tmpl
var solveTemplate string
var solveGo = template.Must(template.New("solve").Parse(solveTemplate))

const RUNNER_ENTRYPOINT = "runner/cmd/runner.go"

type templateData struct {
	Package string
	Day     int
	Year    int
}

func deduplicateSortedStrings(s []string) []string {
	if len(s) == 0 {
		return s
	}

	deduped := make([]string, 0)
	deduped = append(deduped, s[0])
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			deduped = append(deduped, s[i])
		}
	}
	return deduped
}

func insertImport(file *os.File, pkg string) error {
	// - read whole file
	// - find import block (// START IMPORTS, // END IMPORTS)
	// - extract existing imports
	// - insert new import
	// - sort and deduplicate imports
	// - write back to file

	lines := make([]string, 0)
	imports := make([]string, 0)
	importBlock := false
	importBlockStart := 0
	importBlockEnd := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "\t// START IMPORTS" {
			importBlock = true
			importBlockStart = len(lines)
		} else if line == "\t// END IMPORTS" {
			importBlock = false
			importBlockEnd = len(lines)
		} else if importBlock {
			imports = append(imports, line)
		}
		lines = append(lines, line)
	}

	imports = append(imports, fmt.Sprintf("\t_ \"github.com/itsbth/aoc-2023/%s\"", pkg))

	sort.Strings(imports)
	imports = deduplicateSortedStrings(imports)

	lines = append(lines[:importBlockStart+1], append(imports, lines[importBlockEnd:]...)...)

	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	for _, line := range lines {
		fmt.Fprintln(file, line)
	}

	return nil
}

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Scaffold a solution",
	Long:  `Scaffold a solution`,
	Run: func(cmd *cobra.Command, args []string) {
		// Steps:
		// - Check if directory exists warn-and-clear/error depending on force
		// - Create directory
		// - Create solve.go based on template
		// - modify runner/cmd/runner.go to include the new package

		targetDir := fmt.Sprintf("d%d", scaffoldDay)
		if _, err := os.Stat(targetDir); err == nil {
			if scaffoldForce {
				err := os.RemoveAll(targetDir)
				if err != nil {
					cmd.Printf("Error removing directory: %s\n", err)
					os.Exit(1)
				}
			} else {
				fmt.Printf("Directory %s already exists, use --force to overwrite\n", targetDir)
				os.Exit(1)
			}
		}

		err := os.Mkdir(targetDir, 0o755)
		if err != nil {
			cmd.Printf("Error creating directory: %s\n", err)
			os.Exit(1)
		}

		f, err := os.Create(fmt.Sprintf("%s/solve.go", targetDir))
		if err != nil {
			cmd.Printf("Error creating solve.go: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		err = solveGo.Execute(f, templateData{
			Package: targetDir,
			Day:     scaffoldDay,
			Year:    2023,
		})
		if err != nil {
			cmd.Printf("Error executing template: %s\n", err)
			os.Exit(1)
		}

		f, err = os.OpenFile(RUNNER_ENTRYPOINT, os.O_RDWR, 0o755)
		if err != nil {
			cmd.Printf("Error opening runner entrypoint: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		err = insertImport(f, targetDir)
		if err != nil {
			cmd.Printf("Error inserting import: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)

	scaffoldCmd.Flags().IntVarP(&scaffoldDay, "day", "d", 0, "Day to scaffold")
	scaffoldCmd.Flags().BoolVarP(&scaffoldForce, "force", "f", false, "Force overwrite")
}
