package main

import "fmt"
import "log"
import "io/ioutil"
import "strings"
import "os"
import "bytes"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Pos struct {
	X, Y, Z int
}

func (p Pos) Norm() int {
	return (Abs(p.X) + Abs(p.Y) + Abs(p.Z)) / 2
}

func (p Pos) Move(dir string) Pos {
	switch dir {
	case "n":
		p.X++
		p.Z++
	case "ne":
		p.X++
		p.Y--
	case "se":
		p.Y--
		p.Z--
	case "s":
		p.X--
		p.Z--
	case "sw":
		p.X--
		p.Y++
	case "nw":
		p.Y++
		p.Z++
	default:
		log.Fatal("Unknown direction:", dir)
	}
	return p
}

func main() {
	str, _ := ioutil.ReadAll(os.Stdin)
	str = bytes.TrimRight(str, "\n") // trailing newline
	dirs := strings.Split(string(str), ",")
	p := Pos{0, 0, 0}
	maxNorm := 0
	for _, dir := range dirs {
		p = p.Move(dir)
		if p.Norm() > maxNorm {
			maxNorm = p.Norm()
		}
	}
	fmt.Println("Part 1:", p.Norm())
	fmt.Println("Part 2:", maxNorm)
}
