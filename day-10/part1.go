package main

import "fmt"
import "strings"

func LookAndSay(s0 string) string {
	parts := make([]string, 0)
	for i0, i1, prev := 0, 0, s0[0]; ; i1++ {
		if i1 == len(s0) {
			parts = append(parts, s0[i0:i1])
			break
		}
		if s0[i1] != prev {
			parts = append(parts, s0[i0:i1])
			prev = s0[i1]
			i0 = i1
		}
	}

	for i, _ := range parts {
		parts[i] = fmt.Sprintf("%d%c", len(parts[i]), parts[i][0])
	}
	return strings.Join(parts, "")
}

func main() {
	s := "1113222113"
	for i := 0; i < 40; i++ {
		s = LookAndSay(s)
	}
	fmt.Printf("%v\n", len(s))
	for i := 0; i < 10; i++ {
		s = LookAndSay(s)
	}
	fmt.Printf("%v\n", len(s))
}
