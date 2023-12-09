package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	games := ReadInput()
	part1(games)
	part2(games)
}

type Card struct {
	WinningNums map[int]bool
	ActualNums  map[int]bool
	ID          int
	Matching    int
	Copies      int
}

func part1(cards []Card) {
	points := 0
	for _, game := range cards {
		points += game.Points()
	}

	fmt.Println("Part 1:", points)
}

const maxWinningNumbers = 10

func part2(cards []Card) {
	for i := 0; i < len(cards); i++ {
		// Find num copies of this one by looking up
		from := i - maxWinningNumbers
		if from < 0 {
			from = 0
		}
		for j := i - 1; j >= from; j-- {
			// If that card had enough matches to affect this card,
			// add copies to this one.
			if cards[j].ID+cards[j].Matching >= cards[i].ID {
				cards[i].Copies += cards[j].Copies
			}
		}

		// Find matches for this one
		cards[i].Matching = len(Intersection(cards[i].WinningNums, cards[i].ActualNums))
	}

	numTotal := 0
	for _, card := range cards {
		numTotal += card.Copies
	}
	fmt.Println("Part 2:", numTotal)
}

func (g Card) Points() int {
	pointNums := Intersection(g.WinningNums, g.ActualNums)
	points := 0
	if len(pointNums) > 0 {
		points = 1
	}
	for i := 1; i < len(pointNums); i++ {
		points *= 2
	}
	return points
}

func Intersection(winningNums, actualNums map[int]bool) (scoreNums map[int]bool) {
	scoreNums = make(map[int]bool)
	for winningNum := range winningNums {
		if actualNums[winningNum] {
			scoreNums[winningNum] = true
		}
	}
	return scoreNums
}

func ReadInput() []Card {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cards := make([]Card, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

		// Card 1
		// 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		parts := strings.Split(scanner.Text(), ": ")

		// "  1"
		gameID, err := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], "Card ")[1]))
		check(err)

		// 41 48 83 86 17
		// 83 86  6 31 17  9 48 53
		numberParts := strings.Split(parts[1], " | ")

		winningNums := make(map[int]bool)
		for _, num := range strings.Split(strings.TrimSpace(numberParts[0]), " ") {
			if num == "" {
				continue
			}
			// 41
			winningNum, err := strconv.Atoi(num)
			check(err)
			winningNums[winningNum] = true
		}

		actualNums := make(map[int]bool)
		for _, num := range strings.Split(numberParts[1], " ") {
			if num == "" {
				continue
			}
			// 83
			actualNum, err := strconv.Atoi(strings.TrimSpace(num))
			check(err)
			actualNums[actualNum] = true
		}

		card := Card{
			WinningNums: winningNums,
			ActualNums:  actualNums,
			ID:          gameID,
			Copies:      1,
		}
		cards = append(cards, card)
	}

	return cards
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
