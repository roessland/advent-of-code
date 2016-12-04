package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var re = regexp.MustCompile(`([a-z\-]+)\-(\d+)\[([a-z]+)\]`)

func Extract(line string) (name string, number int, checksum string) {
	match := re.FindStringSubmatch(line)
	name = match[1]
	number, _ = strconv.Atoi(match[2])
	checksum = match[3]
	return
}

type Pair struct {
	Fst int
	Snd rune
}

type ByFst []Pair

func (s ByFst) Len() int {
	return len(s)
}

func (s ByFst) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByFst) Less(i, j int) bool {
	if s[i].Fst == s[j].Fst {
		return s[i].Snd < s[j].Snd
	} else {
		return s[i].Fst > s[j].Fst
	}
}

func Checksum(name string) string {
	freqsMap := make(map[rune]int)
	for _, char := range name {
		if unicode.IsLetter(char) {
			freqsMap[char]++
		}
	}
	freqsList := make([]Pair, 0)
	for char, freq := range freqsMap {
		freqsList = append(freqsList, Pair{freq, char})
	}
	sort.Sort(ByFst(freqsList))
	l := freqsList[:5]
	return fmt.Sprintf("%c%c%c%c%c", l[0].Snd, l[1].Snd, l[2].Snd, l[3].Snd, l[4].Snd)
}

var cipher = map[rune]rune{
	'a': 'b',
	'b': 'c',
	'c': 'd',
	'd': 'e',
	'e': 'f',
	'f': 'g',
	'g': 'h',
	'h': 'i',
	'i': 'j',
	'j': 'k',
	'k': 'l',
	'l': 'm',
	'm': 'n',
	'n': 'o',
	'o': 'p',
	'p': 'q',
	'q': 'r',
	'r': 's',
	's': 't',
	't': 'u',
	'u': 'v',
	'v': 'w',
	'w': 'x',
	'x': 'y',
	'y': 'z',
	'z': 'a',
	' ': '-',
	'-': ' ',
}

func DecryptChar(char rune, steps int) rune {
	for i := 0; i < steps; i++ {
		char = cipher[char]
	}
	return char
}

func Decrypt(name string, steps int) string {
	decrypted := make([]rune, len(name))
	for i := 0; i < len(name); i++ {
		decrypted[i] = DecryptChar(rune(name[i]), steps)
	}
	return string(decrypted)
}

func main() {

	buf, _ := ioutil.ReadFile("input.txt")
	sum := 0
	for _, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		name, number, checksum := Extract(line)
		if checksum == Checksum(name) {
			sum += number
			realName := Decrypt(name, number)
			if strings.Contains(realName, "north") {
				fmt.Println(realName, number)
			}
		}
	}
	fmt.Println(sum)
}
