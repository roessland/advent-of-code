package aocutil

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// Atoi wraps strconv.Atoi and panics on error.
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func GetIntsInString(line string) []int {
	return GetNumsInString[int](line)
}

func GetNumsInString[N ~int](line string) []N {
	lineRe := regexp.MustCompile(`(^|[^a-zA-Z])(-?\d+)`)
	matches := lineRe.FindAllStringSubmatch(line, -1)
	nums := make([]N, 0, len(matches)-1)
	for _, match := range matches {
		n, err := strconv.Atoi(strings.TrimSpace(match[2]))
		if err != nil {
			spew.Dump(matches)
			panic(err)
		}
		nums = append(nums, N(n))
	}
	return nums
}

// ReadFile wraps string(os.ReadFile) and panics on error.
// Also removes trailing newline.
func ReadFile(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(f), "\n")
}

func ReadFileAsBytes(filename string) []byte {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.TrimSuffix(f, []byte("\n"))
}

// ReadLines wraps strings.Split(string(os.ReadFile)) and panics on error.
func ReadLines(filename string) []string {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSuffix(string(f), "\n"), "\n")
}

// ReadLinesAsBytes is the same as ReadLines but returns [][]byte instead of []string.
func ReadLinesAsBytes(filename string) [][]byte {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.Split(bytes.TrimSuffix(f, []byte("\n")), []byte("\n"))
}
