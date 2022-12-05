package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt input-ex.txt
var inputFiles embed.FS

func main() {
	rounds := ReadInput()
	part1(rounds)
	part2(rounds)
}

func ReadInput() []Round {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var rounds []Round

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		p1, p2 := parts[0], parts[1]
		round := Round{
			Opponent: getMove1(p1),
			Myself:   getMove1(p2),
		}
		switch p2 {
		case "X":
			round.Target = -1
		case "Y":
			round.Target = 0
		case "Z":
			round.Target = 1
		}
		rounds = append(rounds, round)
	}
	return rounds
}

type Move string

const Rock Move = "rock"
const Paper Move = "paper"
const Scissors Move = "scissors"

type Round struct {
	Opponent Move
	Myself   Move
	Target   int
}

func part1(rounds []Round) {
	totalScore := 0
	for _, round := range rounds {
		totalScore += getScore1(round)
	}
	fmt.Println(totalScore)
}

func getScore1(round Round) int {
	score := 0

	switch round.Myself {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}

	switch getOutcome1(round) {
	case -1:
		score += 0
	case 0:
		score += 3
	case 1:
		score += 6
	}

	return score
}

func getMove1(coded string) Move {
	switch coded {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	panic("nope")
}

func getOutcome1(round Round) int {
	if round.Opponent == round.Myself {
		return 0
	}

	if round.Opponent == Rock && round.Myself == Scissors {
		return -1
	}
	if round.Opponent == Scissors && round.Myself == Rock {
		return 1
	}

	if round.Opponent == Scissors && round.Myself == Paper {
		return -1
	}
	if round.Opponent == Paper && round.Myself == Scissors {
		return 1
	}

	if round.Opponent == Paper && round.Myself == Rock {
		return -1
	}
	if round.Opponent == Rock && round.Myself == Paper {
		return 1
	}

	panic("nope")
}

func part2(rounds []Round) {
	totalScore := 0
	for _, round := range rounds {
		totalScore += getScore2(round)
	}
	fmt.Println(totalScore)
}

func getScore2(round Round) int {
	score := 0

	switch getMove2(round) {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}

	switch getOutcome2(round) {
	case -1:
		score += 0
	case 0:
		score += 3
	case 1:
		score += 6
	}

	return score
}

func getMove2(round Round) Move {
	if round.Target == 0 {
		return round.Opponent
	} else if round.Target == 1 {
		switch round.Opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	} else if round.Target == -1 {
		switch round.Opponent {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}
	panic("nope")
}

func getOutcome2(round Round) int {
	return round.Target
}
