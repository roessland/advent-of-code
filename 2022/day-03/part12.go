package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"unicode"
)

//go:embed input.txt input-ex.txt
var inputFiles embed.FS

func main() {
	backpacks := ReadInput()
	part1(backpacks)
	part2(backpacks)
}

func part1(backpacks []Backpack) {
	totalPri := 0
	for _, b := range backpacks {
		common := b.FindCommonItemInCompartments()
		if unicode.IsLower(common) {
			pri := common - 'a' + 1
			totalPri += int(pri)
		} else if unicode.IsUpper(common) {
			pri := common - 'A' + 27
			totalPri += int(pri)
		} else {
			panic("nope")
		}
	}
	fmt.Println(totalPri)
}

func part2(backpacks []Backpack) {
	totalPri := 0
	for i := 0; i < len(backpacks)-2; i += 3 {
		b1 := backpacks[i+0]
		b2 := backpacks[i+1]
		b3 := backpacks[i+2]
		common := FindCommonItemInSets(b1.Both, b2.Both, b3.Both)
		if unicode.IsLower(common) {
			pri := common - 'a' + 1
			totalPri += int(pri)
		} else if unicode.IsUpper(common) {
			pri := common - 'A' + 27
			totalPri += int(pri)
		} else {
			panic("nope")
		}
	}
	fmt.Println(totalPri)
}

type Backpack struct {
	Both, First, Second map[rune]bool
}

func (b *Backpack) FindCommonItemInCompartments() rune {
	var c rune
	for c1 := range b.First {
		if b.Second[c1] {
			c = c1
			break
		}
	}
	if c == 0 {
		panic("no common")
	}
	return c
}

func FindCommonItemInSets(b1, b2, b3 map[rune]bool) rune {
	for r, _ := range b1 {
		if b2[r] && b3[r] {
			return r
		}
	}
	panic("nothing in common")
}

func SetOf(section string) map[rune]bool {
	m := map[rune]bool{}
	for _, r := range section {
		m[r] = true
	}
	return m
}

func ReadInput() []Backpack {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var backpacks []Backpack
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		b := Backpack{}
		b.Both = SetOf(line)
		b.First = SetOf(line[:len(line)/2])
		b.Second = SetOf(line[len(line)/2:])
		backpacks = append(backpacks, b)
	}
	return backpacks
}
