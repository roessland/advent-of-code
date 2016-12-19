package main

import "fmt"
import "errors"

const N int = 3001330

type Node struct {
	Number int
	Prev   *Node
	Next   *Node
	Items  int
}

func NewCircle() *Node {
	first := &Node{1, nil, nil, 1}
	prev := first
	for i := 2; i <= N; i++ {
		prev.Next = &Node{i, prev, nil, 1}
		prev = prev.Next
	}
	prev.Next = first
	first.Prev = prev
	return first
}

func (n *Node) Print() {
	first := n
	for {
		fmt.Printf("[[%d->( %d )->%d has %d]], ", n.Prev.Number, n.Number, n.Next.Number, n.Items)
		n = n.Next
		if n == first {
			break
		}
	}
	fmt.Println()
}

func (to *Node) Take(n int, from *Node) (*Node, *Node, error) {
	if to == from || to.Next == to {
		return to, nil, errors.New("no-op")
	}
	to.Items += from.Items
	from.Prev.Next = from.Next
	from.Next.Prev = from.Prev

	if to.Next == from {
		return from.Next, nil, errors.New("done")
	} else {
		// Draw it up to see the logic here.
		// Jump two if currently odd in circle,
		// jump one if currently even in circle.
		if n%2 == 0 {
			return to.Next, from.Next, nil
		} else {
			return to.Next, from.Next.Next, nil
		}
	}
}

func main() {
	circ := NewCircle()
	from := circ
	for i := 0; i < N/2; i++ {
		from = from.Next
	}
	n := N
	var err error
	for err == nil {
		circ, from, err = circ.Take(n, from)
		n--
	}
	circ.Print()
}
