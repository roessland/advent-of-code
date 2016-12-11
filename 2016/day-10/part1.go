package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

type Bot struct {
	In   chan int
	Lo   chan int
	Hi   chan int
	Name string
}

func NewBot(name string) *Bot {
	return &Bot{make(chan int), nil, nil, name}
}

func (bot *Bot) Run() {
	for {
		A := <-bot.In
		B := <-bot.In

		if (A == 61 && B == 17) || (A == 17 && B == 61) {
			fmt.Println("Bot", bot.Name, "compares 61 and 17")
		}
		if A <= B {
			bot.Lo <- A
			bot.Hi <- B
		} else {
			bot.Lo <- B
			bot.Hi <- A
		}
	}
}

func GetOrCreateBot(bots map[string]*Bot, name string) *Bot {
	if _, ok := bots[name]; !ok {
		bots[name] = NewBot(name)
	}
	return bots[name]
}

func GetOrCreateOutput(outputs map[string]chan int, name string) chan int {
	if _, ok := outputs[name]; !ok {
		outputs[name] = make(chan int, 100) // has to be buffered to prevent deadlock
	}
	return outputs[name]
}

func main() {
	bots := make(map[string]*Bot)
	outputs := make(map[string]chan int)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[0] == "bot" {
			giver := GetOrCreateBot(bots, fields[1])
			var loChan, hiChan chan int
			if fields[5] == "bot" {
				loChan = GetOrCreateBot(bots, fields[6]).In
			} else {
				loChan = GetOrCreateOutput(outputs, fields[6])
			}
			if fields[10] == "bot" {
				hiChan = GetOrCreateBot(bots, fields[11]).In
			} else {
				hiChan = GetOrCreateOutput(outputs, fields[11])
			}
			giver.Lo = loChan
			giver.Hi = hiChan
			go giver.Run()

		} else if fields[0] == "value" {
			value := Atoi(fields[1])
			receiver := GetOrCreateBot(bots, fields[5])
			go func() { receiver.In <- value }()
		}
	}
	a, b, c := <-outputs["0"], <-outputs["1"], <-outputs["2"]
	fmt.Printf("out[0]*out[1]*out[2] = %v * %v * %v = %v\n", a, b, c, a*b*c)
}
