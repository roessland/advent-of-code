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
	One   uint64
	Float uint64
}

func NewMask(mask string) Mask  {
	ones := uint64(0)
	floating := uint64(0)
	for i, val := range mask {
		shift := len(mask)-i-1
		switch val {
		case '0':
			continue
		case '1':
			ones = ones | (1<<shift)
		case 'X':
			floating = floating | (1<<shift)
		default:
			log.Fatalf("unknown bit %c", val)
		}
	}
	return Mask{ones, floating}
}

func (m Mask) String() string {
	return fmt.Sprintf("{\n\t1: %010b\n\tX: %010b\n}", m.One, m.Float)
}

func (m Mask) Decode(addr uint64) []uint64 {
	addr = addr | m.One

	if m.Float == 0 {
		return []uint64{addr}
	}

	var addrs []uint64
	// linear search for least significant floating bit
	for shift := 0; shift < 36; shift++ {
		if m.Float& (1<<shift) != 0 {
			m.Float ^= 1<<shift // remove floating bit we found

			// floating bit becomes zero, changing addr
			addrs = append(addrs, Mask{
				One:   m.One,
				Float: m.Float,
			}.Decode(^((^addr) | (1<<shift)))...)

			// floating bit becomes one, changing addr
			addrs = append(addrs, Mask{
				One:   m.One,
				Float: m.Float,
			}.Decode(addr | (1<<shift))...)

			return addrs
		}
	}
	panic("should never happen)")
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
	var mem = make(map[uint64]uint64)

	for scanner.Scan() {
		line := scanner.Text()
		maskMatches := maskRe.FindAllStringSubmatch(line, -1)
		memMatches := memRe.FindAllStringSubmatch(line, -1)

		if maskMatches != nil {

			maskStr := maskMatches[0][1]
			mask = NewMask(maskStr)

		} else if memMatches != nil {

			memAddrInt, err := strconv.Atoi(memMatches[0][1])
			memAddr := uint64(memAddrInt)
			if err != nil {
				log.Fatalf("cant atoi mem address %s: %s", memMatches[0][1], err)
			}
			valInt, err := strconv.Atoi(memMatches[0][2])
			val := uint64(valInt)
			if err != nil {
				log.Fatalf("cant atoi val %s: %s", memMatches[0][2], err)
			}

			for _, decodedAddr := range mask.Decode(memAddr) {
				mem[decodedAddr] = val
			}

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