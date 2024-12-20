package day20

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/mathutil"
	"github.com/roessland/gopkg/priorityqueue2"
)

//go:embed input*.txt
var Input embed.FS

func ReadInput(inputName string) (m [][]byte) {
	return aocutil.FSReadLinesAsBytes(Input, inputName)
}

type Pos struct {
	X, Y int
}

func Neighbors(p Pos) []Pos {
	return []Pos{
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
	}
}

func CheatDsts1(p Pos) []Pos {
	return []Pos{
		{p.X - 1, p.Y - 1},
		{p.X + 1, p.Y - 1},
		{p.X - 1, p.Y + 1},
		{p.X + 1, p.Y + 1},
		{p.X - 2, p.Y},
		{p.X + 2, p.Y},
		{p.X, p.Y - 2},
		{p.X, p.Y + 2},
	}
}

func CheatDsts2(p Pos, radius int) []Pos {
	ps := []Pos{}
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if mathutil.AbsInt(x)+mathutil.AbsInt(y) > radius {
				continue
			}
			ps = append(ps, Pos{p.X + x, p.Y + y})
		}
	}
	return ps
}

func Dijkstra(m [][]byte, origin Pos) map[Pos]int {
	dist := map[Pos]int{origin: 0}

	distTo := func(p Pos) int {
		d, ok := dist[p]
		if ok {
			return d
		}
		return 9999999
	}

	q := priorityqueue2.New[Pos, int]()
	q.Push(origin, 0)
	for q.Len() > 0 {
		currPos, currDist := q.PopPri()
		if d, ok := dist[currPos]; ok && d < currDist {
			continue
		}
		for _, neighborPos := range Neighbors(currPos) {
			if m[neighborPos.Y][neighborPos.X] == '#' {
				continue
			}
			neighborDist := distTo(neighborPos)
			altDist := currDist + 1
			if altDist < neighborDist {
				dist[neighborPos] = altDist
				q.Push(neighborPos, altDist)
			}
		}
	}
	return dist
}

type Cheat struct {
	Src, Dst Pos
}

func Part1(m [][]byte, distToEnd map[Pos]int) int {
	width := len(m[0])
	height := len(m)
	cheats := map[Cheat]int{}
	for y, row := range m {
		for x := range row {
			if m[y][x] == '#' {
				continue
			}
			cheatSrc := Pos{x, y}
			for _, cheatDst := range CheatDsts1(cheatSrc) {
				if cheatDst.X < 0 || cheatDst.X >= width || cheatDst.Y < 0 || cheatDst.Y >= height {
					continue
				}
				if m[cheatDst.Y][cheatDst.X] == '#' {
					continue
				}
				saves := distToEnd[cheatSrc] - distToEnd[cheatDst] - 2
				if saves >= 100 {
					cheats[Cheat{cheatSrc, cheatDst}] = saves
				}
			}
		}
	}
	return len(cheats)
}

func Part2(m [][]byte, distToEnd map[Pos]int) int {
	height := len(m)
	width := len(m[0])

	cheats := map[Cheat]int{}
	for y, row := range m {
		for x := range row {
			if m[y][x] == '#' {
				continue
			}
			cheatSrc := Pos{x, y}
			for _, cheatDst := range CheatDsts2(cheatSrc, 20) {

				if cheatDst.X < 0 || cheatDst.X >= width || cheatDst.Y < 0 || cheatDst.Y >= height {
					continue
				}

				if m[cheatDst.Y][cheatDst.X] == '#' {
					continue
				}

				if _, isReachable := distToEnd[cheatDst]; !isReachable {
					continue
				}

				cheatLen := TaxicabDist(cheatSrc, cheatDst)
				saves := distToEnd[cheatSrc] - distToEnd[cheatDst] - cheatLen
				if saves >= 100 {
					cheats[Cheat{cheatSrc, cheatDst}] = saves
				}
			}
		}
	}

	freq := map[int]int{}
	for _, v := range cheats {
		freq[v]++
	}
	return len(cheats)
}

func TaxicabDist(p1, p2 Pos) int {
	return mathutil.AbsInt(p1.X-p2.X) + mathutil.AbsInt(p1.Y-p2.Y)
}

func Part12(inputName string) (int, int) {
	m := ReadInput(inputName)
	m = aocutil.PadMap(m, 0, '#')

	_, endPos := FindStartEnd(m)
	distToEnd := Dijkstra(m, endPos)

	return Part1(m, distToEnd), Part2(m, distToEnd)
}

func FindStartEnd(m [][]byte) (start, end Pos) {
	for y, row := range m {
		for x, cell := range row {
			switch cell {
			case 'S':
				start = Pos{x, y}
			case 'E':
				end = Pos{x, y}
			}
		}
	}
	return
}

// 998943 pt 2 too high
// 986545
