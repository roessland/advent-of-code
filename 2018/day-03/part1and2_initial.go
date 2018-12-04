package main

import "fmt"
import "strconv"
import "log"
import "bufio"
import "os"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "github.com/roessland/gopkg/disjointset"

type Vec struct {
	X, Y int
}

type Claim struct {
	X, Y, W, H int
}

func main() {
	sheet := map[Vec]int{}
	owners := map[Vec][]string{}
	beenstomped := map[string]bool{}
	size := map[string]int{}
	pieces := 0

	//claims := []Claim{}
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		f := strings.Split(line, " @ ")
		owner := strings.Trim(f[0], "# ")
		locs := strings.Split(f[1], ",")
		fmt.Println(locs)
		y := atoi(strings.Trim(strings.Split(locs[1], ": ")[0], ": "))
		x := atoi(locs[0])
		wh := strings.Split(strings.Split(f[1], ": ")[1], "x")
		w := atoi(wh[0])
		h := atoi(wh[1])
		size[owner] = w * h
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				sheet[Vec{i, j}]++
				if sheet[Vec{i, j}] == 2 {
					pieces++
				}
				owners[Vec{i, j}] = append(owners[Vec{i, j}], owner)
				if sheet[Vec{i, j}] >= 2 {
					for _, owner := range owners[Vec{i, j}] {
						beenstomped[owner] = true
					}
				}
			}
		}
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			v := Vec{i, j}
			if owners[v] == nil {
				continue
			}
			if len(owners[v]) > 1 {
			} else {
				if !beenstomped[owners[v][0]] {
					fmt.Printf("%s,", owners[v][0])
				}
			}
		}
	}
	fmt.Println("\npart1:", pieces)
}

var _ = fmt.Println
var _ = strconv.Atoi
var _ = log.Fatal
var _ = bufio.NewScanner // (os.Stdin) -> scanner.Scan(), scanner.Text()
var _ = os.Stdin
var _ = strings.Split    // "str str" -> []string{"str", "str"}
var _ = ioutil.ReadFile  // ("input.txt") -> (buf, err)
var _ = csv.NewReader    // (os.Stdin)
var _ = disjointset.Make // (10) -> ds. ds.Union(a,b), ds.Connected(a,b), ds.Count

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
