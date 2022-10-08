package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Vec struct {
	X, Y int
}

type State struct {
	Pos, Vel Vec
}

func IsInTargetArea(p Vec) bool {
	return 137 <= p.X && p.X <= 171 && -98 <= p.Y && p.Y <= -73
}

func CouldEndUpInTargetArea(p Vec) bool {
	return p.X <= 171 && p.Y >= -98
}

//func IsInTargetArea(p Vec) bool {
//	return 20 <= p.X && p.X <= 30 && -10 <= p.Y && p.Y <= -5
//}
//
//func CouldEndUpInTargetArea(p Vec) bool {
//	return p.X <= 30 && p.Y >= -10
//}

func (s State) Next() State {
	s.Pos.X += s.Vel.X
	s.Pos.Y += s.Vel.Y
	if s.Vel.X > 0 {
		s.Vel.X--
	}
	s.Vel.Y--
	return s
}

func (s State) HitsTarget() bool {
	if IsInTargetArea(s.Pos) {
		return true
	}
	if !CouldEndUpInTargetArea(s.Pos) {
		return false
	}
	return s.Next().HitsTarget()
}

func (s State) HighestPoint() int {
	if CouldEndUpInTargetArea(s.Next().Pos) {
		return Max(s.Pos.Y, s.Next().HighestPoint())
	} else {
		return 0
	}
}

func main() {
	uniq := map[Vec]bool{}
	maxHeight := -1
	for vx := 1; vx < 172; vx++ {
		for vy := -98; vy < 2000; vy++ {
			v := Vec{vx, vy}
			s := State{Vec{0, 0}, v}
			if s.HitsTarget() {
				maxHeight = Max(maxHeight, s.HighestPoint())
				fmt.Println(maxHeight)
				uniq[v] = true
			}
		}
	}
	fmt.Println(maxHeight)
	fmt.Println(len(uniq)) // 141 too low
}
