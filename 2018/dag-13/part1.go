package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Pair struct {
	X, Y int
}

type Cart struct {
	PosX, PosY  int
	DirX, DirY  int
	Crashed     bool
	TurnCounter int
}

func Sort(carts []Cart) {
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].PosY == carts[j].PosY {
			return carts[i].PosX < carts[j].PosX
		}
		return carts[i].PosY < carts[j].PosY
	})
}

func Turn(cart *Cart) {
	switch cart.TurnCounter {
	case 0:
		TurnLeft(cart)
	case 1:
		break
	case 2:
		TurnRight(cart)
	}
	cart.TurnCounter = (cart.TurnCounter + 1) % 3
}

func TurnLeft(cart *Cart) {
	if cart.DirX == 1 {
		cart.DirX = 0
		cart.DirY = -1
	} else if cart.DirX == -1 {
		cart.DirX = 0
		cart.DirY = 1
	} else if cart.DirY == 1 {
		cart.DirX = 1
		cart.DirY = 0
	} else if cart.DirY == -1 {
		cart.DirX = -1
		cart.DirY = 0
	}
}

func TurnRight(cart *Cart) {
	TurnLeft(cart)
	cart.DirX *= -1
	cart.DirY *= -1
}

func main() {
	carts := []Cart{}
	mapp := [][]byte{}
	buf, _ := ioutil.ReadFile("example2.txt")
	for y, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			break
		}
		mapp = append(mapp, make([]byte, len(line)))
		for x, c := range line {
			mapp[y][x] = ' '
			switch c {
			case '<':
				carts = append(carts, Cart{x, y, -1, 0, false, 0})
			case '>':
				carts = append(carts, Cart{x, y, 1, 0, false, 0})
			case '^':
				carts = append(carts, Cart{x, y, 0, -1, false, 0})
			case 'v':
				carts = append(carts, Cart{x, y, 0, 1, false, 0})
			default:
				mapp[y][x] = line[x]
			}
		}
	}
out:
	for k := 0; k < 200000; k++ {
		Sort(carts)

		cartByPos := map[Pair]*Cart{}
		for i, cart := range carts {
			cartByPos[Pair{cart.PosX, cart.PosY}] = &carts[i]
		}

		// Update carts
		for i := range carts {
			cart := &carts[i]
			delete(cartByPos, Pair{cart.PosX, cart.PosY})
			cart.PosX += cart.DirX
			cart.PosY += cart.DirY
			if mapp[cart.PosY][cart.PosX] == '+' {
				Turn(cart)
			}
			switch mapp[cart.PosY][cart.PosX] {
			case '+':
				Turn(cart)
			case '\\':
				if cart.DirX != 0 {
					TurnRight(cart)
				} else {
					TurnLeft(cart)
				}
			case '/':
				if cart.DirX != 0 {
					TurnLeft(cart)
				} else {
					TurnRight(cart)
				}
			}
			if cartByPos[Pair{cart.PosX, cart.PosY}] != nil {
				fmt.Println("crashed!")
				fmt.Printf("%d,%d\n", cart.PosX, cart.PosY)
				break out
			}
			cartByPos[Pair{cart.PosX, cart.PosY}] = cart
		}

		// Draw map
		for y, line := range mapp {
			for x, c := range line {
				_ = c
				if cartByPos[Pair{x, y}] != nil {
					//fmt.Printf("#")
				} else {
					//fmt.Printf("%c", c)
				}
			}
			//fmt.Printf("\n")
		}
	}
}
