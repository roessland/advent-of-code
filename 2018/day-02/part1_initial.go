package main

import "fmt"
import "strconv"
import "log"
import "bufio"
import "os"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "github.com/roessland/gopkg/disjointset"

func wordCount(str string) map[rune]int {
	freq := map[rune]int{}
	for _, r := range str {
		freq[r]++
	}
	return freq
}

func main() {
	words := []string{}
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		words = append(words, line)
	}

	hasTwoSum := 0
	hasThreeSum := 0
	for _, word := range words {
		hasTwo := false
		hasThree := false
		freqs := wordCount(word)
		for _, freq := range freqs {
			if freq == 2 {
				hasTwo = true
			}
			if freq == 3 {
				hasThree = true
			}
		}
		if hasTwo {
			hasTwoSum++
		}
		if hasThree {
			hasThreeSum++
		}
	}
	fmt.Println(hasTwoSum * hasThreeSum)
}

var _ = fmt.Println
var _ = strconv.Atoi
var _ = log.Fatal
var _ = bufio.NewScanner // (os.Stdin) -> scanner.Scan(), scanner.Text()
var _ = os.Stdin
var _ = strings.Split    // "str str" -> []string{"str", "str"}
var _ = ioutil.ReadFile  // ("input.txt") -> (buf, err)
var _ = csv.NewReader    // (os.Stdin)
var _ = disjointset.Make // (10) -> ds. ds.Union(a,b), ds.Connected(a,b), ds.Count

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
