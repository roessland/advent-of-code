package main

import (
	"fmt"
	"math"
)

func Min(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func PrintDiffusion(dm [][]int) {
	for y := 0; y < len(dm); y++ {
		for x := 0; x < len(dm[0]); x++ {
			v := dm[y][x]
			if v < 10 {
				fmt.Printf("%d", v)
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func (m Map) Diffusion(sources []Pos) [][]int {
	Ny := len(m)
	Nx := len(m[0])
	dm := make([][]int, 0, Ny)
	for y := range m {
		dm = append(dm, make([]int, Nx))
		for x := 0; x < Nx; x++ {
			dm[y][x] = math.MaxInt32
		}
	}
	for _, p := range sources {
		dm[p.Y][p.X] = 0
	}

	for {
		updated := false
		for y := 1; y < Ny-1; y++ {
			for x := 1; x < Nx-1; x++ {
				if m[y][x].Unit != nil {
					continue
				}
				p := Pos{x, y}
				d := dm[y][x]
				for _, adj := range p.Adjacent() {
					if m[adj.Y][adj.X].Type != Open {
						continue
					}
					if m[adj.Y][adj.X].Unit != nil {
						continue
					}
					if d+1 < dm[adj.Y][adj.X] {
						dm[adj.Y][adj.X] = d + 1
						updated = true
					}
				}
			}
		}
		if !updated {
			return dm
		}
	}
}
