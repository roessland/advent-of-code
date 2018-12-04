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

func countEqual(a, b string) bool {
	cn := 0
	for i := 0; i < len(b); i++ {
		if a[i] == b[i] {
			cn++
		}
	}
	return cn == len(a)-1
}

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

	for _, word1 := range words {
		for _, word2 := range words {
			if countEqual(word1, word2) {
				fmt.Println(word1, word2)
			}
		}
	}
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
