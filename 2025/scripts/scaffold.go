package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"text/template"
)

const mainTemplate = `// Package {{.PkgName}} solves AoC 2025 Day {{.DayNum}}
package {{.PkgName}}

import (
	"embed"

	"github.com/roessland/advent-of-code/2025/aocutil"
)

//go:embed input*.txt
var InputFS embed.FS

func ReadInput(inputName string) [][]byte {
	return aocutil.FSReadLinesAsBytes(InputFS, inputName)
}

func Part1(input [][]byte) int {
	// TODO: Implement Part 1
	return 0
}

func Part2(input [][]byte) int {
	// TODO: Implement Part 2
	return 0
}
`

const testTemplate = `package {{.PkgName}}_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2025/{{.DayDir}}/{{.PkgName}}"
	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	// TODO: Replace with actual test input filename once puzzle input is available
	// inputEx := {{.PkgName}}.ReadInput("input-ex1.txt")
	// require.Greater(t, len(inputEx), 0)

	_ = t
}
`

type TemplateData struct {
	DayNum   string // "5"
	DayPad   string // "05"
	DayDir   string // "day-05"
	PkgName  string // "day05"
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: go run scaffold.go <day>\n")
		os.Exit(1)
	}

	dayNum := args[0]

	// Validate input
	if _, err := strconv.Atoi(dayNum); err != nil {
		fmt.Fprintf(os.Stderr, "Error: day must be a number\n")
		os.Exit(1)
	}

	// Pad with zero
	dayPad := fmt.Sprintf("%02s", dayNum)
	dayDir := fmt.Sprintf("day-%s", dayPad)
	pkgName := fmt.Sprintf("day%s", dayPad)

	data := TemplateData{
		DayNum:  dayNum,
		DayPad:  dayPad,
		DayDir:  dayDir,
		PkgName: pkgName,
	}

	// Create directory
	pkgDir := filepath.Join(dayDir, pkgName)
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Parse and execute main template
	mainTmpl, err := template.New("main").Parse(mainTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing main template: %v\n", err)
		os.Exit(1)
	}

	mainFile := filepath.Join(pkgDir, fmt.Sprintf("%s.go", pkgName))
	f, err := os.Create(mainFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating main file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := mainTmpl.Execute(f, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing main template: %v\n", err)
		os.Exit(1)
	}

	// Parse and execute test template
	testTmpl, err := template.New("test").Parse(testTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing test template: %v\n", err)
		os.Exit(1)
	}

	testFile := filepath.Join(pkgDir, fmt.Sprintf("%s_test.go", pkgName))
	f, err = os.Create(testFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating test file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := testTmpl.Execute(f, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing test template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Scaffolded %s with %s package\n", dayDir, pkgName)

	// Try to download input
	cmd := exec.Command("go", "run", "scripts/download.go", dayNum)
	if err := cmd.Run(); err != nil {
		// Download errors are non-fatal (might not be logged in yet)
	}
}
