package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Particle struct {
	Px, Py, Vx, Vy int
}

func (p *Particle) Update(dir int) {
	p.Px += p.Vx * dir
	p.Py += p.Vy * dir
}

func Extent(ps []Particle) int {
	minX, minY, maxX, maxY := 99999, 99999, -99999, -99999
	for _, p := range ps {
		if p.Px < minX {
			minX = p.Px
		}
		if p.Py < minY {
			minY = p.Py
		}
		if p.Px > maxX {
			maxX = p.Px
		}
		if p.Py > maxY {
			maxY = p.Py
		}
	}
	return maxX - minX + maxY - minY
}

func Print(ps []Particle) {
	minX := 129
	maxX := 199
	minY := 120
	maxY := 138
	//minX := -20
	//maxX := 20
	//minY := -20
	//maxY := 20
	arr := make([][]rune, maxY-minY+1)
	for iy := range arr {
		arr[iy] = make([]rune, maxX-minX+1)
		for ix := range arr[iy] {
			arr[iy][ix] = '.'
		}
	}
	for _, p := range ps {
		arr[p.Py-minY][p.Px-minX] = '▓'
	}
	for _, line := range arr {
		fmt.Println(string(line))
	}

}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	particles := []Particle{}
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		line = strings.Replace(line, " ", "", -1)
		var a, b, c, d int
		fmt.Sscanf(line, "position=<%d,%d>velocity=<%d,%d>", &a, &b, &c, &d)
		particles = append(particles, Particle{a, b, c, d})
	}
	fmt.Println(particles)
	lastExtent := 9999999
	time := 0
	for {
		for i := range particles {
			particles[i].Update(1)
		}
		extent := Extent(particles)
		if extent > lastExtent {
			for i := range particles {
				particles[i].Update(-1)
			}
			fmt.Println(particles)
			Print(particles)
			fmt.Println(time)
			break
		}
		lastExtent = extent
		time++
	}
}
