package main

import "fmt"
import "os"
import "strconv"
import "encoding/csv"
import "log"

type Multiset map[Comp]int

func NewMultiset(comps []Comp) Multiset {
	ms := make(Multiset)
	for _, comp := range comps {
		ms[comp]++
	}
	return ms
}

func (ms Multiset) Clone() Multiset {
	clone := make(Multiset)
	for comp, count := range ms {
		clone[comp] = count
	}
	return clone
}

func (ms Multiset) Without(comp Comp) Multiset {
	without := ms.Clone()
	if without[comp] > 0 {
		without[comp]--
		if without[comp] == 0 {
			delete(without, comp)
		}
		return without
	}
	if without[comp.Rotate()] > 0 {
		without[comp.Rotate()]--
		if without[comp.Rotate()] == 0 {
			delete(without, comp.Rotate())
		}
		return without
	}
	log.Fatal("Can't have less than zero...")
	return nil
}

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

type Comp struct {
	A, B int
}

func (comp Comp) Rotate() Comp {
	return Comp{comp.B, comp.A}
}

func ReadInput() []Comp {
	comps := []Comp{}
	reader := csv.NewReader(os.Stdin)
	reader.Comma = '/'
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		comps = append(comps, Comp{Atoi(record[0]), Atoi(record[1])})
	}
	return comps
}

type Node struct {
	A, B     int
	Children []*Node
}

func NewNode(A, B int) *Node {
	n := Node{A, B, make([]*Node, 0)}
	return &n
}

func BuildTree(parent *Node, remaining Multiset) {
	for comp, _ := range remaining {
		if parent.B == comp.A {
			node := NewNode(comp.A, comp.B)
			parent.Children = append(parent.Children, node)
			BuildTree(node, remaining.Without(comp))
		} else if parent.B == comp.B {
			node := NewNode(comp.B, comp.A)
			parent.Children = append(parent.Children, node)
			BuildTree(node, remaining.Without(comp))
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxLength(root *Node, depth int) (int, int) {
	childMaxLength := 0
	childMaxLengthStrength := 0
	for _, child := range root.Children {
		childLength, childLengthStrength := MaxLength(child, depth+1)
		if childLength == childMaxLength {
			if childLengthStrength > childMaxLengthStrength {
				childMaxLengthStrength = childLengthStrength
			}
		}
		if childLength > childMaxLength {
			childMaxLength = childLength
			childMaxLengthStrength = childLengthStrength
		}
	}
	return 1 + childMaxLength, root.A + root.B + childMaxLengthStrength
}

func main() {
	comps := ReadInput()
	remaining := NewMultiset(comps)
	root := NewNode(0, 0)
	BuildTree(root, remaining)
	lengthPlusOne, itsStrength := MaxLength(root, 0)
	fmt.Println("Length:", lengthPlusOne-1, "Strength:", itsStrength)
}
