package main

import (
	"fmt"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"github.com/roessland/gopkg/priorityqueue2"
)

type Map []string

func (m Map) SizeY() int {
	return len(m)
}

func (m Map) SizeX() int {
	return len(m[0])
}

func (m Map) At1(p Pos) byte {
	if p.Y < 0 || p.Y >= m.SizeY() || p.X < 0 || p.X >= m.SizeX() {
		return '#'
	}
	y := (p.Y%m.SizeY() + m.SizeY()) % m.SizeY()
	x := (p.X%m.SizeX() + m.SizeX()) % m.SizeX()
	return m[y][x]
}

func (m Map) At2(p Pos) byte {
	y := (p.Y%m.SizeY() + m.SizeY()) % m.SizeY()
	x := (p.X%m.SizeX() + m.SizeX()) % m.SizeX()
	return m[y][x]
}

func (m Map) FindStartPos() Pos {
	for y, line := range m {
		for x, c := range line {
			if c == 'S' {
				return Pos{Y: y, X: x}
			}
		}
	}
	panic("No start position found")
}

// TileOrigin returns the top-left position of the given tile.
// Origin tile has top-left position (0,0).
func (m Map) TileOrigin(tile Tile) Pos {
	return Pos{tile.Y * m.SizeY(), tile.X * m.SizeX()}
}

func main() {
	m := ReadInput()
	// Part1(m)
	// f, err := os.Create("pprof.txt")
	// if err != nil {
	// 	log.Fatal("could not create CPU profile: ", err)
	// }
	// defer f.Close() // error handling omitted for example
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	log.Fatal("could not start CPU profile: ", err)
	// }
	//
	// defer pprof.StopCPUProfile()
	Part2(m)
}

func Part1(m *Map) {
	start := m.FindStartPos()
	stepsToEach := BFS1(m, start)
	reachable := 0
	for y := 0; y < m.SizeY(); y++ {
		for x := 0; x < m.SizeX(); x++ {
			pos := Pos{y, x}
			steps, ok := stepsToEach[pos]
			if !ok {
				continue
			}
			if steps <= 64 && steps%2 == 0 {
				reachable++
			}
		}
	}
	fmt.Println(reachable)
}

// BFS1 returns the number of steps needed to reach any reachable position.
func BFS1(m *Map, start Pos) map[Pos]int {
	steps := map[Pos]int{}

	queue := priorityqueue2.New[ProblemNode, int]()
	queue.Push(ProblemNode{start}, 0)
	for queue.Len() > 0 {
		currNode, currSteps := queue.PopPri()

		// Already visited
		if _, ok := steps[currNode.Pos]; ok {
			continue
		}

		// Visit and enqueue neighbors
		steps[currNode.Pos] = currSteps
		for _, neighborPos := range currNode.Pos.Neighbors() {
			if m.At1(neighborPos) == '#' {
				continue
			}
			queue.Push(ProblemNode{Pos: neighborPos}, currSteps+1)
		}
	}
	return steps
}

type Quadrant int

const (
	InvalidQuadrant Quadrant = iota
	E
	NE
	N
	NW
	W
	SW
	S
	SE
	Origin
)

var AllQuadrants = []Quadrant{E, NE, N, NW, W, SW, S, SE}

func (q Quadrant) TileSign() Tile {
	switch q {
	case E:
		return Tile{0, 1}
	case NE:
		return Tile{-1, 1}
	case N:
		return Tile{-1, 0}
	case NW:
		return Tile{-1, -1}
	case W:
		return Tile{0, -1}
	case SW:
		return Tile{1, -1}
	case S:
		return Tile{1, 0}
	case SE:
		return Tile{1, 1}
	case Origin:
		return Tile{0, 0}
	default:
		panic("invalid quadrant")
	}
}

type Tile Pos // (0,0) is the start tile, (0,1) is the east tile, etc.

func Sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func (t Tile) Sign() Tile {
	return Tile{Sign(t.Y), Sign(t.X)}
}

// TileInQuadrant returns true if the given tile is in the given quadrant.
// Origin is in all quadrants.
func TileInQuadrant(tile Tile, quadrant Quadrant) bool {
	return tile == Tile{0, 0} || tile.Sign() == quadrant.TileSign()
}

func NumReachableInExactly(m *Map, cachedBFS2 map[Pos]int, S Pos, exactSteps int) int {
	// Pre-compute distances for a small area around the origin. The assumption
	// is that the number of reachable tiles is less predictable near the origin.
	// cachedBFS2 := BFS2(m, S, m.SizeX()*m.SizeY())

	// For each position in small map
	sumTotal := 0
	for y0 := 0; y0 < m.SizeY(); y0++ {
		for x0 := 0; x0 < m.SizeX(); x0++ {
			p0 := Pos{y0, x0}

			// Don't count walls
			if m.At1(p0) == '#' {
				continue
			}
			// Count reachable equivalent positions on entire map
			sumPos := NumReachablePosInExactly(m, cachedBFS2, S, p0, exactSteps)
			sumTotal += sumPos
		}
	}
	return sumTotal
}

// NumReachablePosInExactly returns the number of equivalent positions
// reachable from the given position in exactly the given number of steps.
func NumReachablePosInExactly(m *Map, cachedBFS2 map[Pos]int, S, p0 Pos, exactSteps int) int {
	sumTotal := 0
	for _, quadrant := range AllQuadrants {
		sumQuadrant := NumReachablePosQuadrantInExactly(m, cachedBFS2, S, p0, exactSteps, quadrant)
		sumTotal += sumQuadrant
	}
	// Since all 8 quadrants include the origin tile, they have been double-counted.
	// Subtract 7 of those to compensate.
	s, ok := cachedBFS2[p0]
	if ok && s%2 == exactSteps%2 && s <= exactSteps {
		sumTotal -= 7
	}

	return sumTotal
}

