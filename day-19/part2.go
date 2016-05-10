package main

import "io/ioutil"
import "strings"
import "math/rand"
import "fmt"

type Rep [2]string

func ReadInput(filename string) ([]Rep, string) {
	buf, _ := ioutil.ReadFile(filename)
	reps := []Rep{}
	var mol string
	lastLine := false
	for _, line := range strings.Split(string(buf), "\n") {
		if lastLine {
			mol = line
			break
		}
		if len(line) == 0 {
			lastLine = true
			continue
		}
		fromto := strings.Split(line, " => ")
		from, to := fromto[0], fromto[1]
		reps = append(reps, Rep{from, to})
	}
	return reps, mol
}

func main() {
	reps, mol0 := ReadInput("input.txt")
	mol := mol0

	count := 0
	for mol != "e" {
		molBefore := mol
		// Shuffle replacements
		for i := range reps {
			j := rand.Intn(i + 1)
			reps[i], reps[j] = reps[j], reps[i]
		}

		// Apply all possible replacements
		for _, rep := range reps {
			if strings.Index(mol, rep[1]) == -1 {
				continue
			}
			mol = strings.Replace(mol, rep[1], rep[0], 1)
			count++
		}
		if mol == molBefore {
			// No replacements could be made, restart with original molecule.
			mol = mol0
			count = 0
		}
	}
	fmt.Printf("Depth: %v\n", count)
}
