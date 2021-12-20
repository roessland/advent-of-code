package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Pos struct {
	X, Y int
}

type Im map[Pos]byte

func CountLit(im Im) int {
	num := 0
	for _, lit := range im {
		if lit == 1 {
			num++
		}
	}
	return num
}

func GetNext(values []byte, algorithm []byte) byte {
	var index uint16
	for _, value := range values {
		index = index*2 + uint16(value)
	}
	return algorithm[index]
}

func EnhanceTwice(im0 Im, algo []byte) Im {
	// Can spread out a distance of 2
	targetArea := make(map[Pos]struct{})
	for pos, _ := range im0 {
		x, y := pos.X, pos.Y
		for x0 := x - 2; x0 <= x+2; x0++ {
			for y0 := y - 2; y0 <= y+2; y0++ {
				targetArea[Pos{x0, y0}] = struct{}{}
			}
		}
	}

	// Find all values for next step
	cache := map[Pos]byte{}
	for pos, _ := range targetArea {
		x1, y1 := pos.X, pos.Y
		var values0 = make([]byte, 0, 9)
		for y0 := y1 - 1; y0 <= y1+1; y0++ {
			for x0 := x1 - 1; x0 <= x1+1; x0++ {
				values0 = append(values0, im0[Pos{x0, y0}])
			}
		}
		cache[pos] = GetNext(values0, algo)
	}

	// Find all values for next, next step
	im := make(Im)
	for pos, _ := range targetArea {
		x2, y2 := pos.X, pos.Y
		var values1 = make([]byte, 0, 9)
		for y1 := y2 - 1; y1 <= y2+1; y1++ {
			for x1 := x2 - 1; x1 <= x2+1; x1++ {
				cached, ok := cache[Pos{x1, y1}]
				if !ok {
					cached = 1
				}
				values1 = append(values1, cached)
			}
		}

		if GetNext(values1, algo) == 1 {
			im[pos] = 1
		}
	}
	return im
}

func main() {
	t0 := time.Now()
	im, algo := ReadInput()
	im = EnhanceTwice(im, algo)
	fmt.Println(CountLit(im))
	for i := 2; i < 50; i += 2 {
		im = EnhanceTwice(im, algo)
	}
	fmt.Println(CountLit(im))
	fmt.Println(time.Since(t0))
}

func ReadInput() (Im, []byte) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	var algo []byte
	scanner.Scan()
	for _, c := range scanner.Text() {
		if c == '#' {
			algo = append(algo, 1)
		} else if c == '.' {
			algo = append(algo, 0)
		}
	}
	scanner.Scan()

	var im = make(map[Pos]byte)
	y := 0
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			if c == '#' {
				im[Pos{x, y}] = 1
			}
		}
		y++
	}
	return im, algo
}
