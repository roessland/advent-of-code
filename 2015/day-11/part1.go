package main

import "fmt"

type Password struct {
	chars [8]int
}

func (p *Password) next() {
	i := len(p.chars) - 1
	for {
		p.chars[i]++
		if p.chars[i] >= 26 {
			p.chars[i] = 0
			i--
		} else {
			break
		}
	}
}

func NewPassword(pass string) Password {
	p := Password{}
	for i, c := range pass {
		p.chars[i] = int(c - 'a')
	}
	return p
}

func (p *Password) ToString() string {
	pBytes := make([]byte, 8)
	for i, num := range p.chars {
		pBytes[i] = byte(num + 'a')
	}
	return string(pBytes)
}

func (p *Password) Valid() bool {
	return p.HasIncreasingStraight() &&
		p.NoForbiddenChars() &&
		p.HasTwoPairs()
}

func (p *Password) HasIncreasingStraight() bool {
	for i := 0; i < len(p.chars)-3; i++ {
		if p.chars[i] == p.chars[i+1]-1 &&
			p.chars[i] == p.chars[i+2]-2 {
			return true
		}
	}
	return false
}

func (p *Password) NoForbiddenChars() bool {
	forbidden := map[int]bool{
		'i' - 'a': true,
		'o' - 'a': true,
		'l' - 'a': true,
	}
	for _, num := range p.chars {
		if forbidden[num] {
			return false
		}
	}
	return true
}

func (p Password) HasTwoPairs() bool {
	busy := make(map[int]bool)
	pairs := 0
	for i := 0; i <= len(p.chars)-2; i++ {
		if p.chars[i] == p.chars[i+1] && !busy[i] {
			busy[i], busy[i+1] = true, true
			pairs++
			if pairs >= 2 {
				return true
			}
		}
	}
	return false
}

func main() {
	p := NewPassword("vzbxkghb")
	for !p.Valid() {
		p.next()
	}
	fmt.Printf("Next password: %v\n", p.ToString())
	p.next()
	for !p.Valid() {
		p.next()
	}
	fmt.Printf("Next password: %v\n", p.ToString())
}