// NumReachablePosQuadrantInExactly returns the number of equivalent positions
// reachable from the given position in exactly the given number of steps, in
// the given quadrant.
func NumReachablePosQuadrantInExactly(m *Map, cachedBFS2 map[Pos]int, S, p0 Pos, exactSteps int, quadrant Quadrant) int {
	if m.At1(p0) == '#' {
		panic("p0 is a wall")
	}
	if _, ok := cachedBFS2[p0]; !ok {
		// Not reachable at all, enclosed
		return 0
	}
	mustBeEven := exactSteps%2 == 0
	mustBeOdd := !mustBeEven

	// Go outwards until the number of steps to reach the positions is predictable.
	numSteps := []int{}
	numDupes := []int{}
	for radius := 0; ; radius++ {
		tile := PickTileInQuadrantAtRadius(quadrant, radius)
		dupes := CountTilesInQuadrantAtRadius(quadrant, radius)
		if tile == nil {
			continue // No tile in that quadrant of that radius
		}
		// Find equivalent position p in the tile
		tileOrigin := m.TileOrigin(*tile)
		p := tileOrigin.Add(p0)

		// Find steps from S to p
		steps, ok := cachedBFS2[p]
		if !ok {
			deltas := []int{}
			for i := 0; i < len(numSteps)-1; i++ {
				deltas = append(deltas, numSteps[i+1]-numSteps[i])
			}
			typ := m.At2(p)
			msg := fmt.Sprintf("could not find steps from %v to %v, deltas=%v, expand radius? type is %c", S, p, deltas, typ)
			panic(msg)
		}

		// Only add if reachable
		if steps%2 == 0 && mustBeEven || steps%2 == 1 && mustBeOdd {
			numSteps = append(numSteps, steps)
			numDupes = append(numDupes, dupes)
		}

		// Break if the number of steps is predictable
		if len(numSteps) <= 5 {
			continue
		}
		end := len(numSteps) - 1
		delta := numSteps[end] - numSteps[end-1]
		if delta == numSteps[end-1]-numSteps[end-2] && delta == numSteps[end-2]-numSteps[end-3] && delta == numSteps[end-3]-numSteps[end-4] && delta == numSteps[end-4]-numSteps[end-5] {
			break
		}
	}

	// fmt.Println("became predictable at radius", len(numSteps)-1)

	// Count even/odd step counts less than or equal to exactSteps.
	reachable := 0

	// First, count those we already computed Remember to multiply by the number
	// of tiles for each radius. For example, p in the origin is counted once,
	// while p in "next radius" exists in two tiles, and should be counted
	// twice.
	//
	//    ┌───┐┌───┐┌───┐
	//    │ s ││   ││   │
	//    └───x└───┘└───┘
	//    ┌───┐┌───┐┌───┐
	//    │   ││   ││   │
	//    └───┘└───┘└───x
	//    ┌───┐┌───┐┌───┐
	//    │   ││   ││   │
	//    └───┘└───x└───┘
	//
	for i, steps := range numSteps {
		if steps > exactSteps {
			break
		}
		isEven := steps%2 == 0
		isOdd := !isEven

		if isEven && mustBeEven {
			reachable += numDupes[i]
		} else if isOdd && mustBeOdd {
			reachable += numDupes[i]
		}
	}

	// Then, count the matches in the infinite series {a0 + i*da*(r0+dr*i)}, i in integers [0, inf)
	da := numSteps[len(numSteps)-1] - numSteps[len(numSteps)-2]
	a0 := numSteps[len(numSteps)-1] + da
	dr := numDupes[len(numDupes)-1] - numDupes[len(numDupes)-2]
	r0 := numDupes[len(numDupes)-1] + dr
	reachableInfinite := SumSeries(a0, da, r0, dr, exactSteps)
	// fmt.Println("reachable steps are", numSteps)
	// fmt.Println("reachableInfinite", reachableInfinite)
	reachable += reachableInfinite

	//

	return reachable
}

func SumSeries(a0, da, r0, dr, max int) int {
	if a0 > max {
		return 0
	}
	// S0 = {a0 + i*da*(r0+dr*i)}, i in integers [0, inf)
	//
	// S1 = {
	//         S0 where a0 + i*da is odd, if isOdd
	//         S0 where a0 + i*da is even, if !isOdd
	//      }
	//
	// S2 = S1 where a0 + i*da is <= max
	//
	// N = CountSeriesEvenOdd(a0, da, max, isOdd)
	// Sum = Sum from i=1 to N: i*(r0+dr*i)
	//
	// Sum { a0 + i*da is even ? 1 : 0 } * (r0 + dr*i)
	//
	// 1 + 2 + ... + N = N*(N+1)/2
	//
	// dr + 2dr + ... + Ndr = dr*(N*(N+1)/2)
	// => (r0) + (r0+dr) + (r0+2dr) + ... + (r0+Ndr) = (N+1)r0 + dr*(N*(N+1)/2)
	//
	// Steps:
	// a0, a0+da, a0+2da, a0+3da, ..., a0+ida
	//
	// What is N?
	// a0 + i*da <= max
	// => i*da <= max - a0
	// => i <= (max - a0)/da
	// => N = floor((max - a0)/da)
	N := (max - a0) / da
	ret := (N+1)*r0 + dr*(N*(N+1)/2)
	if ret < 0 {
		msg := fmt.Sprintf("ret < 0: %d, %d, %d, %d, %d", a0, da, r0, dr, max)
		panic(msg)
	}
	return ret
}

