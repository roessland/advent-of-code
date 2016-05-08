package main

import "fmt"
import "strconv"
import "io/ioutil"
import "strings"
import "math/rand"

const N int = 8 + 1

func TotalGain(names []string, gain map[string]map[string]int) int {
	var order []int = rand.Perm(N)
	totalGain := 0
	for i, _ := range order {
		person := names[order[i]]
		left := names[order[(i-1+N)%N]]
		right := names[order[(i+1+N)%N]]
		totalGain += gain[person][left] + gain[person][right]
	}
	return totalGain
}

func main() {
	gain := map[string]map[string]int{}
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			break
		}
		words := strings.Split(line, " ")
		name1 := words[0]
		name2 := words[10][:len(words[10])-1]
		sign := "+"
		if words[2] == "lose" {
			sign = "-"
		}
		gain1, _ := strconv.Atoi(sign + words[3])
		if _, ok := gain[name1]; !ok {
			gain[name1] = map[string]int{}
		}
		gain[name1][name2] = gain1

	}
	names := []string{}
	for name, _ := range gain {
		names = append(names, name)
	}

	// Add myself
	names = append(names, "Myself")
	gain["Myself"] = map[string]int{}
	for _, name := range names {
		gain["Myself"][name] = 0
		gain[name]["Myself"] = 0
	}

	// Bogosort
	maxgain := 0
	for i := 0; i < 1000000; i++ {
		g := TotalGain(names, gain)
		if g > maxgain {
			maxgain = g
			fmt.Printf("Max: %v\n", maxgain)
		}
	}
}
