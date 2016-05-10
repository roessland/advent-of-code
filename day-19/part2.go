package main

import "regexp"
import "fmt"
import "io/ioutil"
import "strings"
import "os"

func ReadInput(filename string) (map[string][]string, string) {
	buf, _ := ioutil.ReadFile(filename)
	rep := make(map[string][]string)
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
		if _, ok := rep[from]; !ok {
			rep[from] = []string{}
		}
		rep[from] = append(rep[from], to)
	}
	return rep, mol
}

func PossibleResults(mol0 string, rep map[string][]string) map[string]bool {
	possibleMols := make(map[string]bool)
	var atoms0 []string
	if mol0 == "e" {
		atoms0 = []string{"e"}
	} else {
		atoms0 = regexp.MustCompile(`[A-Z][a-z]*`).FindAllString(mol0, -1)
	}
	for i, atom0 := range atoms0 {
		originalAtom0 := atom0
		for _, to := range rep[originalAtom0] {
			atoms0[i] = to
			possibleMols[strings.Join(atoms0, "")] = true
		}
		atoms0[i] = originalAtom0
	}
	return possibleMols
}

func main() {
	rep, target := ReadInput("input.txt")
	steps := make([]map[string]bool, 1)
	steps[0] = map[string]bool{"e": true}
	for step := 1; step < 600; step++ {
		steps = append(steps, map[string]bool{})
		for mol, _ := range steps[step-1] {
			for res, _ := range PossibleResults(mol, rep) {
				steps[step][res] = true
				if res == target {
					fmt.Printf("Found it during step %v\n", step)
					os.Exit(0)
				}
			}
		}
	}
	//fmt.Printf("%v\n", steps)
}
