package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HasAbba(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	howMany := 0
	for scanner.Scan() {
		parts := strings.FieldsFunc(scanner.Text(), func(c rune) bool {
			return c == '[' || c == ']'
		})
		supportsTLS := false
		for i, part := range parts {
			if i%2 == 0 {
				supportsTLS = supportsTLS || HasAbba(part)
			} else {
				if HasAbba(part) {
					supportsTLS = false
					break
				}
			}
		}
		if supportsTLS {
			howMany++
		}
	}
	fmt.Println("Supports TLS: ", howMany)
}
