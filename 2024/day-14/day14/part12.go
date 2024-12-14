package day14

import (
	"embed"
	"fmt"
	"time"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/mathutil"
)

//go:embed input*.txt
var Input embed.FS

func ChristmasFactor(pos map[[2]int]bool, Nx, Ny int) int {
	left := 0
	right := 0
	for i := 0; i < Ny; i++ {
		unique := map[[2]int]bool{}
		for j := 0; j < Nx; j++ {
			if pos[[2]int{i, j}] {
				unique[[2]int{i, j}] = true
				if j < Nx/2 {
					left++
				} else {
					right++
				}
			}
		}
	}
	return -mathutil.AbsInt(left - right)
}

func Part12(inputName string, Nx, Ny int) (int, int) {
	input := aocutil.FSGetIntsInStringLines(Input, inputName)

	// t := 100

	q := map[int]int{}

	for t := 0; t < 10520; t++ {
		pos := map[[2]int]bool{}

		for _, row := range input {

			px0, py0, vx0, vy0 := row[0], row[1], row[2], row[3]
			vx0 = (vx0 + Nx) % Nx
			vy0 = (vy0 + Ny) % Ny

			px := (px0 + t*vx0) % Nx
			py := (py0 + t*vy0) % Ny

			pos[[2]int{py, px}] = true

			if t == 100 {
				if px < Nx/2 && py < Ny/2 {
					q[0]++
				}
				if px > Nx/2 && py < Ny/2 {
					q[1]++
				}
				if px < Nx/2 && py > Ny/2 {
					q[2]++
				}
				if px > Nx/2 && py > Ny/2 {
					q[3]++
				}
			}
		}

		cf := ChristmasFactor(pos, Nx, Ny)
		if cf > -352 {
			continue
		}

		if cf < -362 {
			continue
		}

		fmt.Println()
		for i := 0; i < 60; i++ {
			for j := 0; j < 103; j++ {
				if pos[[2]int{i, j}] {
					print("█")
				} else {
					print(".")
				}
			}
			println()
		}
		fmt.Println(t)
		time.Sleep(100 * time.Millisecond * time.Duration(cf))
		fmt.Println()
	}

	sum1 := q[0] * q[1] * q[2] * q[3]
	return sum1, 0
}
