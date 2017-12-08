package main

import "fmt"

const NumBanks int = 16

type BankConfig struct {
	blocks [NumBanks]uint8
}

func (bc BankConfig) Redistribute() BankConfig {
	i := bc.MostBlocksIndex()
	blocksLeft := bc.blocks[i]
	bc.blocks[i] = 0
	for blocksLeft > 0 {
		i = (i + 1) % NumBanks
		bc.blocks[i]++
		blocksLeft--
	}
	return bc
}

func (bc BankConfig) MostBlocksIndex() int {
	minBlocks := bc.blocks[0]
	minIndex := 0
	for i := 1; i < NumBanks; i++ {
		if bc.blocks[i] > minBlocks {
			minIndex = i
			minBlocks = bc.blocks[i]
		}
	}
	return minIndex
}

func main() {
	bc := BankConfig{[16]uint8{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}}
	seen := map[BankConfig]int{}
	i := 1
	for seen[bc] == 0 {
		seen[bc] = i
		bc = bc.Redistribute()
		i++
	}
	fmt.Println("Part 1: ", len(seen))
	fmt.Println("Part 2: ", i-seen[bc])
}
