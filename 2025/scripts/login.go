package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const sessionFile = "scripts/session.cookie"
const aocURL = "https://adventofcode.com/2025/day/1/input"

func main() {
	// Check if session cookie already exists
	if cookie, err := loadSession(); err == nil && cookie != "" {
		if verifySession(cookie) {
			fmt.Println("✓ Already logged in!")
			return
		}
		fmt.Println("⚠ Session cookie expired or invalid. Please log in again.")
	}

	fmt.Println("\nTo get your session cookie:")
	fmt.Println("1. Go to: https://adventofcode.com/2025/day/1")
	fmt.Println("2. Open DevTools (F12)")
	fmt.Println("3. Firefox: Storage > Cookies > adventofcode.com")
	fmt.Println("   Chrome:  Application > Cookies > adventofcode.com")
	fmt.Println("4. Copy the 'session' cookie value")
	fmt.Println("   e.g. 53616c7465642068616e64202d2068616e64202d...")
	fmt.Println()

	cookie := promptForCookie()
	if cookie == "" {
		fmt.Fprintf(os.Stderr, "Error: no cookie provided\n")
		os.Exit(1)
	}

	// Verify the cookie works
	if !verifySession(cookie) {
		fmt.Fprintf(os.Stderr, "Error: invalid session cookie\n")
		os.Exit(1)
	}

	// Save the cookie
	if err := saveSession(cookie); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving session: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Logged in successfully!")
}

func loadSession() (string, error) {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func saveSession(cookie string) error {
	return os.WriteFile(sessionFile, []byte(cookie), 0600)
}

func promptForCookie() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Paste your session cookie: ")
	cookie, _ := reader.ReadString('\n')
	return strings.TrimSpace(cookie)
}

func verifySession(cookie string) bool {
	// Try to fetch day 1 input with the cookie
	req, err := http.NewRequest("GET", aocURL, nil)
	if err != nil {
		return false
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	// Check if we got actual content (not a redirect or error page)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	// If we got content, it should be the puzzle input (multiple lines of digits)
	return len(body) > 0 && !strings.Contains(string(body), "Please log in")
}
