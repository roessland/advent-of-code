package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func GCD(a, b int) int {
	if a < 0 {
		return GCD(-a, b)
	}
	if b < 0 {
		return GCD(a, -b)
	}
	if b > a {
		return GCD(b, a)
	}
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

type Rat struct {
	P, Q int
}

func (r Rat) Simplify() Rat {
	d := GCD(r.P, r.Q)
	return Rat{r.P/d, r.Q/d}
}

func NewRat(P, Q int) Rat {
	return Rat{P, Q}.Simplify()
}

type Asteroid struct {
	I, J int
}

func ReadInput(filename string) []Asteroid {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var asteroids []Asteroid
	i := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if c == '#' {
				asteroids = append(asteroids, Asteroid{i, j})
			}
		}
		i++
	}
	return asteroids
}

func Part1(as []Asteroid) Asteroid {
	max := 0
	var best Asteroid
	for _, a := range as {
		visible := map[Rat]bool{}
		for _, b := range as {
			if a == b {
				continue
			}
			visible[NewRat(a.I-b.I, a.J-b.J)] = true
		}
		numVisible := len(visible)
		if numVisible > max {
			max = numVisible
			best = a
		}
	}
	fmt.Println("Part 1:", max)
	return best
}

type Angle struct {
	Rat
	Asteroids []Asteroid
}

func SortByDistanceFrom(as []Asteroid, s Asteroid) {
	sort.Slice(as, func(i, j int)bool{
		dI := (as[i].I-s.I)* (as[i].I-s.I) + (as[i].J-s.J)* (as[i].J-s.J)
		dJ := (as[j].I-s.I)* (as[j].I-s.I) + (as[j].J-s.J)* (as[j].J-s.J)
		return dI > dJ
	})
}

func Part2(as []Asteroid, station Asteroid) {

	// Group by direction from station
	asteroidsByDir := map[Rat][]Asteroid{}
	for _, a := range as {
		if a == station {
			continue
		}
		dir := NewRat(a.I-station.I, a.J-station.J)
		if asteroidsByDir[dir] == nil {
			asteroidsByDir[dir] = []Asteroid{}
		}
		asteroidsByDir[dir] = append(asteroidsByDir[dir], a)
	}

	// Convert to list, sort in radial distance
	var angles []Angle
	for dir, bs := range asteroidsByDir {
		SortByDistanceFrom(bs, station) // Closest is placed at end of list, for easy pop.
		angles = append(angles, Angle{dir, bs})
	}

	// Sort angles by shooting order
	sort.Slice(angles, func(i,j int)bool{
		a := angles[i]
		b := angles[j]
		// Small adjustment to make it start at 12 o clock
		return math.Atan2(float64(-a.Q)-0.000001, float64(a.P)) < math.Atan2(float64(-b.Q)-0.000001, float64(b.P))
	})

	// Loop around angles, shooting.
	shotsFired := 0
	Free:
	for {
		shotsFiredThisRound := 0
		for idx, angle := range angles {
			if len(angle.Asteroids) == 0 {
				continue
			}
			shotAsteroid := angle.Asteroids[len(angle.Asteroids)-1]
			angles[idx].Asteroids = angle.Asteroids[:len(angle.Asteroids)-1]
			shotsFiredThisRound++
			shotsFired++
			if shotsFired == 200 {
				fmt.Println("Part 2:", 100*shotAsteroid.J + shotAsteroid.I)
				break Free
			}
		}
		if shotsFiredThisRound == 0 {
			break
		}
	}
}

func main() {
	asteroids := ReadInput("input.txt")
	station := Part1(asteroids)
	Part2(asteroids, station)
}