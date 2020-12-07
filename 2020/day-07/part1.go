package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Color string
	Bags map[*Bag]int
}

func getBag(bags map[string]*Bag, color string) *Bag {
	if bags[color] == nil {
		bags[color] = &Bag{Color: color, Bags: map[*Bag]int{}}
	}
	return bags[color]
}

// if a bag is color, or contains color
func eventuallyContains(bag *Bag, color string, path string) (bool, string) {
	if path == "" {
		path = bag.Color
	}
	if bag.Color == color {
		return true, path
	}
	for inner := range bag.Bags {
		res, resPath := eventuallyContains(inner, color, path + "->" + inner.Color)
		if res {
			return true, resPath
		}
	}
	return false, ""
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	bags := map[string]*Bag{}
	scanner := bufio.NewScanner(f)
	// dotted aqua contain 1 mirrored green, 5 shiny maroon
	// dark silver contain no other
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " bags", "")
		line = strings.ReplaceAll(line, " bag", "")
		line = strings.Trim(line, ".")
		parts := strings.Split(line, " contain ")
		outerColor := parts[0]
		outerBag := getBag(bags, outerColor)
		innerBagsStr := parts[1]
		if innerBagsStr == "no other" {
			continue
		}
		nInnerBags := strings.Split(innerBagsStr, ", ")
		for _, nInnerBag := range nInnerBags {
			nColor := strings.SplitN(nInnerBag, " ", 2)
			n, err := strconv.Atoi(nColor[0])
			if err != nil {
				log.Fatal(err)
			}
			innerColor := nColor[1]
			innerBag := getBag(bags, innerColor)
			outerBag.Bags[innerBag] = n
		}
	}

	sum := -1 // subtract one for the shiny gold bag, which doesn't contain itself.
	for color := range bags {
		contains, _ := eventuallyContains(getBag(bags, color), "shiny gold", "")
		if contains {
			sum++
		}
	}
	fmt.Println(sum)
}
