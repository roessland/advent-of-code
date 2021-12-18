package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Number struct {
	parent *Number
	value  int
	left   *Number
	right  *Number
}

func (n *Number) Copy() *Number {
	if n == nil {
		return nil
	}
	c := &Number{}
	if n.HasValue() {
		c.value = n.value
	} else {
		c.left = n.left.Copy()
		c.left.parent = c
		c.right = n.right.Copy()
		c.right.parent = c
	}
	return c
}

func (n *Number) Add(m *Number) *Number {
	sum := &Number{left: n.Copy(), right: m.Copy()}
	sum.left.parent = sum
	sum.right.parent = sum
	sum.Reduce()
	return sum
}

func (n *Number) HasValue() bool {
	return n.left == nil && n.right == nil
}

func (n *Number) String() string {
	if n.HasValue() {
		return fmt.Sprintf("%d", n.value)
	} else {
		return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
	}
}

func (n *Number) Reduce() {
	for {
		explodable := n.findExplodable(0)
		if explodable != nil {
			explodable.Explode()
			continue
		}

		splittable := n.findSplittable()
		if splittable != nil {
			splittable.Split()
			continue
		}

		break
	}
}

func (n *Number) Split() {
	down := n.value / 2
	up := n.value / 2
	if n.value%2 == 1 {
		up++
	}
	n.left = &Number{value: down, parent: n}
	n.right = &Number{value: up, parent: n}
	n.value = 0
}

func (n *Number) Magnitude() int {
	if n.HasValue() {
		return n.value
	}
	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func (n *Number) findSplittable() *Number {
	if n == nil {
		return nil
	}
	if n.value >= 10 {
		return n
	}
	if leftSplittable := n.left.findSplittable(); leftSplittable != nil {
		return leftSplittable
	}
	if rightSplittable := n.right.findSplittable(); rightSplittable != nil {
		return rightSplittable
	}
	return nil
}

func (n *Number) Explode() {
	if n.left == nil || n.right == nil {
		panic("cannot explode number without left and right")
	}
	leftInorder := n.findLeft()
	rightInorder := n.findRight()
	if leftInorder != nil {
		leftInorder.value += n.left.value
	}
	if rightInorder != nil {
		rightInorder.value += n.right.value
	}
	n.left = nil
	n.right = nil
	n.value = 0
}

func (n *Number) findExplodable(level int) *Number {
	if n == nil || n.HasValue() {
		return nil
	}
	if level == 4 {
		return n
	}
	if leftExplodable := n.left.findExplodable(level + 1); leftExplodable != nil {
		return leftExplodable
	}
	if rightExplodable := n.right.findExplodable(level + 1); rightExplodable != nil {
		return rightExplodable
	}
	return nil
}

func (n *Number) findLeft() *Number {
	// Move up until some node has a left branch
	for {
		if n.parent == nil {
			return nil
		}
		if n.parent.left != n && n.parent.left != nil {
			n = n.parent.left
			break
		}
		n = n.parent
	}
	// Move down all right branches until a value is found
	for !n.HasValue() {
		n = n.right
	}
	return n
}

func (n *Number) findRight() *Number {
	// Move up until some node has a right branch
	for {
		if n.parent == nil {
			return nil
		}
		if n.parent.right != n && n.parent.right != nil {
			n = n.parent.right
			break
		}
		n = n.parent
	}
	// Move down all left branches until a value is found
	for !n.HasValue() {
		n = n.left
	}
	return n
}

func FromString(str string) *Number {
	var v []interface{}
	decoder := json.NewDecoder(strings.NewReader(str))
	decoder.UseNumber()
	err := decoder.Decode(&v)
	if err != nil {
		log.Fatal(err)
	}

	var f func(listOrInt interface{}, parent *Number) *Number
	f = func(listOrInt interface{}, parent *Number) *Number {
		switch v := listOrInt.(type) {
		case json.Number:
			n, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			return &Number{value: n, parent: parent}
		case []interface{}:
			n := &Number{parent: parent}
			n.left = f(v[0], n)
			n.right = f(v[1], n)
			return n
		default:
			log.Fatalf("unknown type %T", v)
		}
		return nil
	}
	return f(v, nil)
}

func main() {
	t0 := time.Now()
	part1()
	part2()
	fmt.Println(time.Since(t0))
}

func part1() {
	nums := ReadInput()
	for i := 1; i < len(nums); i++ {
		nums[0] = nums[0].Add(nums[i])
	}
	fmt.Println("Part 1:", nums[0].Magnitude())
}

func part2() {
	nums := ReadInput()
	maxMagnitude := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			magnitude := nums[i].Add(nums[j]).Magnitude()
			maxMagnitude = mathutil.MaxInt(maxMagnitude, magnitude)
		}

	}
	fmt.Println("Part 2:", maxMagnitude)
}

func ReadInput() []*Number {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var nums []*Number
	for scanner.Scan() {
		num := FromString(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