// CountSeriesEven counts the number of even numbers in the series <= max.
// {a0 + i*da}, i in integers [0, inf).
// 1, 3, 5, 7, 9
// 10-1 = 9
func CountSeriesEven(a0, da, max int) int {
	if da == 0 {
		panic("da == 0")
	}
	if max < a0 {
		return 0
	}

	// a0 + I*da <= max < a0 + (I+1)*da
	I := (max - a0) / da
	addRem := 0
	if a0+I*da == max {
		addRem++
	}
	_ = addRem

	// (2, 2, 10, false)
	// 2, 4, 6, 8, 10
	// All are even
	if a0%2 == 0 && da%2 == 0 {
		return I + 1
	}

	// (1, 2, 10, false)
	// 1, 3, 5, 7, 9
	// None are even
	if a0%2 == 1 && da%2 == 0 {
		return 0
	}

	// (2, 3, 10, false)
	// 2, 5, 8
	// Half are even
	if a0%2 == 0 && da%2 == 1 {
		return (I / 2) + 1
	}

	// (1, 3, 10, false)
	// 1, 4, 7, 10
	// Half are even
	if a0%2 == 1 && da%2 == 1 {
		return (I + 1) / 2
	}

	msg := fmt.Sprintf("unhandled case: %d, %d, %d", a0, da, max)
	panic(msg)
}

func CountSeriesEvenOdd(a0, da, max int, odd bool) int {
	if odd {
		// Count 1 3 5 7 9 where odd and <= 10 is equivalent to
		// Count 0 2 4 6 8 where even and <= 9
		return CountSeriesEven(a0-1, da, max-1)
	}
	return CountSeriesEven(a0, da, max)
}

// PickTileInQuadrantAtRadius returns a tile in the given quadrant, at the given radius.
//
// Tiles are scanned in this order, and the first matching one is returned:
//
//	             ╱     ╲
//	            ╱       ╲
//	1  ─────────────────────────▶
//	          ╱           ╲
//	2  ─────────────────────────▶
//	        ╱               ╲
//	3  ─────────────────────────▶
func PickTileInQuadrantAtRadius(quadrant Quadrant, radius int) *Tile {
	for Dy := -radius; Dy <= radius; Dy++ {
		for Dx := -radius; Dx <= radius; Dx++ {
			tile := Tile{Dy, Dx}
			// Not at radius
			if Abs(Dy)+Abs(Dx) != radius {
				continue
			}
			if !TileInQuadrant(tile, quadrant) {
				continue
			}
			return &Tile{Dy, Dx}
		}
	}
	return nil
}

// CountTilesInQuadrantAtRadius returns the number of tiles in the given quadrant,
// at the given radius. Assumes the map is square.
// For example, there are 1 tile in the SE quadrant with radius 2,
// and 2 tiles with radius 3.
// ┌───┐┌───┐┌───┐
// │ s ││   ││   │
// └───┘└───┘└───┘
// ┌───┐┌───┐┌───┐
// │   ││   ││   │
// └───┘└───2└───3
// ┌───┐┌───┐┌───┐
// │   ││   ││   │
// └───┘└───3└───4
func CountTilesInQuadrantAtRadius(quadrant Quadrant, radius int) int {
	q := quadrant.TileSign()
	if q.Y == 0 || q.X == 0 {
		return 1
	}
	if radius <= 1 {
		return 1
	}
	return radius - 1
}

func Part2(m Map) {
	fmt.Println("Map size:", m.SizeY(), "x", m.SizeX())
	maxSteps := 26501365
	start := m.FindStartPos()
	cachedBFS2 := BFS2(&m, start, 15*m.SizeX())
	num := NumReachableInExactly(&m, cachedBFS2, start, maxSteps)
	fmt.Println("Part 2:", num)
}

// Dist is the distance between the center of two tiles.
func TileDist(m Map, Dy, Dx int) int {
	if Dy < 0 {
		Dy = -Dy
	}
	if Dx < 0 {
		Dx = -Dx
	}
	// return Dy*m.SizeY() + Dx*m.SizeX()
	return Dy*131 + Dx*131
}

