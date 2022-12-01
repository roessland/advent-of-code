package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/advent-of-code/2018/aocutil"
	"github.com/roessland/gopkg/mathutil"
	"log"
	"math/rand"
	"os"
)

const N = 2000

type TileType rune

const (
	None         TileType = 0
	Sand         TileType = '.'
	Clay         TileType = '#'
	Spring       TileType = '+'
	StableWater  TileType = '~'
	FallingWater TileType = '|'
)

type Tile struct {
	Type             TileType
	VelX             int
	ReachedWallLeft  bool
	ReachedWallRight bool
}

type State [N][N]Tile

func NewState() State {
	s := State{}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			s[y][x] = Tile{Type: Sand}
		}
	}
	return s
}

func ReadInput() (state0 State, xmin, xmax, ymin, ymax int) {
	xmin, xmax, ymin, ymax = N, 0, 0, 0
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	state0 = NewState()
	state0[0][500].Type = Spring
	for scanner.Scan() {
		line := scanner.Text()
		nums := aocutil.GetIntsInString(line)
		// x=504, y=10..13
		// y=13, x=498..504
		fmt.Println(nums)
		switch line[0] {
		case 'x':
			x := nums[0]
			xmin, xmax = mathutil.MinInt(x, xmin), mathutil.MaxInt(x, xmax)
			ymax = mathutil.MaxInt(nums[2], ymax)
			for y := nums[1]; y <= nums[2]; y++ {
				state0[y][x].Type = Clay
			}
		case 'y':
			y := nums[0]
			ymax = mathutil.MaxInt(y, ymax)
			xmin, xmax = mathutil.MinInt(nums[1], xmin), mathutil.MaxInt(nums[2], xmax)
			for x := nums[1]; x <= nums[2]; x++ {
				state0[y][x].Type = Clay
			}
		}
	}
	return
}

func NextTile(prev, up, down, left, right, downleft, downright Tile) Tile {
	// Basic tiles
	next := Tile{}
	if prev.Type == Clay {
		next.Type = Clay
	} else if prev.Type == StableWater {
		next.Type = StableWater
	} else if up.Type == Spring {
		next.Type = FallingWater
	} else if up.Type == FallingWater && prev.Type == Sand {
		next.Type = FallingWater
	} else if up.Type == FallingWater && prev.Type == FallingWater {
		next = prev
		next.VelX = 0
	}

	// Falling water staying put and transformation to stable water
	if prev.Type == FallingWater && (down.Type == StableWater || down.Type == Clay) && prev.VelX == 0 {
		next = prev
		if left.Type == Clay || left.Type == StableWater {
			next.ReachedWallLeft = true
		}
		if right.Type == Clay || right.Type == StableWater {
			next.ReachedWallRight = true
		}
		if left.Type == FallingWater && left.VelX == 0 && left.ReachedWallLeft {
			next.ReachedWallLeft = true
		}
		if right.Type == FallingWater && right.VelX == 0 && right.ReachedWallLeft {
			next.ReachedWallRight = true
		}
		if next.ReachedWallLeft && next.ReachedWallRight {
			next.Type = StableWater
		}
	}

	// Falling water starting to move sideways
	if prev.Type == FallingWater &&
		next.Type == FallingWater &&
		next.VelX == 0 &&
		(down.Type == StableWater || down.Type == Clay) {
		if left.Type == Clay && right.Type == Sand {
			next.VelX = 1
		} else if left.Type == Sand && right.Type == Clay {
			next.VelX = -1
		} else if (left.Type == Sand || left.Type == FallingWater) &&
			(right.Type == Sand || right.Type == FallingWater) {
			if rand.Intn(2) == 1 {
				next.VelX = 1
			} else {
				next.VelX = -1
			}
		}
	}

	// Falling water stopping since it hits something
	if next.Type == None && prev.Type == FallingWater && (down.Type == StableWater || down.Type == Clay) && prev.VelX != 0 {
		if (left.Type == StableWater || left.Type == Clay) && (right.Type == StableWater || right.Type == Clay) {
			next = prev
			next.VelX = 0
		} else if prev.VelX == -1 && (left.Type == Clay || left.Type == StableWater || left.Type == FallingWater) {
			next = prev
			next.VelX = 0
		} else if prev.VelX == 1 && (right.Type == Clay || right.Type == StableWater || right.Type == FallingWater) {
			next = prev
			next.VelX = 0
		}
	}

	// Falling water moving sideways
	if next.Type == None && prev.Type == Sand && right.Type == FallingWater && right.VelX == -1 &&
		(downright.Type == Clay || downright.Type == StableWater) {
		next = right
	} else if next.Type == None && prev.Type == Sand && left.Type == FallingWater && left.VelX == 1 &&
		(downleft.Type == Clay || downleft.Type == StableWater) {
		next = left
	}

	// Squeezed falling water becomes stable
	if next.Type == FallingWater &&
		(down.Type == Clay || down.Type == StableWater) &&
		(left.Type == Clay || left.Type == StableWater) &&
		(right.Type == Clay || right.Type == StableWater) {
		next.Type = StableWater
		next.VelX = 0
	}

	// Fallback to sand
	if next.Type == None {
		next.Type = Sand
	}

	return next
}

func Next(A State) State {
	B := NewState()
	B[0][500].Type = Spring
	for y := 1; y < N-1; y++ {
		for x := 1; x < N-1; x++ {
			prev := A[y][x]
			up := A[y-1][x]
			down := A[y+1][x]
			left := A[y][x-1]
			right := A[y][x+1]
			downleft := A[y+1][x-1]
			downright := A[y+1][x+1]

			B[y][x] = NextTile(prev, up, down, left, right, downleft, downright)
		}
	}
	return B
}

func Print(s State, xmin, xmax, ymin, ymax int) {
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			if s[y][x].Type == FallingWater && s[y][x].VelX == -1 {
				fmt.Printf("<")
			} else if s[y][x].Type == FallingWater && s[y][x].VelX == 1 {
				fmt.Printf(">")
			} else {
				fmt.Printf("%c", s[y][x].Type)
			}
		}
		fmt.Println()
	}
}

type Pos struct {
	Y, X int
}

func main() {
	s, xmin, xmax, ymin, ymax := ReadInput()
	_, _, _, _ = xmin, xmax, ymin, ymax

	waterTiles := map[Pos]struct{}{}
	for i := 0; i < 30000; i++ {
		//Print(s, xmin, xmax, ymin, ymax)
		s = Next(s)

		for y := 0; y < ymax+1; y++ {
			for x := 0; x < N; x++ {
				t := s[y][x].Type
				if t == FallingWater || t == StableWater {
					waterTiles[Pos{y, x}] = struct{}{}
				}
			}
		}
		//for y := ymin; y <= ymax+5; y++ {
		//	for x := xmin; x <= xmax; x++ {
		//		if _, ok := waterTiles[Pos{y, x}]; ok {
		//			fmt.Printf("@")
		//			continue
		//		}
		//		if s[y][x].Type == FallingWater && s[y][x].VelX == -1 {
		//			fmt.Printf("<")
		//		} else if s[y][x].Type == FallingWater && s[y][x].VelX == 1 {
		//			fmt.Printf(">")
		//		} else {
		//			fmt.Printf("%c", s[y][x].Type)
		//		}
		//	}
		//	fmt.Println()
		//}
		fmt.Println(len(waterTiles))
	}
}
