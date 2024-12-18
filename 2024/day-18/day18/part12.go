package day18

import (
	"embed"
	"fmt"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/priorityqueue2"
)

//go:embed input*.txt
var Input embed.FS

type Pos struct {
	Y, X int
}

type Corrupted map[Pos]bool

func ReadInput(inputName string, n int, numBytes int) []Pos {
	positions := []Pos{}
	nums := aocutil.FSGetIntsInStringLines(Input, inputName)
	for _, ns := range nums {
		positions = append(positions, Pos{ns[1], ns[0]})
	}

	return positions
}

func PrintMap(m [][]byte) {
	for _, row := range m {
		fmt.Println(string(row))
	}
}

type State struct {
	Pos Pos
}

func (s State) ID(Nx int) int {
	return s.Pos.Y*Nx + s.Pos.X
}

func Right(p Pos) Pos {
	return Pos{p.Y, p.X + 1}
}

func Down(p Pos) Pos {
	return Pos{p.Y + 1, p.X}
}

func Left(p Pos) Pos {
	return Pos{p.Y, p.X - 1}
}

func Up(p Pos) Pos {
	return Pos{p.Y - 1, p.X}
}

const Inf = 9999999

func Part12(inputName string, n, numBytes int) (int, int) {
	corruptedPositions := ReadInput(inputName, n, numBytes)

	corrupted := make(map[Pos]bool)

	// Add walls
	for i := -1; i <= n; i++ {
		corrupted[Pos{-1, i}] = true // top
		corrupted[Pos{n, i}] = true  // bottom
		corrupted[Pos{i, -1}] = true // left
		corrupted[Pos{i, n}] = true  // right
	}

	// Add corrupted positions one by one
	numCorrupt := 0
	for _, newCorruptedPos := range corruptedPositions {
		corrupted[newCorruptedPos] = true
		numCorrupt++

		// Dijkstra/BFS for each of them.
		dist := map[Pos]int{{0, 0}: 0}
		pq := priorityqueue2.New[Pos, int]()
		pq.Push(Pos{0, 0}, 0)
		for pq.Len() > 0 {
			currPos, currDist := pq.PopPri()
			if dist[currPos] < currDist {
				continue
			}

			for _, neighbor := range []Pos{Right(currPos), Down(currPos), Left(currPos), Up(currPos)} {
				if corrupted[neighbor] {
					continue
				}
				neighborDist, ok := dist[neighbor]
				if !ok {
					neighborDist = Inf
				}
				altDist := currDist + 1
				if altDist < neighborDist {
					pq.Push(neighbor, altDist)
					dist[neighbor] = altDist
				}
			}
		}

		if numCorrupt == numBytes {
			fmt.Println(dist[Pos{n - 1, n - 1}])
		}

		if val, ok := dist[Pos{n - 1, n - 1}]; val == Inf || !ok {
			fmt.Println("Part 2:", newCorruptedPos.X, newCorruptedPos.Y)
			break
		}
	}

	// print := func() {
	// 	for y := 0; y < n; y++ {
	// 		for x := 0; x < n; x++ {
	// 			if corrupted[Pos{y, x}] {
	// 				fmt.Print("###")
	// 			} else {
	// 				d, ok := dist[Pos{y, x}]
	// 				if !ok {
	// 					fmt.Print("   ")
	// 				} else {
	// 					fmt.Printf("%3d", d)
	// 				}
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	return 0, 0
}
