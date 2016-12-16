package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var cache map[int]string

func Hash(n int) string {
	cached, ok := cache[n]
	if ok {
		return cached
	}
	salt := "cuanljph"
	hash := md5.Sum([]byte(salt + strconv.Itoa(n)))
	digest := hex.EncodeToString(hash[:])

	for i := 0; i < 2016; i++ {
		hash = md5.Sum([]byte(digest))
		digest = hex.EncodeToString(hash[:])
	}
	cache[n] = digest
	return digest
}

func GetTriplet(hash string) (byte, bool) {
	for i, j, k := 0, 1, 2; k < len(hash); i, j, k = i+1, j+1, k+1 {
		if hash[i] == hash[j] && hash[j] == hash[k] {
			return hash[i], true
		}
	}
	return 0, false
}

func HasQuintet(hash string, char byte) bool {
	for a, b, c, d, e := 0, 1, 2, 3, 4; e < len(hash); a, b, c, d, e = a+1, b+1, c+1, d+1, e+1 {
		if char == hash[a] && hash[a] == hash[b] && hash[b] == hash[c] && hash[c] == hash[d] && hash[d] == hash[e] {
			return true
		}
	}
	return false
}

func main() {
	cache = make(map[int]string)
	fmt.Println("done")
	numKeys := 0
	for n := 0; ; n++ {
		c, hasTriplet := GetTriplet(Hash(n))
		if hasTriplet {
			hasQuintet := false
			for i := n + 1; i <= n+1000; i++ {
				if HasQuintet(Hash(i), c) {
					hasQuintet = true
					break
				}
			}
			if hasQuintet {
				numKeys++
				fmt.Println(numKeys, n)
				if numKeys == 64 {
					break
				}
			}
		}
	}
}
