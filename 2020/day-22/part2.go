package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card int

// Bottom has highest index, top has index 0
type Deck struct {
	Cards []Card
}

func (deck *Deck) Copy() *Deck {
	c := Deck{}
	c.Cards = make([]Card, len(deck.Cards))
	copy(c.Cards, deck.Cards)
	return &c
}

func (deck *Deck) CopyN(N int) *Deck {
	c := Deck{}
	c.Cards = make([]Card, N)
	copy(c.Cards, deck.Cards)
	return &c
}

func (deck *Deck) Hash() string {
	var buf bytes.Buffer
	for _, card := range deck.Cards {
		buf.WriteByte(byte(card))
	}
	return buf.String()
}

func (deck *Deck) InsertBottom(card Card) {
	deck.Cards = append(deck.Cards, card)
}

func (deck *Deck) InsertTop(card Card) {
	deck.Cards = append([]Card{card}, deck.Cards...)
}

func (deck *Deck) PopTop() Card {
	if len(deck.Cards) == 0 {
		panic("attempted pop from empty deck")
	}
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

type Player struct {
	Deck *Deck
}

func (p *Player) Score() int {
	score := 0
	for i, j := 1, len(p.Deck.Cards)-1; j >= 0; i, j = i+1, j-1 {
		score += int(p.Deck.Cards[j])*i
	}
	return score
}

type GlobalState struct {
	HasAlreadyPlayed map[string]bool
	StartDeck1       *Deck
	StartDeck2       *Deck
	GameId int
}

func (g *GlobalState) GetGameId() int {
	fmt.Printf(".")
	g.GameId++
	return g.GameId
}

type Game struct {
	GlobalState *GlobalState
	Player1 *Player
	Player2 *Player
	Winner *Player
}

func (game *Game) NewSubGame(card1, card2 Card) *Game {
	c := Game{}
	c.GlobalState = game.GlobalState
	c.Player1 = &Player{Deck: game.Player1.Deck.CopyN(int(card1))}
	c.Player2 = &Player{Deck: game.Player2.Deck.CopyN(int(card2))}
	c.Winner = nil
	return &c
}

func (game *Game) Hash() string {
	return fmt.Sprintf("1:%s===2:%s", game.Player1.Deck.Hash(), game.Player2.Deck.Hash())
}

func ReadInput(filename string) *Game {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	game := Game{
		Player1: &Player{Deck: &Deck{}},
		Player2: &Player{Deck: &Deck{}},
	}
	player := game.Player1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Player") {
			continue
		}
		if line == "" {
			player = game.Player2
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		player.Deck.Cards = append(player.Deck.Cards, Card(n))
	}
	return &game
}

func (game *Game) Play() int {
	//gameId := game.GlobalState.GetGameId()
	//fmt.Printf("=== Game %d ===\n", gameId)

	seenRounds := make(map[string]bool)
	round := 1
	for len(game.Player1.Deck.Cards) > 0 && len(game.Player2.Deck.Cards) > 0 {
		hash := game.Hash()
		if seenRounds[hash] {
			game.Winner = game.Player1
			//fmt.Printf("Player 1 wins game %d (infinite game detected)!\n",  gameId)
			return 1
		}
		seenRounds[hash] = true
		//fmt.Printf("-- Round %d (Game %d) --\n", round, gameId)
		//fmt.Printf("Player 1's deck: %v\n", game.Player1.Deck.Cards)
		//fmt.Printf("Player 2's deck: %v\n", game.Player2.Deck.Cards)
		card1 := game.Player1.Deck.PopTop()
		card2 := game.Player2.Deck.PopTop()
		//fmt.Printf("Player 1 plays: %v\n", card1)
		//fmt.Printf("Player 2 plays: %v\n", card2)


		var winner int
		if len(game.Player1.Deck.Cards) >= int(card1) && len(game.Player2.Deck.Cards) >= int(card2) {
			//fmt.Println("Playing a sub-game to determine the winner...")
			subGame := game.NewSubGame(card1, card2)
			winner = subGame.Play()
		} else if card1 > card2 {
			winner = 1
		} else if card2 > card1 {
			winner = 2
		} else {
			panic("nope!")
		}

		//fmt.Printf("Player %d wins round %d of game %d!\n", winner, round, gameId)

		if winner == 1 {
			game.Player1.Deck.InsertBottom(card1)
			game.Player1.Deck.InsertBottom(card2)
		} else if winner == 2 {
			game.Player2.Deck.InsertBottom(card2)
			game.Player2.Deck.InsertBottom(card1)
		} else {
			panic("didnt plan for this")
		}
		round++
	}



	if len(game.Player1.Deck.Cards) > 0 {
		game.Winner = game.Player1
		//fmt.Printf("Player 1 wins game %d!\n",  gameId)

		return 1
	} else {
		game.Winner = game.Player2
		//fmt.Printf("Player 2 wins game %d!\n",  gameId)
		return 2
	}
}

func main () {
	game := ReadInput("input.txt")
	globalState := GlobalState{
		HasAlreadyPlayed: make(map[string]bool),
		StartDeck1:       game.Player1.Deck.Copy(),
		StartDeck2:       game.Player2.Deck.Copy(),
	}
	game.GlobalState = &globalState
	game.Play()
	fmt.Println("Part 2:", game.Winner.Score())
}
