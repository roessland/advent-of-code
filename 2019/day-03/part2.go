package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Pos struct {
	Left, Up int
}

type Cable struct {
	Visits map[Pos]int
	HeadPos Pos
	DistanceTo map[Pos]int
	DistanceToHead int
}

type Cables []Cable

type Section struct {
	Dir Pos
	Length int
}

func (cables Cables) FindDifferentCableIntersections() []Pos {
	var intersections []Pos
	for pos, _ := range cables[0].Visits {
		fstVisits := cables[0].Visits[pos]
		sndVisits := cables[1].Visits[pos]
		if fstVisits >= 1 && sndVisits >= 1 {
			intersections = append(intersections, pos)
		}
	}
	return intersections
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (pos Pos) ManhattanDistance() int {
	return Abs(pos.Up) + Abs(pos.Left)
}

func Closest(positions []Pos) Pos {
	closest := Pos{99999, 99999}
	closestDist := closest.ManhattanDistance()
	origin := Pos{}
	for _, pos := range positions {
		if pos == origin {
			// Doesn't count
			continue
		}
		dist := pos.ManhattanDistance()
		if dist < closestDist {
			closest = pos
			closestDist = dist
		}
	}
	return closest
}

func (cables Cables) SignalDelay(pos Pos) int {
	return cables[0].DistanceTo[pos] + cables[1].DistanceTo[pos]
}

func (cables Cables) SmallestDelay(intersections []Pos) int {
	smallestDelay := 999999999
	for _, intersection := range intersections {
		signalDelay := cables.SignalDelay(intersection)
		if signalDelay < smallestDelay {
			smallestDelay = signalDelay
		}
	}
	return smallestDelay
}

func (a Pos) Add(b Pos) Pos {
	return Pos{a.Left + b.Left, a.Up + b.Up}
}

func LoopSegment(p0 Pos, length int, dir Pos, f func(Pos)) Pos {
	p := p0
	for i := 0; i < length; i++ {
		p = p.Add(dir)
		f(p)
	}
	return p
}

func AddToCable(cable *Cable, section Section) {
	cable.HeadPos = LoopSegment(cable.HeadPos, section.Length, section.Dir, func(p Pos) {
		cable.Visits[p]++
		cable.DistanceToHead++
		if _, alreadyExists := cable.DistanceTo[p]; !alreadyExists {
			cable.DistanceTo[p] = cable.DistanceToHead
		}
	})
}

func GetDir(c rune) Pos {
	switch c {
	case 'L': return Pos{Left: 1, Up: 0}
	case 'R': return Pos{Left: -1, Up: 0}
	case 'U': return Pos{Left: 0, Up: 1}
	case 'D': return Pos{Left: 0, Up: -1}
	default: log.Fatal("Unknown direction", c)
	}
	return Pos{}
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(buf), "\n")
	cables := make(Cables, 2)
	cables[0] = Cable{Visits: make(map[Pos]int), HeadPos: Pos{0, 0}, DistanceTo: make(map[Pos]int), DistanceToHead: 0}
	cables[1] = Cable{Visits: make(map[Pos]int), HeadPos: Pos{0, 0}, DistanceTo: make(map[Pos]int), DistanceToHead: 0}
	for cableNum, line := range lines {
		if line == "" {
			continue
		}
		words := strings.Split(line, ",")
		for _, word := range words {
			dir := GetDir(rune(word[0]))
			lengthStr := word[1:]
			length, err := strconv.Atoi(lengthStr)
			if err != nil {
				log.Fatal(err)
			}
			AddToCable(&cables[cableNum], Section{dir, length})
		}
	}

	fmt.Println(Closest(cables.FindDifferentCableIntersections()).ManhattanDistance())
	fmt.Println(cables.SmallestDelay(cables.FindDifferentCableIntersections()))
}
