package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const Black = 0
const White = 1
const Transparent = 2

// Image is a bunch of layers
type Image struct {
	Data   []int
	Width  int
	Height int
	Depth  int
}

// NewImage creates an image from a string of runes.
func NewImage(rd io.Reader, width, height int) Image {
	im := Image{
		Width:  width,
		Height: height,
	}
	b := bufio.NewReader(rd)
	for {
		r, _, err := b.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		num := int(r - '0')
		im.Data = append(im.Data, num)
	}
	im.Depth = len(im.Data) / (im.Width * im.Height)
	return im
}

// At gets the value at column i, row j, layer k
func (im Image) At(i, j, k int) int {
	return im.Data[k*(im.Width*im.Height)+j*im.Width+i]
}

// Set sets the value at column i, row j, layer k
func (im Image) Set(i, j, k int, val int) {
	if val != Transparent {
		im.Data[k*(im.Width*im.Height)+j*im.Width+i] = val
	}
}

// String creates a string representation of all the layers
func (im Image) String() string {
	chars := []rune{}
	for k := im.Depth - 1; k >= 0; k-- {
		for j := 0; j < im.Height; j++ {
			for i := 0; i < im.Width; i++ {
				val := im.At(i, j, k)
				switch val {
				case Black:
					chars = append(chars, '▓')
				case White:
					chars = append(chars, '░')
				case Transparent:
					chars = append(chars, ' ')
				}
			}
			chars = append(chars, '\n')
		}
		chars = append(chars, '\n')
	}
	return string(chars)
}

// Flatten returns a new image of the same dimensions but with just one layer.
// The layers are rendered back to front.
func (im Image) Flatten() Image {
	flattened := Image{
		Data:   make([]int, im.Width*im.Height),
		Width:  im.Width,
		Height: im.Height,
		Depth:  1,
	}
	for idx := range flattened.Data {
		flattened.Data[idx] = Transparent
	}
	for k := im.Depth - 1; k >= 0; k-- {
		for j := 0; j < im.Height; j++ {
			for i := 0; i < im.Width; i++ {
				flattened.Set(i, j, 0, im.At(i, j, k))
			}
		}
	}
	return flattened
}

// GetLayerData gets the raw data for layer k
func (im Image) GetLayerData(k int) []int {
	layerSize := im.Width * im.Height
	return im.Data[k*layerSize : (k+1)*layerSize]
}

// CountZeroesForLayer counts the black values in layer k
func (im Image) CountBlacksForLayer(k int) int {
	count := 0
	layer := im.GetLayerData(k)
	for _, val := range layer {
		if val == 0 {
			count++
		}
	}
	return count
}

func Part1(im Image) {
	blackCounts := make([]int, im.Depth)
	for k := 0; k < im.Depth; k++ {
		blackCounts[k] = im.CountBlacksForLayer(k)
	}
	k := ArgMin(blackCounts)
	layerData := im.GetLayerData(k)
	numWhite := 0
	numTransparent := 0
	for _, val := range layerData {
		if val == White {
			numWhite++
		} else if val == Transparent {
			numTransparent++
		}
	}
	fmt.Println(numWhite * numTransparent)
}

func ArgMin(s []int) int {
	minIdx := 0
	minVal := s[0]
	for idx, val := range s {
		if val < minVal {
			minIdx, minVal = idx, val
		}
	}
	return minIdx
}

func Part2(im Image) {
	flattened := im.Flatten()
	fmt.Println(flattened)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	width := 25
	height := 6
	im := NewImage(f, width, height)
	Part1(im)
	Part2(im)
}

/*
░░░░▓░▓▓▓░░░░▓▓░▓▓▓▓░▓▓░▓
▓▓▓░▓░▓▓▓░░▓▓░▓░▓▓▓▓░▓▓░▓
▓▓░▓▓▓░▓░▓░░░▓▓░▓▓▓▓░░░░▓
▓░▓▓▓▓▓░▓▓░▓▓░▓░▓▓▓▓░▓▓░▓
░▓▓▓▓▓▓░▓▓░▓▓░▓░▓▓▓▓░▓▓░▓
░░░░▓▓▓░▓▓░░░▓▓░░░░▓░▓▓░▓
*/
