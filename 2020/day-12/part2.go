package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Vec struct {
	E, N int
}

func (p Vec) Rot(degrees int) Vec {
	degrees = degrees % 360
	for degrees < 0 {
		degrees += 360
	}
	switch degrees {
	case 0:
		return p
	case 90:
		return Vec{E: -p.N, N: p.E}
	case 180:
		return Vec{E: -p.E, N: -p.N}
	case 270:
		return Vec{E: p.N, N: -p.E}
	default:
		log.Fatalf("rotating degrees %d", degrees)
		return Vec{}
	}
}

func (p Vec) Add(q Vec) Vec {
	return Vec{p.E + q.E, p.N + q.N}
}

func (p Vec) Sub(q Vec) Vec {
	return Vec{p.E - q.E, p.N - q.N}
}

func (p Vec) Mul(n int) Vec {
	return Vec{n*p.E, n*p.N}
}

type Ship struct {
	Pos Vec
	WaypointPos Vec
}

func (s Ship) Do(cmd uint8, n int) Ship {
	switch cmd {
	case 'N':
		s.WaypointPos.N += n
	case 'S':
		s.WaypointPos.N -= n
	case 'E':
		s.WaypointPos.E += n
	case 'W':
		s.WaypointPos.E -= n
	case 'L':
		delta := s.WaypointPos.Sub(s.Pos)
		s.WaypointPos = s.Pos.Add(delta.Rot(n))
	case 'R':
		delta := s.WaypointPos.Sub(s.Pos)
		s.WaypointPos = s.Pos.Add(delta.Rot(-n))
	case 'F':
		delta := s.WaypointPos.Sub(s.Pos)
		s.Pos = s.Pos.Add(delta.Mul(n))
		s.WaypointPos = s.Pos.Add(delta)
	default:
		log.Fatalf("unknown command %c", cmd)
	}
	return s
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (p Vec) Dist() int {
	return Abs(p.E) + Abs(p.N)
}

func main () {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ship := Ship{
		WaypointPos: Vec{E: 10, N: 1},
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := line[0]
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		ship = ship.Do(cmd, n)
	}
	fmt.Println(ship.Pos.Dist())
}