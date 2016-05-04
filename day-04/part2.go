package main

import "fmt"
import "crypto/md5"

func main() {
	for i := 0; ; i++ {
		data := []byte(fmt.Sprintf("bgvyzdsv%d", i))
		s := fmt.Sprintf("%x", md5.Sum(data))
		if s[0] == '0' && s[1] == '0' && s[2] == '0' && s[3] == '0' && s[4] == '0' && s[5] == '0' {
			fmt.Printf("%s\n", string(data))
			fmt.Printf("%s\n", s)
			fmt.Printf("yeah %d", i)
			break
		}
	}

}
