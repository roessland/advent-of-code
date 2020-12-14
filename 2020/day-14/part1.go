package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Mask struct {
	AndZeros uint64
	OrOnes uint64
}

func (m Mask) Apply(num uint64) uint64 {
	num = num & m.AndZeros
	num = num | m.OrOnes
	return num
}

func (m Mask) String() string {
	return fmt.Sprintf("{\n\t0: %b\n\t1: %b\n}", m.AndZeros, m.OrOnes)
}

// & with zeros, |
func computeMask(mask string) Mask  {
	//fmt.Println("computing mask for string", mask)
	zeros := uint64(1<<36-1)
	ones := uint64(0)
	for i, val := range mask {
		shift := len(mask)-i-1
		switch val {
		case 'X':
			continue
		case '0':
			zeros = ^(^zeros | (1<<shift))
		case '1':
			ones = ones | (1<<shift)
		default:
			log.Fatalf("unknown bit %c", val)
		}
	}
	return Mask{zeros, ones}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	maskRe := regexp.MustCompile(`^mask\s=\s([X01]+)$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\]\s=\s(\d+)$`)

	var mask Mask
	var mem = make([]uint64, 100000)

	for scanner.Scan() {
		line := scanner.Text()
		maskMatches := maskRe.FindAllStringSubmatch(line, -1)
		memMatches := memRe.FindAllStringSubmatch(line, -1)

		if maskMatches != nil {

			maskStr := maskMatches[0][1]
			mask = computeMask(maskStr)

		} else if memMatches != nil {

			memAddr, err := strconv.Atoi(memMatches[0][1])
			if err != nil {
				log.Fatalf("cant atoi mem address %s: %s", memMatches[0][1], err)
			}
			valInt, err := strconv.Atoi(memMatches[0][2])
			val := uint64(valInt)
			if err != nil {
				log.Fatalf("cant atoi val %s: %s", memMatches[0][2], err)
			}

			mem[memAddr] = mask.Apply(val)

		} else {
			log.Fatal("neither mask nor mem matched")
		}
	}

	var sum uint64
	for _, val := range mem {
		sum += val
	}
	fmt.Println(sum)
}