func TileReachable(m Map, Dy, Dx int, maxSteps int, sizeY, sizeX int) (int, Pos) {
	startPos0 := Pos{65, 65}
	tileDist := (Abs(Dy) + Abs(Dx)) * 131
	toEast := 65
	toWest := 65
	toNorth := 65
	toSouth := 65

	var startPos Pos

	// Determine position on that edge
	var distToStartPos int
	if Dy == 0 && Dx == 0 {
		// ...
		// .X.
		// ...
		startPos = startPos0
	} else if Abs(Dy) == Abs(Dx) && Dy < 0 && Dx > 0 {
		// ..X
		// .S.
		// ...
		startPos = Pos{sizeY - 1, 0} // bottom left
		distToStartPos = tileDist - toEast - toNorth
	} else if Abs(Dy) == Abs(Dx) && Dy < 0 && Dx < 0 {
		// X..
		// .S.
		// ...
		startPos = Pos{sizeY - 1, sizeX - 1} // bottom right
		distToStartPos = tileDist - toEast - toNorth
	} else if Abs(Dy) == Abs(Dx) && Dy > 0 && Dx < 0 {
		// ...
		// .S.
		// X..
		startPos = Pos{0, sizeX - 1} // top right
		distToStartPos = tileDist - toEast - toNorth
	} else if Abs(Dy) == Abs(Dx) && Dy > 0 && Dx > 0 {
		// ...
		// .S.
		// ..X
		startPos = Pos{0, 0} // top left
		distToStartPos = tileDist - toEast - toNorth
	} else if Dy > 0 && Dx > 0 {
		// ...
		// .S.
		// ..X
		startPos = Pos{0, 0} // top left
		distToStartPos = tileDist - toEast - toSouth
	} else if Dy < 0 && Dx > 0 {
		// ..X
		// .S.
		// ...
		startPos = Pos{sizeY - 1, 0} // bottom left
		distToStartPos = tileDist - toEast - toNorth
	} else if Dy < 0 && Dx < 0 {
		// X..
		// .S.
		// ...
		startPos = Pos{sizeY - 1, sizeX - 1} // bottom right
		distToStartPos = tileDist - toWest - toNorth
	} else if Dy > 0 && Dx < 0 {
		// ...
		// .S.
		// X..
		startPos = Pos{0, sizeX - 1} // top right
		distToStartPos = tileDist - toWest - toSouth
	} else if Dy == 0 && Dx > 0 {
		// ...
		// .SX
		// ...
		startPos = Pos{startPos0.Y, 0} // middle left
		distToStartPos = tileDist - toEast
	} else if Dy == 0 && Dx < 0 {
		// ...
		// XS.
		// ...
		startPos = Pos{startPos0.Y, sizeX - 1} // middle right
		distToStartPos = tileDist - toWest
	} else if Dy < 0 && Dx == 0 {
		// .X.
		// .S.
		// ...
		startPos = Pos{sizeY - 1, startPos0.X} // bottom middle
		distToStartPos = tileDist - toNorth
	} else if Dy > 0 && Dx == 0 {
		// ...
		// .S.
		// .X.
		startPos = Pos{0, startPos0.X} // top middle
		distToStartPos = tileDist - toSouth
	} else {
		panic(fmt.Sprint("YOLO: ", Dy, Dx))
	}

	maxSteps -= distToStartPos

	if maxSteps < 0 {
		return 0, Pos{0, 0}
	}

	reachableEven, reachableOdd := CachedBFS1(&m, startPos, maxSteps)

	if distToStartPos%2 == 0 {
		return reachableEven, startPos
	} else {
		return reachableOdd, startPos
	}
}

type CacheKey struct {
	Start    Pos
	MaxSteps int
}

type CacheVal struct {
	Even, Odd int
}

var (
	cache    = []CacheVal{}
	bfsCache = map[Pos]map[Pos]int{}
)

func init() {
	for i := 0; i < 100000; i++ {
		cache = append(cache, CacheVal{-1, -1})
	}
}

