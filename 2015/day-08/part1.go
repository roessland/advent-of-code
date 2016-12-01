package main

import "fmt"
import "strings"
import "io/ioutil"

func main() {
	buf, _ := ioutil.ReadFile("input.txt")

	total := 0
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		literal_len := len(line)
		memory_len := len(line) - 2
		orig_line := line

		// Replace double backslash
		doubleslashes := strings.Count(line, `\\`)
		line = strings.Replace(line, `\\`, "__", -1)
		memory_len -= 1 * doubleslashes

		// Replace hex characters
		hexes := strings.Count(line, `\x`)
		line = strings.Replace(line, `\x`, "__", -1)
		memory_len -= 3 * hexes

		// Replace backslash
		backslashes := strings.Count(line, `\`)
		memory_len -= 1 * backslashes

		total += literal_len
		total -= memory_len
	}
	fmt.Printf("%v\n", total)
}
