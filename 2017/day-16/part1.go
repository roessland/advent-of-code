package main

import "fmt"
import "strconv"
import "bufio"
import "os"
import "strings"
import "log"

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func FindIndex(state []rune, r rune) int {
	i := 0
	for state[i] != r {
		i++
	}
	return i
}

func Spin(state []rune, X int) []rune {
	Y := len(state) - X
	first := state[0:Y]
	second := state[Y:len(state)]
	newState := append(second, first...)
	return newState
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	moves := strings.Split(scanner.Text(), ",")
	state := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	for _, move := range moves {
		if move[0] == 's' {
			state = Spin(state, Atoi(move[1:]))
		} else if move[0] == 'x' {
			params := strings.Split(move[1:], "/")
			A := Atoi(params[0])
			B := Atoi(params[1])
			state[A], state[B] = state[B], state[A]
		} else if move[0] == 'p' {
			params := strings.Split(move[1:], "/")
			A := rune(params[0][0])
			B := rune(params[1][0])
			i := FindIndex(state, A)
			j := FindIndex(state, B)
			state[i], state[j] = state[j], state[i]
		}
	}
	fmt.Println(string(state))
}
