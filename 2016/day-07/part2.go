package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IP struct {
	ABAs map[string]bool
	BABs map[string]bool
}

func AddTriplesIfABA(s string, triples map[string]bool) {
	for i := 0; i <= len(s)-3; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			triples[string([]byte{s[i], s[i+1], s[i+2]})] = true
		}
	}
}

func (ip IP) SupportsSSL() bool {
	for aba, _ := range ip.ABAs {
		bab := string([]byte{aba[1], aba[0], aba[1]})
		if ip.BABs[bab] {
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
		ip := IP{make(map[string]bool), make(map[string]bool)}
		for i, part := range parts {
			if i%2 == 0 {
				AddTriplesIfABA(part, ip.ABAs)
			} else {
				AddTriplesIfABA(part, ip.BABs)
			}
		}
		if ip.SupportsSSL() {
			howMany++
		}
	}
	fmt.Println(howMany)
}
