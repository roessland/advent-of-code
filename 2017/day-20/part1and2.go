package main

import "fmt"
import "strings"
import "log"
import "strconv"
import "os"
import "bufio"
import "sort"

type Vec struct {
	X, Y, Z int
}

type Particle struct {
	Num  int
	Pos  Vec
	Vel  Vec
	Acc  Vec
	Dead bool
}

func (p Particle) Tick() Particle {
	p.Vel.X += p.Acc.X
	p.Vel.Y += p.Acc.Y
	p.Vel.Z += p.Acc.Z
	p.Pos.X += p.Vel.X
	p.Pos.Y += p.Vel.Y
	p.Pos.Z += p.Vel.Z
	return p
}

func ParseTuple(s string) Vec {
	from := 0
	for s[from] != '<' {
		from++
	}
	from++
	to := len(s) - 1
	for s[to] != '>' {
		to--
	}
	nums := strings.Split(s[from:to], ",")
	x, errX := strconv.Atoi(nums[0])
	y, errY := strconv.Atoi(nums[1])
	z, errZ := strconv.Atoi(nums[2])
	if errX != nil || errY != nil || errZ != nil {
		log.Fatal(errX, errY, errZ)
	}
	return Vec{x, y, z}
}

func ReadInput() []Particle {
	ps := []Particle{}
	scanner := bufio.NewScanner(os.Stdin)
	num := 0
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ", ")
		p := Particle{}
		p.Num = num
		p.Pos = ParseTuple(fields[0])
		p.Vel = ParseTuple(fields[1])
		p.Acc = ParseTuple(fields[2])
		p.Dead = false
		ps = append(ps, p)
		num++
	}
	return ps
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func RemoveColliding(ps []Particle) {
	positions := map[Vec][]int{}

	live := 0
	for i, p := range ps {
		if p.Dead {
			continue
		}
		live++
		if positions[p.Pos] == nil {
			positions[p.Pos] = []int{i}
		} else {
			positions[p.Pos] = append(positions[p.Pos], i)
		}
	}
	for _, is := range positions {
		if len(is) > 1 {
			for _, i := range is {
				ps[i].Dead = true
			}
		}
	}
	fmt.Printf("%d ", live)
}

func main() {
	ps := ReadInput()

	sort.Slice(ps, func(i, j int) bool {
		return Abs(ps[i].Acc.X)+Abs(ps[i].Acc.Y)+Abs(ps[i].Acc.Z) < Abs(ps[j].Acc.X)+Abs(ps[j].Acc.Y)+Abs(ps[j].Acc.Z)
	})

	fmt.Println("Part 1: Particle closest to origin:", ps[0].Num)

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Num < ps[j].Num
	})

	fmt.Println("Part 2: Remaining particles")
	for k := 0; k < 50; k++ {
		RemoveColliding(ps)
		for i, _ := range ps {
			ps[i] = ps[i].Tick()
		}
	}
}
