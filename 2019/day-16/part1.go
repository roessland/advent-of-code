package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func FFT(s []int) []int {
	kernel := []int{0, 1, 0, -1}
	output := make([]int, len(s))
	for i, n := range s {
		for j := range s {
			kernelValue := kernel[(i+1)/(j+1)%len(kernel)]
			output[j] += n * kernelValue
		}
	}
	for i := range output {
		output[i] %= 10
		if output[i] < 0 {
			output[i] = -output[i]
		}
	}
	return output
}

func IteratedFFT(s []int) []int {
	for i := 0; i < 100; i++ {
		s = FFT(s)
	}
	return s
}

func Input(str string) []int {
	s := []int{}
	for _, c := range str {
		n := int(c - '0')
		if n < 10 && n >= 0 {
			s = append(s, n)
		}
	}
	return s
}

func Input10000(str string) []int {
	s := []int{}
	for i := 0; i < 2; i++ {
		for _, c := range str {
			n := int(c - '0')
			if n < 10 && n >= 0 {
				s = append(s, n)
			}
		}
	}
	return s
}

func main() {
	fmt.Println(FFT(Input("00000001")))
	fmt.Println(FFT(Input10000("00000001")))

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	o := IteratedFFT(Input(string(input)))
	fmt.Printf("Part 1: %d%d%d%d%d%d%d%d\n", o[0], o[1], o[2], o[3], o[4], o[5], o[6], o[7])
	o = IteratedFFT(Input10000(string(input)))
	fmt.Printf("Part 2: %d%d%d%d%d%d%d%d\n", o[0], o[1], o[2], o[3], o[4], o[5], o[6], o[7])
}
