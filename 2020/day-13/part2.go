package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan() // skip irrelevant first line
	scanner.Scan()
	line2 := scanner.Text()

	ms := []int64{} // ids
	msStr := strings.Split(line2, ",")
	M := int64(1)   // product of ids
	as := []int64{} // RHS in Chinese Remainder Theorem
	for offset, mStr := range msStr {
		if mStr == "x" {
			continue
		}
		m, err := strconv.Atoi(mStr)
		if err != nil {
			log.Fatal(err)
		}
		ms = append(ms, int64(m))
		// Make sure the RHS is positive
		as = append(as, (-int64(offset)+int64(10)*int64(m))%int64(m))
		M *= int64(m)
	}

	// chinese remainder theorem
	var t int64
	for i, m := range ms {
		a := as[i]
		b := M / m // prod of all _other_ ms
		bInv := new(big.Int).ModInverse(big.NewInt(b), big.NewInt(m)).Int64()
		t = (t + a*b*bInv) % M
	}
	fmt.Println(t)
}
