package main

import (
	"fmt"
	"github.com/roessland/gopkg/mathutil"
)

type Die struct {
	prevRoll int
}

func (d *Die) Roll() int {
	d.prevRoll++
	return d.prevRoll
}

type Game struct {
	timesRolled int
	positions   [2]int
	scores      [2]int
	numSpaces   int
	prevPlayer  int
	die         *Die
}

func (g *Game) DoTurn() {
	g.prevPlayer = (g.prevPlayer + 1) % 2
	currPlayer := g.prevPlayer
	pos := g.positions[currPlayer]
	g.timesRolled += 3
	roll := g.die.Roll() + g.die.Roll() + g.die.Roll()
	pos = ((pos-1)+roll)%g.numSpaces + 1
	g.positions[currPlayer] = pos
	g.scores[currPlayer] += pos
}

func main() {
	game := Game{
		timesRolled: 0,
		positions:   [2]int{4, 3},
		numSpaces:   10,
		die:         &Die{},
		prevPlayer:  1,
	}

	for i := 0; i < 5000; i++ {

		game.DoTurn()

		if game.scores[0] >= 1000 {
			fmt.Println("plkayer 1 won")
			break
		}
		if game.scores[1] >= 1000 {
			fmt.Println("plkayer 2 won")
			break
		}
	}

	fmt.Printf("%#v\n", game)
	fmt.Println(game.timesRolled * mathutil.MinInt(game.scores[0], game.scores[1]))
}
