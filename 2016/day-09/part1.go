package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"unicode"
)

func Atoi(s []byte) int {
	num, _ := strconv.Atoi(string(s))
	return num
}

func Decompress(data []byte, w io.Writer) {
	re := regexp.MustCompile(`\((\d+)x(\d+)\)`)
	for len(data) > 0 {
		loc := re.FindSubmatchIndex(data) // 6 indices
		if loc == nil {
			w.Write(data)
			return
		}
		size := Atoi(data[loc[2]:loc[3]])
		times := Atoi(data[loc[4]:loc[5]])

		// From beginning to start of marker
		w.Write(data[:loc[0]])

		// Expand marker
		for i := 0; i < times; i++ {
			w.Write(data[loc[1] : loc[1]+size])
		}

		// For next iteration
		if loc[1]+size < len(data) {
			data = data[loc[1]+size : len(data)]
		} else {
			return
		}
	}
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	var buf bytes.Buffer
	f := bufio.NewWriter(&buf)
	Decompress(data, f)
	f.Flush()
	length := 0
	for _, r := range buf.String() {
		if !unicode.IsSpace(r) {
			length++
		}
	}
	fmt.Println("Final result has", length, "non-space characters")
}
