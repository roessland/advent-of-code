package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Chemical string

type Quantity struct {
	Chemical Chemical
	Amount   int
}

func NewQuantity(str string) Quantity {
	numName := strings.Split(str, " ")
	n, err := strconv.Atoi(numName[0])
	if err != nil {
		log.Fatal(err)
	}
	return Quantity{
		Chemical: Chemical(numName[1]),
		Amount:   n,
	}
}

type Reaction struct {
	Name   string
	Inputs []Quantity
	Output Quantity
}

type Reactions map[Chemical]*Reaction

func NewReactions() Reactions {
	return make(Reactions)
}

func (rs Reactions) AddFromString(line string) {
	reaction := Reaction{Name: line}
	insOut := strings.Split(line, " => ")
	for _, inStr := range strings.Split(insOut[0], ", ") {
		reaction.Inputs = append(reaction.Inputs, NewQuantity(inStr))
	}
	reaction.Output = NewQuantity(insOut[1])
	rs[reaction.Output.Chemical] = &reaction
}

func (rs Reactions) Print() {
	for outChemical, reaction := range rs {
		fmt.Println(outChemical, reaction.Inputs, reaction.Output)
	}
}

func (rs Reactions) Fix(heap map[Chemical]int) bool {
	fixed := false
	for chemical, amount := range heap {
		if amount >= 0 || chemical == "ORE" {
			continue
		}
		r := rs[chemical]
		mult := -amount / r.Output.Amount
		for mult*r.Output.Amount+amount < 0 {
			mult++
		}
		for _, input := range r.Inputs {
			heap[input.Chemical] -= mult * input.Amount
		}
		heap[r.Output.Chemical] += mult * r.Output.Amount
		fixed = true
	}
	return fixed
}

func (rs Reactions) OreRequired(heap map[Chemical]int) int {
	for rs.Fix(heap) {
	}
	return -heap["ORE"]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	reactions := NewReactions()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		reactions.AddFromString(scanner.Text())
	}

	maxOrePerFuel := reactions.OreRequired(map[Chemical]int{
		"FUEL": -1,
	})
	fmt.Println("Part 1:", maxOrePerFuel)

	// Bisection search
	minOrePerFuel := maxOrePerFuel / 100
	trillion := 1000000000000
	loFuse := trillion / maxOrePerFuel
	hiFuse := trillion / minOrePerFuel
	for loFuse < hiFuse-1 {
		midFuse := (loFuse + hiFuse) / 2
		midOre := reactions.OreRequired(map[Chemical]int{
			"FUEL": -midFuse,
		})
		if midOre < trillion {
			loFuse = midFuse
		} else {
			hiFuse = midFuse
		}
	}
	fmt.Println("Part 2:", loFuse)
}
