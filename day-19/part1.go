package main

import "regexp"
import "fmt"
import "io/ioutil"
import "strings"

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

func main() {
	rep, mol0 := ReadInput("input.txt")
	possibleMols := make(map[string]bool)
	atoms0 := regexp.MustCompile(`[A-Z][a-z]*`).FindAllString(mol0, -1)
	for i, atom0 := range atoms0 {
		originalAtom0 := atom0
		for _, to := range rep[originalAtom0] {
			atoms0[i] = to
			possibleMols[strings.Join(atoms0, "")] = true
		}
		atoms0[i] = originalAtom0
	}
	fmt.Printf("%v\n", len(possibleMols))
}
