package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/advent-of-code/2018/aocutil"
	"log"
	"os"
)

const (
	A = 1
	B = 2
	C = 3
)

type State [4]int
type Instr [4]int

type OpCode int

func Addr(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] + s[instr[B]]
	return s
}

func Addi(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] + instr[B]
	return s
}

func Mulr(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] * s[instr[B]]
	return s
}

func Muli(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] * instr[B]
	return s
}

func Banr(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] & s[instr[B]]
	return s
}

func Bani(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] & instr[B]
	return s
}

func Borr(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] | s[instr[B]]
	return s
}

func Bori(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]] | instr[B]
	return s
}

func Setr(s State, instr Instr) State {
	s[instr[C]] = s[instr[A]]
	return s
}

func Seti(s State, instr Instr) State {
	s[instr[C]] = instr[A]
	return s
}

func Gtir(s State, instr Instr) State {
	if instr[A] > s[instr[B]] {
		s[instr[C]] = 1
	} else {
		s[instr[C]] = 0
	}
	return s
}

func Gtrr(s State, instr Instr) State {
	if s[instr[A]] > s[instr[B]] {
		s[instr[C]] = 1
	} else {
		s[instr[C]] = 0
	}
	return s
}

func Gtri(s State, instr Instr) State {
	if s[instr[A]] > instr[B] {
		s[instr[C]] = 1
	} else {
		s[instr[C]] = 0
	}
	return s
}

func Eqir(s State, instr Instr) State {
	if instr[A] == s[instr[B]] {
		s[instr[C]] = 1
	} else {
		s[instr[C]] = 0
	}
	return s
}

func Eqrr(s State, instr Instr) State {
	if s[instr[A]] == s[instr[B]] {
		s[instr[C]] = 1
	} else {
		s[instr[C]] = 0
	}
	return s
}

func Eqri(s State, instr Instr) State {
	if s[instr[A]] == instr[B] {
		s[instr[C]] = 1
	} else {
		s[instr[C]] = 0
	}
	return s
}

type Op func(State, Instr) State

var OpsByName = map[string]Op{
	"addr": Addr,
	"addi": Addi,
	"mulr": Mulr,
	"muli": Muli,
	"banr": Banr,
	"bani": Bani,
	"borr": Borr,
	"bori": Bori,
	"setr": Setr,
	"seti": Seti,
	"gtir": Gtir,
	"gtrr": Gtrr,
	"gtri": Gtri,
	"eqir": Eqir,
	"eqrr": Eqrr,
	"eqri": Eqri,
}

func ReadInput() ([]Sample, []Instr) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var samples []Sample
	scanner := bufio.NewScanner(f)
	for {
		scanner.Scan()
		regs0Line := scanner.Text()

		scanner.Scan()
		opsLine := scanner.Text()

		if regs0Line == "" && opsLine == "" {
			break
		}

		scanner.Scan()
		regs1Line := scanner.Text()

		scanner.Scan()

		regNums0 := aocutil.GetIntsInString(regs0Line)
		instrNums := aocutil.GetIntsInString(opsLine)
		regNums1 := aocutil.GetIntsInString(regs1Line)

		state0 := *(*[4]int)(regNums0)
		instr := *(*[4]int)(instrNums)
		state1 := *(*[4]int)(regNums1)

		samples = append(samples, Sample{state0, instr, state1})

	}
	var testProgram []Instr
	for scanner.Scan() {
		is := aocutil.GetIntsInString(scanner.Text())
		testProgram = append(testProgram, Instr{is[0], is[1], is[2], is[3]})
	}
	return samples, testProgram
}

type Sample struct {
	State0 State
	Instr  Instr
	State1 State
}

// 516 too low

func Intersection(A, B map[string]struct{}) (S map[string]struct{}) {
	S = make(map[string]struct{})
	for a := range A {
		if _, ok := B[a]; ok {
			S[a] = struct{}{}
		}
	}
	return S
}

func GetAny(A map[string]struct{}) string {
	for s := range A {
		return s
	}
	panic("empty set")
}

func main() {
	samples, testProgram := ReadInput()

	// Part 1
	var behavesLikeThreeOrMoreCount int
	for _, sample := range samples {
		behavesLikeCount := 0
		for _, op := range OpsByName {
			if op(sample.State0, sample.Instr) == sample.State1 {
				behavesLikeCount++
			}
		}
		if behavesLikeCount >= 3 {
			behavesLikeThreeOrMoreCount++
		}
	}
	fmt.Println("Part 1:", behavesLikeThreeOrMoreCount)

	// Build a set of all ops for all opcodes.
	possibleOpsSet := map[OpCode]map[string]struct{}{}
	for _, sample := range samples {
		opCode := OpCode(sample.Instr[0])
		possibleOpsSet[opCode] = map[string]struct{}{}
		for opName := range OpsByName {
			possibleOpsSet[opCode][opName] = struct{}{}
		}
	}

	// Narrow down the sets of opcodes using the samples.
	for _, sample := range samples {
		opCode := OpCode(sample.Instr[0])
		couldBeOps := map[string]struct{}{}
		for opName, op := range OpsByName {
			if op(sample.State0, sample.Instr) == sample.State1 {
				couldBeOps[opName] = struct{}{}
			}
		}
		possibleOpsSet[opCode] = Intersection(possibleOpsSet[opCode], couldBeOps)
	}

	// Find sets of size 1. Consider those as known, and delete them
	// from other sets. Repeat until all ops have been found.
	knownOps := map[OpCode]string{}
	for len(knownOps) < len(possibleOpsSet) {
		for opCode := range possibleOpsSet {
			if len(possibleOpsSet[opCode]) == 1 {
				opName := GetAny(possibleOpsSet[opCode])
				knownOps[opCode] = opName
				for unknownOpCode := range possibleOpsSet {
					delete(possibleOpsSet[unknownOpCode], opName)
				}
			}
		}
	}

	// Run the test program
	state := State{}
	for _, instr := range testProgram {
		opCode := OpCode(instr[0])
		opName := knownOps[opCode]
		op := OpsByName[opName]
		state = op(state, instr)
	}
	fmt.Println("Part 2:", state[0])
}
