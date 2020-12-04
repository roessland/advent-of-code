package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pass struct {
	Cid string
	Hcl string
	Hgt string
	Iyr string
	Byr string
	Eyr string
	Pid string
	Ecl string
}

func (p *Pass) AddInfos(infos string) {
	infos = strings.ReplaceAll(infos, "\n", " ")
	parts := strings.Split(infos, " ")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		keyVal := strings.Split(part, ":")
		p.AddInfo(keyVal[0], strings.Trim(keyVal[1], " "))
	}
	return
}

func (p *Pass) AddInfo(key, val string) {
	fmt.Printf(`"%s" "%s"`+"\n", key, val)
	switch key {
	case "iyr":
		p.Iyr = val
	case "cid":
		p.Cid = val
	case "eyr":
		p.Eyr = val
	case "hcl":
		p.Hcl = val
	case "hgt":
		p.Hgt = val
	case "ecl":
		p.Ecl = val
	case "byr":
		p.Byr = val
	case "pid":
		p.Pid = val
	default:
		log.Fatal("unknown key " + key)
	}
}

func (p *Pass) Valid() bool {
	if len(p.Iyr) == 0 {
		return false
	}
	if len(p.Cid) == 0 {
		// return false
	}
	if len(p.Eyr) == 0 {
		return false
	}
	if len(p.Hcl) == 0 {
		return false
	}
	if len(p.Hgt) == 0 {
		return false
	}
	if len(p.Ecl) == 0 {
		return false
	}
	if len(p.Byr) == 0 {
		return false
	}
	if len(p.Pid) == 0 {
		return false
	}

	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var ps []Pass
	var p Pass
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ps = append(ps, p)
			p = Pass{}
		}
		p.AddInfos(line)
	}
	ps = append(ps, p) // the final one

	numValid := 0
	for _, p := range ps {
		if p.Valid() {
			numValid++
		}
	}
	fmt.Println(numValid) // 238 too low
}