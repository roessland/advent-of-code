package main

import (
	"fmt"
)

func Next(a []uint8) []uint8 {
	a0b := make([]uint8, len(a)*2+1)
	copy(a0b, a)
	a0b[len(a)] = 0
	for i := 0; i < len(a); i++ {
		a0b[len(a)+1+i] = a[len(a)-1-i] ^ 1
	}
	return a0b
}

func Checksum(in []uint8) []uint8 {
	out := make([]uint8, len(in)/2)
	for i, j := 0, 1; j < len(in); i, j = i+2, j+2 {
		out[i/2] = (in[i] ^ in[j]) ^ 1
	}
	return out
}

func main() {
	//const diskSize = 272 // part 1
	const diskSize = 35651584 // part 2
	data := []uint8{0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0}
	for len(data) < diskSize {
		data = Next(data)
	}

	checksum := data[0:diskSize]
	for len(checksum)%2 == 0 {
		checksum = Checksum(checksum)
	}
	fmt.Println(checksum)

}
