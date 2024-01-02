package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

var palette [][]string

func init() {
	generatePalette()
}

var nextColor = -1

func getColor(blockID int, distance int) lipgloss.Color {
	return lipgloss.Color(palette[blockID%len(palette)][distance])
}

func generatePalette() {
	baseColors, err := colorful.HappyPalette(10)
	if err != nil {
		panic(err)
	}

	palette = make([][]string, len(baseColors))

	for i, c0 := range baseColors {
		for j := 0; j <= 9; j++ {
			black, _ := colorful.Hex("#333333")
			c := c0.BlendHcl(black, 0.8*float64(j)/9)
			palette[i] = append(palette[i], c.Hex())
		}
	}
}
