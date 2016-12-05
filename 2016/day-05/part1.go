package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	input := "wtnhxymk"
	pass := make([]rune, 0)

	for i := 0; i < 1000000000000; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v", input, i))))
		if hash[0:5] == "00000" {
			pass = append(pass, rune(hash[5]))
			fmt.Println(string(pass))
			if len(pass) == 8 {
				break
			}
		}
	}
	fmt.Println("There you go!")
}
