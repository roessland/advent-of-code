package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"github.com/roessland/advent-of-code/2022/aocutil"
	"github.com/roessland/gopkg/sliceutil"
	"github.com/roessland/gopkg/stackqueue"
	"log"
	"unicode"
)

//go:embed input.txt input-ex.txt
var inputFiles embed.FS

func main() {
	part1()
	part2()
}

func part1() {
	stacks, moves := ReadInput()
	Print(stacks)
	for _, move := range moves {
		for move.No > 0 {
			crate := stacks[move.From].Pop()
			stacks[move.To].Push(crate)
			move.No--
		}
	}
	msg := ""
	for _, stack := range stacks {
		msg = fmt.Sprintf("%s%c", msg, stack.Peek())
	}
	fmt.Println("Part 1:", msg)
	Print(stacks)
}

func part2() {
	stacks, moves := ReadInput()
	Print(stacks)
	for _, move := range moves {
		remaining := move.No
		for remaining > 0 {
			crate := stacks[move.From].Pop()
			stacks[move.To].Push(crate)
			remaining--
		}
		sliceutil.Reverse(stacks[move.To][len(stacks[move.To])-move.No:])
	}
	msg := ""
	for _, stack := range stacks {
		msg = fmt.Sprintf("%s%c", msg, stack.Peek())
	}
	fmt.Println("Part 1:", msg)
	Print(stacks)
}

type Move struct {
	No, From, To int
}

func Print(stacks []stackqueue.Stack[byte]) {
	for _, s := range stacks {
		fmt.Println(string(s))
	}
}

func ReadInput() (stacks []stackqueue.Stack[byte], moves []Move) {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		N := (len(line) + 1) / 4
		if stacks == nil {
			stacks = make([]stackqueue.Stack[byte], N)
		}
		if line == "" {
			break
		}
		if unicode.IsNumber(rune(line[1])) {
			continue
		}

		for i := 0; i < N; i++ {
			c := line[i*4+1]
			if unicode.IsLetter(rune(c)) {
				stacks[i] = append(stacks[i], c)
			}
		}
	}

	for i := range stacks {
		sliceutil.Reverse(stacks[i])
	}

	for scanner.Scan() {
		line := scanner.Text()
		p := aocutil.GetIntsInString(line)
		moves = append(moves, Move{p[0], p[1] - 1, p[2] - 1})
	}
	return stacks, moves
}
