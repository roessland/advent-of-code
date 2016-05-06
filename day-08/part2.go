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
		memory_len := len(line)
		literal_len += strings.Count(line, `"`) + 2
		literal_len += strings.Count(line, `\`)

		fmt.Printf("%v %v\n", line, literal_len)
		total += literal_len
		total -= memory_len
	}
	fmt.Printf("%v\n", total)
}
