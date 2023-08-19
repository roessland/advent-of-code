package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"github.com/roessland/advent-of-code/2022/aocutil"
	"log"
	"slices"
	"strings"
)

//go:embed input.txt input-ex1.txt
var inputFiles embed.FS

func main() {
	part1()
	part2()
}

type MonkeyID int
type WorryLevel int
type OperationFunc func(worryLevel WorryLevel) WorryLevel
type TestFunc func(worryLevel WorryLevel) MonkeyID

type Monkey struct {
	Items          []WorryLevel
	Operation      OperationFunc
	Test           TestFunc
	NumInspections int
}

type State struct {
	Monkeys      []*Monkey
	ReliefFactor int
	Modulus      int
}

func part1() {
	state := readInput()
	state.ReliefFactor = 3
	for i := 0; i < 20; i++ {
		state.DoRound()
	}
	fmt.Println("Part 1 monkey business level: ", state.monkeyBusinessLevel())
}

func part2() {
	state := readInput()
	state.ReliefFactor = 1
	for i := 0; i < 10000; i++ {
		state.DoRound()
	}
	fmt.Println("Part 2 monkey business level: ", state.monkeyBusinessLevel())
}

func (s *State) DoRound() {
	for _, monkey := range s.Monkeys {
		s.DoTurn(monkey)
	}
}

func (s *State) DoTurn(monkey *Monkey) {
	for _, worryLevel := range monkey.Items {
		// Monkey inspects item, increasing your worry level
		monkey.NumInspections++
		newWorryLevel := monkey.Operation(worryLevel)

		// After inspection, you are relieved the item didn't break.
		newWorryLevel /= WorryLevel(s.ReliefFactor)

		// We can work with worryLevel mod LCM(divisors) since
		// each monkey only cares about the remainder of the worry level.
		newWorryLevel = newWorryLevel % WorryLevel(s.Modulus)

		// Monkey throws the item to another monkey based on the new worry level
		destinationMonkey := monkey.Test(newWorryLevel)
		s.Monkeys[destinationMonkey].Items = append(s.Monkeys[destinationMonkey].Items, newWorryLevel)
	}
	// Monkey threw away all its items
	monkey.Items = nil
}

func (s *State) monkeyBusinessLevel() int {
	var numInspections []int
	for _, monkey := range s.Monkeys {
		numInspections = append(numInspections, monkey.NumInspections)
	}
	slices.Sort(numInspections)
	numMonkeys := len(s.Monkeys)
	return numInspections[numMonkeys-1] * numInspections[numMonkeys-2]
}

func readInput() (state State) {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	state.Modulus = 1

	for {
		if doneScanning(scanner) {
			break
		}
		monkey := Monkey{}
		monkey.Items = readItems(scanner)
		monkey.Operation = readOperation(scanner)
		testFunc, divisor := readTest(scanner)
		monkey.Test = testFunc
		state.Modulus *= divisor // They are all primes, so this is fine.
		state.Monkeys = append(state.Monkeys, &monkey)
		readBlankLine(scanner)
	}
	return
}

func doneScanning(scanner *bufio.Scanner) bool {
	scanner.Scan()
	return len(scanner.Text()) == 0
}

func readItems(scanner *bufio.Scanner) []WorryLevel {
	scanner.Scan()
	return aocutil.GetNumsInString[WorryLevel](scanner.Text())
}

func readOperation(scanner *bufio.Scanner) OperationFunc {
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(strings.TrimSpace(line), " ")
	// 0          1   2 3   4 5
	// Operation: new = old * 19
	return makeOperationFunc(parts[3], parts[4], parts[5])
}

func makeOperationFunc(aStr, opStr, bStr string) OperationFunc {
	binOp := getBinaryOp(opStr)
	if aStr == "old" && bStr == "old" {
		return func(w WorryLevel) WorryLevel { return binOp(w, w) }
	} else if aStr == "old" {
		b := WorryLevel(aocutil.GetIntsInString(bStr)[0])
		return func(w WorryLevel) WorryLevel { return binOp(w, b) }
	} else if bStr == "old" {
		a := WorryLevel(aocutil.GetIntsInString(aStr)[0])
		return func(w WorryLevel) WorryLevel { return binOp(a, w) }
	} else {
		panic(fmt.Sprintf("Unknown operation: %s %s %s", aStr, opStr, bStr))
	}
}

func getBinaryOp(plusOrMult string) func(a, b WorryLevel) WorryLevel {
	if plusOrMult == "+" {
		return func(a, b WorryLevel) WorryLevel { return a + b }
	} else if plusOrMult == "*" {
		return func(a, b WorryLevel) WorryLevel { return a * b }
	} else {
		panic("Unknown operation: " + plusOrMult)
	}
}

func readTest(scanner *bufio.Scanner) (TestFunc, int) {
	scanner.Scan()
	lineTest := scanner.Text()
	divisor := aocutil.GetIntsInString(lineTest)[0]

	scanner.Scan()
	lineIf := scanner.Text()
	targetIf := MonkeyID(aocutil.GetIntsInString(lineIf)[0])

	scanner.Scan()
	lineElse := scanner.Text()
	targetElse := MonkeyID(aocutil.GetIntsInString(lineElse)[0])

	return makeTestFunc(divisor, targetIf, targetElse), divisor
}

func makeTestFunc(divisor int, targetIf, targetElse MonkeyID) TestFunc {
	return func(worryLevel WorryLevel) MonkeyID {
		if int(worryLevel)%divisor == 0 {
			return targetIf
		}
		return targetElse
	}
}

func readBlankLine(scanner *bufio.Scanner) {
	scanner.Scan()
}
