package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Op(op rune, a, b int) int {
	switch op {
	case '*':
		return a * b
	case '+':
		return a + b
	default:
		panic("wut")
	}
}

type TokenType string

const NumberToken TokenType = "number"

const OpToken TokenType = "op"

type Token struct {
	Type TokenType
	Number int
	Op rune
}

func NewToken(r rune) Token {
	if '0' < r && r <= '9' {
		return Token{NumberToken, int(r-'0'), 0}
	}
	if r == '*' || r == '+' {
		return Token{OpToken, -1, r}
	}
	panic("fuck")
}

func Eval(rdr *bufio.Reader) Token {
	var tokens []Token
	for {
		r, _, err := rdr.ReadRune()
		if err == io.EOF {
			break
		} else if r == ' ' {
			continue
		} else if r == '(' {
			tokens = append(tokens, Eval(rdr))
		} else if r == ')' {
			break
		} else {
			tokens = append(tokens, NewToken(r))
		}
	}

	for _, phase := range []rune{'+', '*'} {
		for i := 1; i < len(tokens)-1; i++ {
			for len(tokens) > i && tokens[i].Op == phase {
				L := len(tokens)
				val := Token{NumberToken,Op(tokens[i].Op, tokens[i-1].Number, tokens[i+1].Number), 0}
				tokens[i-1] = val
				copy(tokens[i:L-2], tokens[i+2:L])
				tokens = tokens[:L-2]
			}
		}
	}

	return tokens[0]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		sum += Eval(bufio.NewReader(strings.NewReader(scanner.Text()))).Number
	}
	fmt.Println("Part 2:", sum)
}
