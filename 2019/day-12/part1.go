package main

import "fmt"

type Vec struct {
	X, Y, Z int
}

func (v Vec) Norm1() int {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Moon struct {
	Pos Vec
	Vel Vec
}

func Sign(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func Energy(moons []Moon) int {
	total := 0
	for _, m := range moons {
		total += m.Pos.Norm1() * m.Vel.Norm1()
	}
	return total
}

func Update(prevMoons []Moon) []Moon {
	moons := make([]Moon, len(prevMoons))
	copy(moons, prevMoons)

	// Apply gravity
	for i, this := range prevMoons {
		for j, other := range prevMoons {
			if i == j {
				continue
			}
			moons[i].Vel.X += Sign(other.Pos.X - this.Pos.X)
			moons[i].Vel.Y += Sign(other.Pos.Y - this.Pos.Y)
			moons[i].Vel.Z += Sign(other.Pos.Z - this.Pos.Z)
		}
	}

	// Apply velocity
	for i, its := range moons {
		moons[i].Pos.X += its.Vel.X
		moons[i].Pos.Y += its.Vel.Y
		moons[i].Pos.Z += its.Vel.Z
	}

	return moons
}

func main() {
	moons := []Moon{
		{Pos: Vec{14, 4, 5}},
		{Pos: Vec{12, 10, 8}},
		{Pos: Vec{1, 7, -10}},
		{Pos: Vec{16, -5, 3}},
	}

	for i := 0; i < 1000; i++ {
		moons = Update(moons)
	}
	fmt.Println(Energy(moons)) // 1691 too low
}