func getCacheKey(start Pos, maxSteps int) int {
	y, x := start.Y, start.X
	if y == 65 && x == 65 {
		return 1000 + maxSteps
	} else if y == 0 && x == 65 {
		return 2000 + maxSteps
	} else if y == 65 && x == 0 {
		return 3000 + maxSteps
	} else if y == 65 && x == 130 {
		return 4000 + maxSteps
	} else if y == 130 && x == 65 {
		return 5000 + maxSteps
	} else if y == 0 && x == 0 {
		return 6000 + maxSteps
	} else if y == 0 && x == 130 {
		return 7000 + maxSteps
	} else if y == 130 && x == 0 {
		return 8000 + maxSteps
	} else if y == 130 && x == 130 {
		return 9000 + maxSteps
	} else {
		panic("unknown start pos")
	}
}

func CachedBFS1(m *Map, start Pos, maxSteps int) (int, int) {
	if maxSteps > 1000 {
		maxSteps %= 1000
	}
	cacheKey := getCacheKey(start, maxSteps)
	cachedVal := cache[cacheKey]
	if cachedVal.Even != -1 {
		return cachedVal.Even, cachedVal.Odd
	}

	stepsToReach, ok := bfsCache[start]
	if !ok {
		stepsToReach = BFS1(m, start)
		bfsCache[start] = stepsToReach
	}
	reachableEven := 0
	reachableOdd := 0
	for y := 0; y < 131; y++ {
		for x := 0; x < 131; x++ {
			pos := Pos{y, x}
			steps, ok := stepsToReach[pos]
			if !ok {
				continue
			}
			if steps <= maxSteps && steps%2 == 0 {
				reachableEven++
			}
			if steps <= maxSteps && steps%2 == 1 {
				reachableOdd++
			}
		}
	}
	cache[cacheKey] = CacheVal{reachableEven, reachableOdd}
	return reachableEven, reachableOdd
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Abs2(x, y int) int {
	if x < 0 && y < 0 {
		return -x - y
	}
	if x < 0 {
		return -x + y
	}
	if y < 0 {
		return x - y
	}
	return x + y
}

// BFS2 returns the number of steps needed to reach any reachable position, reachable within maxSteps steps.
func BFS2(m *Map, start Pos, maxSteps int) map[Pos]int {
	steps := map[Pos]int{}

	queue := priorityqueue2.New[ProblemNode, int]()
	queue.Push(ProblemNode{start}, 0)
	for queue.Len() > 0 {
		currNode, currSteps := queue.PopPri()

		// Already visited,
		if _, ok := steps[currNode.Pos]; ok {
			continue
		}

		// Too far away
		if currSteps > maxSteps {
			continue
		}

		// Visit and enqueue neighbors
		steps[currNode.Pos] = currSteps
		for _, neighborPos := range currNode.Pos.Neighbors() {
			if m.At2(neighborPos) == '#' {
				continue
			}
			queue.Push(ProblemNode{Pos: neighborPos}, currSteps+1)
		}
	}
	return steps
}

type Pos struct {
	Y, X int
}

func (p Pos) Add(other Pos) Pos {
	return Pos{p.Y + other.Y, p.X + other.X}
}

func (p Pos) North() Pos {
	return Pos{p.Y - 1, p.X}
}

func (p Pos) South() Pos {
	return Pos{p.Y + 1, p.X}
}

func (p Pos) East() Pos {
	return Pos{p.Y, p.X + 1}
}

func (p Pos) West() Pos {
	return Pos{p.Y, p.X - 1}
}

func (p Pos) Neighbors() []Pos {
	return []Pos{p.North(), p.South(), p.East(), p.West()}
}

type ProblemNode struct {
	Pos Pos
}

func ReadInput() Map {
	return aocutil.ReadLines("input.txt")
}

// 515308687619374 too low
// 592712210123430 too low
// 592723929260582 ?
