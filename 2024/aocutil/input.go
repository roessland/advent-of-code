package aocutil

import (
	"bytes"
	"fmt"
	"io/fs"
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
	if len(matches) == 0 {
		return nil
	}
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

// GetIntsInStringLines parses an input string into [][]int, where each line is a slice of ints.
func GetIntsInStringLines(f string) [][]int {
	ints := make([][]int, 0)
	for _, line := range strings.Split(strings.TrimSuffix(f, "\n"), "\n") {
		ints = append(ints, GetIntsInString(line))
	}
	return ints
}

func FSGetIntsInStringLines(dirFS fs.FS, fileName string) [][]int {
	buf := FSReadFile(dirFS, fileName)
	return GetIntsInStringLines(buf)
}

// ReadFileAsInts as ints returns a [][]int from a file, where each line is a []int.
// Any non-integer characters are ignored. Examples:
//
//	"1, 2, 3" -> [[1,2,3]],
//	1\n1.1" -> [[1], [1], [1]]
func ReadFileAsInts(dirFS fs.FS, fileName string) [][]int {
	f := FSReadFile(dirFS, fileName)
	ints := make([][]int, 0)
	for _, line := range strings.Split(strings.TrimSuffix(f, "\n"), "\n") {
		ints = append(ints, GetIntsInString(line))
	}
	return ints
}

// ReadFile wraps string(os.ReadFile) and panics on error.
// Also removes trailing newline.
func ReadFile(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %w", err))
	}
	return strings.TrimSuffix(string(f), "\n")
}

func FSReadFile(dirFS fs.FS, filename string) string {
	dirFS_ := unwrapReadFileFS(dirFS)
	f, err := dirFS_.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %w", err))
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

func FSReadLinesAsBytes(dirFS fs.FS, filename string) [][]byte {
	dirFS_ := unwrapReadFileFS(dirFS)
	f, err := dirFS_.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.Split(bytes.TrimSuffix(f, []byte("\n")), []byte("\n"))
}

func PadMap(m [][]byte, padding int, c byte) [][]byte {
	height := len(m)
	width := len(m[0])
	height_ := height + 2*padding
	width_ := width + 2*padding

	m_ := make([][]byte, height_)

	for i := 0; i < padding; i++ {
		m_[i] = bytes.Repeat([]byte{c}, width_)
		m_[height_-i-1] = bytes.Repeat([]byte{c}, width_)
	}

	for y_ := padding; y_ < height_-padding; y_++ {
		y := y_ - padding
		m_[y_] = make([]byte, width_)
		copy(m_[y_][padding:], m[y])
		for p := 0; p < padding; p++ {
			m_[y_][p] = c
			m_[y_][width_-p-1] = c
		}
	}

	return m_
}

func PrintCharMap(m [][]byte) {
	for _, row := range m {
		fmt.Println(string(row))
	}
}

func unwrapReadFileFS(dirFS fs.FS) fs.ReadFileFS {
	if rfs, ok := dirFS.(fs.ReadFileFS); ok {
		return rfs
	}
	panic("not a fs.ReadFileFS")
}
