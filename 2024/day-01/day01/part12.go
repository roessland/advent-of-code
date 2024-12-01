package day01

import (
	"embed"
	"sort"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/mathutil"
)

//go:embed input*.txt
var inputDir embed.FS

type ID int

func Unzip(pairs [][]int) (lefts []ID, rights []ID) {
	for _, pair := range pairs {
		lefts = append(lefts, ID(pair[0]))
		rights = append(rights, ID(pair[1]))
	}
	return lefts, rights
}

func ZipWithIndex(locationIDs []int) []Number {
	numbers := make([]Number, 0, len(locationIDs))
	for i, id := range locationIDs {
		numbers = append(numbers, Number{Index: i, LocationID: id})
	}
	return numbers
}

func SortByID(ids []ID) {
	sort.SliceStable(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
}

type Number struct {
	Index      int
	LocationID int
}

func Part1(inputName string) int {
	rows := aocutil.ReadFileAsInts(inputDir, inputName)
	lefts, rights := Unzip(rows)
	SortByID(lefts)
	SortByID(rights)

	sum := 0
	for i := range rows {
		sum += mathutil.AbsInt(int(rights[i] - lefts[i]))
	}
	return sum
}

func Frequencies(ids []ID) map[ID]int {
	freq := make(map[ID]int)
	for _, n := range ids {
		freq[n]++
	}
	return freq
}

func Part2(inputName string) int {
	rows := aocutil.ReadFileAsInts(inputDir, inputName)
	lefts, rights := Unzip(rows)
	rightFreqs := Frequencies(rights)

	sum := 0
	for _, id := range lefts {
		sum += int(id) * rightFreqs[id]
	}
	return sum
}
