package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const Ny = 137
const Nx = 139

type Board [Ny][Nx]byte

func (b *Board) Print() {
	for y := 0; y < Ny; y++ {
		fmt.Println(string(b[y][:]))
	}
}

func rem(a, b int) int {
	rest := a % b
	if rest < 0 {
		return rest + b
	} else {
		return rest
	}
}

func (b *Board) At(y, x int) byte {
	return b[rem(y, Ny)][rem(x, Nx)]
}

func (b *Board) Step() Board {
	c := *b
	for y := 0; y < Ny; y++ {
		for x := 0; x < Nx; x++ {
			if b.At(y, x-1) == '>' && b.At(y, x) == '.' {
				c[y][x] = '>'
			} else if b.At(y, x) == '>' && b.At(y, x+1) == '.' {
				c[y][x] = '.'
			} else if b.At(y, x) == '>' && b.At(y, x+1) != '.' {
				c[y][x] = '>'
			}
		}
	}

	d := c
	for y := 0; y < Ny; y++ {
		for x := 0; x < Nx; x++ {
			if c.At(y-1, x) == 'v' && c.At(y, x) == '.' {
				d[y][x] = 'v'
			} else if c.At(y, x) == 'v' && c.At(y+1, x) == '.' {
				d[y][x] = '.'
			} else if c.At(y, x) == 'v' && c.At(y+1, x) != '.' {
				d[y][x] = 'v'
			}
		}
	}
	return d
}

func ReadInput() Board {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var board Board
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		board[i] = *(*[Nx]byte)(line)
		i++
	}
	return board
}

func main() {
	board := ReadInput()

	for i := 1; ; i++ {
		nextBoard := board.Step()
		if nextBoard == board {
			fmt.Println(i)
			break
		}
		board = nextBoard
	}
}
