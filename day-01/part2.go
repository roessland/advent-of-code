package main

import "io/ioutil"
import "fmt"

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	floor := 0
	for i, button := range []rune(string(buf)) {
		fmt.Printf("%T\n", button)
		if button == '(' {
			floor++
		} else if button == ')' {
			floor--
		} else {
			fmt.Printf("Unknown: %s %x %d\n", button, button, button)
		}
		if floor == -1 {
			fmt.Printf("========= %d =======\n", i+1)
			panic("lol")
		}

	}

	fmt.Printf("Final floor: %d\n", floor)
}
