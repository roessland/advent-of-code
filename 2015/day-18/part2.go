package main

import "io/ioutil"
import "strings"
import "fmt"

//import "time"

const N int = 100

type Lights struct {
	status [N][N]bool
}

func (ls *Lights) Print() {
	fmt.Printf("\n")
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			switch st := ls.status[j][i]; st {
			case true:
				fmt.Printf("#")
			case false:
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (ls *Lights) At(j, i int) bool {
	if j < 0 || i < 0 || N <= j || N <= i {
		return false
	}
	return ls.status[j][i]
}

func (ls *Lights) NeighborsOnCount(j, i int) int {
	n := []bool{
		ls.At(j, i+1),
		ls.At(j-1, i+1),
		ls.At(j-1, i),
		ls.At(j-1, i-1),
		ls.At(j, i-1),
		ls.At(j+1, i-1),
		ls.At(j+1, i),
		ls.At(j+1, i+1),
	}
	c := 0
	for _, st := range n {
		if st {
			c++
		}
	}
	return c
}

func (ls *Lights) Update(to *Lights) {
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			st := ls.At(j, i)
			c := ls.NeighborsOnCount(j, i)
			if st {
				if c == 2 || c == 3 {
					to.status[j][i] = true
				} else {
					to.status[j][i] = false
				}
			}
			if !st {
				if c == 3 {
					to.status[j][i] = true
				} else {
					to.status[j][i] = false
				}
			}

		}
	}
	to.status[0][0] = true
	to.status[0][N-1] = true
	to.status[N-1][0] = true
	to.status[N-1][N-1] = true
}

func (ls *Lights) Count() int {
	c := 0
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			if ls.At(j, i) {
				c++
			}
		}
	}
	return c
}

func main() {
	lights, tmp := Lights{}, Lights{}
	buf, _ := ioutil.ReadFile("input.txt")
	for j, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			break
		}
		for i := 0; i < N; i++ {
			switch ch := line[i]; ch {
			case '#':
				lights.status[j][i] = true
			case '.':
				lights.status[j][i] = false
			default:
				fmt.Printf("Unknown character: %v\n", ch)
			}
		}
	}
	lights.status[0][0] = true
	lights.status[0][N-1] = true
	lights.status[N-1][0] = true
	lights.status[N-1][N-1] = true

	for i := 0; i < 100; i++ {
		lights.Update(&tmp)
		tmp, lights = lights, tmp
	}
	fmt.Printf("Count: %v\n", lights.Count())

}
