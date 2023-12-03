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

type Game struct {
	Sets []map[string]int
	ID   int
}

func part1(games []Game) {
	sum := 0
	for _, game := range games {
		if game.IsPossible() {
			sum += game.ID
		}
	}

	fmt.Println(sum)
}

func (g Game) IsPossible() bool {
	numColors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, subset := range g.Sets {
		for color, num := range subset {
			if num > numColors[color] {
				return false
			}
		}
	}
	return true
}

func part2(games []Game) {
	sum := 0
	for _, game := range games {
		set := Max(game.Sets)
		sum += Power(set)
	}
	fmt.Println(sum)
}

func Max(subsets []map[string]int) map[string]int {
	max := make(map[string]int)
	for _, subset := range subsets {
		for color, num := range subset {
			if num > max[color] {
				max[color] = num
			}
		}
	}
	return max
}

func Power(subset map[string]int) int {
	prod := 1
	for _, num := range subset {
		prod *= num
	}
	return prod
}

func ReadInput() []Game {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	games := make([]Game, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		// Game 1
		// 4 blue, 16 green, 2 red; 5 red, 11 blue, 16 green
		gameSubsets := strings.Split(scanner.Text(), ": ")

		// 1
		gameID, err := strconv.Atoi(strings.Split(gameSubsets[0], " ")[1])
		check(err)

		// 4 blue, 16 green, 2 red
		// 5 red, 11 blue, 16 green
		subsetsStrs := strings.Split(gameSubsets[1], "; ")

		subsets := make([]map[string]int, 0)
		for _, subsetStr := range subsetsStrs {
			// map[red:5 blue:11 green:16]
			subset := make(map[string]int)
			cubes := strings.Split(subsetStr, ", ")
			for _, cube := range cubes {
				numColor := strings.Split(cube, " ")
				num, err := strconv.Atoi(numColor[0])
				check(err)
				color := numColor[1]
				subset[color] = num
			}
			subsets = append(subsets, subset)
		}

		game := Game{
			ID:   gameID,
			Sets: subsets,
		}
		games = append(games, game)
	}

	return games
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
