package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const sessionFile = "scripts/session.cookie"

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: go run download.go <day>\n")
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

	// Check if input file already exists
	inputFile := filepath.Join(dayDir, pkgName, "input.txt")
	if _, err := os.Stat(inputFile); err == nil {
		fmt.Printf("✓ Input already exists: %s\n", inputFile)
		return
	}

	// Load session cookie
	cookie, err := loadSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: not logged in. Run 'just login' first\n")
		os.Exit(1)
	}

	// Fetch the input
	input, err := fetchInput(dayNum, cookie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching input: %v\n", err)
		os.Exit(1)
	}

	// Ensure directory exists
	pkgDir := filepath.Join(dayDir, pkgName)
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Write the input file
	if err := os.WriteFile(inputFile, []byte(input), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing input file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Downloaded input for day %s\n", dayNum)
}

func loadSession() (string, error) {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func fetchInput(dayNum, cookie string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%s/input", dayNum)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	content := string(body)
	if strings.Contains(content, "Please log in") {
		return "", fmt.Errorf("session cookie invalid or expired")
	}

	return content, nil
}
