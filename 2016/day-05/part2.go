package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := "wtnhxymk"
	pass := make([]rune, 20)
	for i := 0; i < len(pass); i++ {
		pass[i] = '_'
	}

	for i := 0; i < 1000000000000; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v", input, i))))
		if hash[0:5] == "00000" {
			pos, err := strconv.Atoi(string(hash[5]))
			if err == nil && pos < 8 && pass[pos] == '_' {
				pass[pos] = rune(hash[6])
			}
			fmt.Println(string(pass))
			fmt.Println("********")
		}
	}
	fmt.Println("There you go!")
}
