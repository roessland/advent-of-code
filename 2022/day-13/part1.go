package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	Children []Node
	IsLeaf   bool
	Num      int
}

func (n Node) String() string {
	if n.IsLeaf {
		return fmt.Sprintf("%d", n.Num)
	}
	parts := []string{}
	for _, child := range n.Children {
		parts = append(parts, child.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ","))
}

func Split(s string) []string {
	parts := []string{}
	if len(s) == 0 {
		return parts
	}

	prev := 0
	depth := 0
	for i, c := range s {
		if c == '[' {
			depth++
		} else if c == ']' {
			depth--
		} else if depth == 0 && c == ',' {
			parts = append(parts, s[prev:i])
			prev = i + 1
		}
	}
	parts = append(parts, s[prev:])
	return parts
}

func Read(s string) Node {
	if len(s) == 0 {
		panic("empty string")
	}
	if s[0] == '[' {
		children := []Node{}
		for _, part := range Split(s[1 : len(s)-1]) {
			children = append(children, Read(part))
		}
		return Node{Children: children, IsLeaf: false}
	} else if len(s) >= 2 && s[0] == '1' && s[1] == '0' {
		return Node{
			IsLeaf: true,
			Num:    10,
		}
	} else if s[0] >= '0' && s[0] <= '9' {
		return Node{
			Num:    int(s[0] - '0'),
			IsLeaf: true,
		}
	} else {
		panic(fmt.Errorf("unexpected input: %s", s))
	}
}

func Compare(left, right Node) int {
	if left.IsLeaf && right.IsLeaf {
		if left.Num < right.Num {
			return -1
		} else if left.Num > right.Num {
			return 1
		} else {
			return 0
		}
	} else if !left.IsLeaf && !right.IsLeaf {
		if len(left.Children) == 0 && len(right.Children) > 0 {
			return -1
		} else if len(left.Children) > 0 && len(right.Children) == 0 {
			return 1
		} else if len(left.Children) == 0 && len(right.Children) == 0 {
			return 0
		} else {
			c := Compare(left.Children[0], right.Children[0])
			if c != 0 {
				return c
			}
			return Compare(Node{Children: left.Children[1:]}, Node{Children: right.Children[1:]})
		}
	} else if left.IsLeaf && !right.IsLeaf {
		return Compare(Node{Children: []Node{left}}, right)
	} else if !left.IsLeaf && right.IsLeaf {
		return Compare(left, Node{Children: []Node{right}})
	} else {
		panic("unreachable")
	}
}

func ReadInput() []Node {
	sum := 0
	for i, pairs := range strings.Split(input, "\n\n") {
		pairs := strings.Split(pairs, "\n")
		left := Read(pairs[0])
		right := Read(pairs[1])
		if Compare(left, right) == -1 {
			index := i + 1
			sum += index
		}
	}
	fmt.Println(sum)
	return nil
}

func main() {
	ReadInput()
}
