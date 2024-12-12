package day09

import (
	"embed"
	"slices"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

func ReadInput(inputName string) []int {
	nums := []int{}
	for _, c := range aocutil.FSReadFile(Input, inputName) {
		nums = append(nums, int(c-'0'))
	}
	return nums
}

type FileBlock struct {
	FileID int
}

type Disk1 struct {
	FreeBlocks     []int
	FileBlocks     []FileBlock
	NumEmptyBlocks int
}

type Disk2 struct {
	FreeBlocks [][]int
	FileBlocks []FileBlock
}

func Decompress1(diskMap []int) Disk1 {
	freeBlocks := []int{}
	fileBlocks := []FileBlock{}
	numEmptyBlocks := 0
	fileID := 0
	for i := 0; i < len(diskMap); i++ {
		if i%2 == 0 {
			for range diskMap[i] {
				fileBlocks = append(fileBlocks, FileBlock{FileID: fileID})
			}
			fileID++
		} else {
			for range diskMap[i] {
				numEmptyBlocks++
				freeBlocks = append(freeBlocks, len(fileBlocks))
				fileBlocks = append(fileBlocks, FileBlock{FileID: -1})
			}
		}
	}
	slices.Reverse(freeBlocks)
	return Disk1{
		FreeBlocks:     freeBlocks,
		FileBlocks:     fileBlocks,
		NumEmptyBlocks: numEmptyBlocks,
	}
}

func (d Disk1) Checksum() int {
	checksum := 0
	for pos, fileBlock := range d.FileBlocks {
		checksum += pos * fileBlock.FileID
	}
	return checksum
}

func (d Disk2) Checksum() int {
	checksum := 0
	for pos, fileBlock := range d.FileBlocks {
		if fileBlock.FileID == -1 {
			continue
		}
		checksum += pos * fileBlock.FileID
	}
	return checksum
}

func Defrag1(disk *Disk1) {
	for disk.NumEmptyBlocks > 0 {
		// Pop last block
		srcBlock := disk.FileBlocks[len(disk.FileBlocks)-1]
		disk.FileBlocks = disk.FileBlocks[:len(disk.FileBlocks)-1]

		// Skip if empty
		if srcBlock.FileID == -1 {
			disk.NumEmptyBlocks--
			continue
		}

		// Pop first free block
		dst := disk.FreeBlocks[len(disk.FreeBlocks)-1]
		disk.FreeBlocks = disk.FreeBlocks[:len(disk.FreeBlocks)-1]
		disk.NumEmptyBlocks--

		// Assign
		disk.FileBlocks[dst] = srcBlock
	}
}

func Decompress2(diskMap []int) Disk2 {
	freeBlocks := make([][]int, 10)
	for i := 0; i <= 9; i++ {
		freeBlocks[i] = []int{}
	}

	fileBlocks := []FileBlock{}
	fileID := 0
	for i := 0; i < len(diskMap); i++ {
		if i%2 == 0 {
			nFile := diskMap[i]
			for range nFile {
				fileBlocks = append(fileBlocks, FileBlock{FileID: fileID})
			}
			fileID++
		} else {
			nFree := diskMap[i]
			freeBlocks[nFree] = append(freeBlocks[nFree], len(fileBlocks))
			for range nFree {
				fileBlocks = append(fileBlocks, FileBlock{FileID: -1})
			}
		}
	}
	for _, freeBlock := range freeBlocks {
		slices.Reverse(freeBlock)
	}
	return Disk2{
		FreeBlocks: freeBlocks,
		FileBlocks: fileBlocks,
	}
}

func (d Disk2) String() string {
	s := []byte{}
	for _, block := range d.FileBlocks {
		if block.FileID == -1 {
			s = append(s, '.')
		} else {
			s = append(s, byte(block.FileID)+'0')
		}
	}
	s = append(s, '\n')
	asdf := make([]byte, len(d.FileBlocks))
	for i := range asdf {
		asdf[i] = ' '
	}
	for sz, bySize := range d.FreeBlocks {
		for _, block := range bySize {
			for i := 0; i < sz; i++ {
				asdf[block+i] = '0' + byte(sz)
			}
		}
	}
	s = append(s, asdf...)

	return string(s)
}

func GetExtent(disk *Disk2, endInclusive int) (int, int) {
	endID := disk.FileBlocks[endInclusive].FileID
	start := endInclusive
	for start-1 >= 0 && disk.FileBlocks[start-1].FileID == endID {
		start--
	}
	return start, endInclusive
}

func Defrag2(disk *Disk2) {
	hasMoved := make(map[int]bool)
	i := len(disk.FileBlocks) - 1
	for i >= 0 {
		if disk.FileBlocks[i].FileID == -1 {
			i--
			continue
		}
		srcStart, srcEnd := GetExtent(disk, i)

		// Find leftmost possible free block big enough
		size := srcEnd - srcStart + 1
		dstStart := 99999
		var minDstFree int
		minDstStart := 99999
		for dstFree := size; dstFree < 10; dstFree++ {
			if len(disk.FreeBlocks[dstFree]) == 0 {
				continue
			}
			dstStart = disk.FreeBlocks[dstFree][len(disk.FreeBlocks[dstFree])-1]
			if dstStart > srcStart {
				continue
			}
			if dstStart <= minDstStart {
				minDstStart = dstStart
				minDstFree = dstFree
				continue
			}
		}

		if dstStart > 9999999 || minDstFree == 0 {
			// Go to next file
			i = srcStart - 1
			continue // no free blocks
		}

		// Pop block of this size
		dstStart = disk.FreeBlocks[minDstFree][len(disk.FreeBlocks[minDstFree])-1]
		disk.FreeBlocks[minDstFree] = disk.FreeBlocks[minDstFree][:len(disk.FreeBlocks[minDstFree])-1]
		// fmt.Println(disk)

		// Add back remaining space
		remaining := minDstFree - size
		if remaining > 0 {
			disk.FreeBlocks[remaining] = append(disk.FreeBlocks[remaining], dstStart+size)
			slices.Sort(disk.FreeBlocks[remaining])
			slices.Reverse(disk.FreeBlocks[remaining])
		}

		// Move
		if hasMoved[disk.FileBlocks[srcStart].FileID] {
			continue
		}
		hasMoved[disk.FileBlocks[srcStart].FileID] = true
		for j := srcStart; j <= srcEnd; j++ {
			dst := dstStart + j - srcStart
			disk.FileBlocks[dst] = disk.FileBlocks[j]
			disk.FileBlocks[j] = FileBlock{FileID: -1}
		}

		// Crop

		// Go to next file
		i = srcStart - 1
	}
}

func Part12(inputName string) (int, int) {
	diskMap := ReadInput(inputName)
	disk1 := Decompress1(diskMap)
	Defrag1(&disk1)

	disk2 := Decompress2(diskMap)
	Defrag2(&disk2)

	return disk1.Checksum(), disk2.Checksum()
}
