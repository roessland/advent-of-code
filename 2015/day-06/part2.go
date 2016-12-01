package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

type Grid struct {
	lights [1000][1000]int
}

type Area struct {
	x0, y0, x1, y1 int
}

func NewAreaFromStrings(x0y0, x1y1 string) Area {
	x0y0str := strings.Split(x0y0, ",")
	x0, _ := strconv.Atoi(x0y0str[0])
	y0, _ := strconv.Atoi(x0y0str[1])
	x1y1str := strings.Split(x1y1, ",")
	x1, _ := strconv.Atoi(x1y1str[0])
	y1, _ := strconv.Atoi(x1y1str[1])
	return Area{x0, y0, x1, y1}
}

func (g *Grid) toggle(area Area) {
	for y := area.y0; y <= area.y1; y++ {
		for x := area.x0; x <= area.x1; x++ {
			g.lights[y][x] += 2
		}

	}
}

func (g *Grid) turnon(area Area) {
	for y := area.y0; y <= area.y1; y++ {
		for x := area.x0; x <= area.x1; x++ {
			g.lights[y][x] += 1
		}
	}
}

func (g *Grid) turnoff(area Area) {
	for y := area.y0; y <= area.y1; y++ {
		for x := area.x0; x <= area.x1; x++ {
			if g.lights[y][x] > 0 {
				g.lights[y][x] -= 1
			}
		}

	}
}

func (g *Grid) count() int {
	count := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			count += g.lights[y][x]
		}
	}
	return count
}

func main() {
	fmt.Printf("haha\n")
	g := Grid{}
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		words := strings.Split(line, " ")
		if words[0] == "turn" && words[1] == "off" {
			area := NewAreaFromStrings(words[2], words[4])
			g.turnoff(area)
		}
		if words[0] == "turn" && words[1] == "on" {
			area := NewAreaFromStrings(words[2], words[4])
			g.turnon(area)
		}
		if words[0] == "toggle" {
			area := NewAreaFromStrings(words[1], words[3])
			g.toggle(area)
		}
	}
	fmt.Printf("%v\n", g.count())

}
