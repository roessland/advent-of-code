package main

import "fmt"
import "errors"

const N int = 3001330

type Node struct {
	Number int
	Next   *Node
	Items  int
}

func NewCircle() *Node {
	last := &Node{N, nil, 1}
	next := last
	for i := N - 1; 1 <= i; i-- {
		next = &Node{i, next, 1}
	}
	last.Next = next
	return next
}

func (n *Node) Print() {
	first := n
	for {
		fmt.Printf("(%d has %d), ", n.Number, n.Items)
		n = n.Next
		if n == first {
			break
		}
	}
	fmt.Println()
}

func (n *Node) Take() (*Node, error) {
	if n == n.Next {
		return n, errors.New("no-op")
	}
	n.Items += n.Next.Items
	n.Next = n.Next.Next
	return n.Next, nil
}

func main() {
	circ := NewCircle()

	var err error
	for {
		circ, err = circ.Take()
		if err != nil {
			circ.Print()
			break
		}
	}
}
