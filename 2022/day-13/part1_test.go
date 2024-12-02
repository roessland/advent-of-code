package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	check := func(s string) {
		n := Read(s)
		assert.Equal(t, s, fmt.Sprintf("%v", n), fmt.Sprintf("%#v", n))
	}

	check("10")
	check("1")
	check("[1]")
	check("[10]")
	check("[]")
	check("[1,2]")
	check("[2,[]]")
}

func TestSplit(t *testing.T) {
	assert.EqualValues(t, []string{"1", "2", "3"}, Split("1,2,3"))
	assert.EqualValues(t, []string{"2", "[]"}, Split("2,[]"))
	assert.EqualValues(t, []string{"1", ""}, Split("1,"))
	assert.EqualValues(t, []string{"10", "[]"}, Split("10,[]"))
	assert.EqualValues(t, []string{}, Split(""))
	assert.EqualValues(t, []string{"[]", "2"}, Split("[],2"))
}

func TestString(t *testing.T) {
	assert.Equal(t, "1", Node{IsLeaf: true, Num: 1}.String())
	assert.Equal(t, "[]", Node{IsLeaf: false}.String())
	assert.Equal(t, "[[]]", Node{IsLeaf: false, Children: []Node{{}}}.String())
}

func TestCompare(t *testing.T) {
	assert.Equal(t, -1, Compare(Node{IsLeaf: true, Num: 1}, Node{IsLeaf: true, Num: 2}))
	assert.Equal(t, 1, Compare(Node{IsLeaf: true, Num: 3}, Node{IsLeaf: true, Num: 2}))
	assert.Equal(t, -1, Compare(Node{IsLeaf: false}, Node{IsLeaf: true, Num: 3}))
	assert.Equal(t, 1, Compare(Node{IsLeaf: true, Num: 3}, Node{IsLeaf: false}))
	assert.Equal(t, -1, Compare(Read("[4,4,4]"), Read("[4,4,4,4]")))
	assert.Equal(t, 0, Compare(Read("[4,[],4]"), Read("[4,[],4]")))
	assert.Equal(t, -1, Compare(Read("[[4,4],4,4]"), Read("[[4,4],4,4,4]")))
	assert.Equal(t, 0, Compare(Read("[[4,4],4,4]"), Read("[[4,4],4,4]")))
	assert.Equal(t, 1, Compare(Read("[[4,4],4,4,4]"), Read("[4,4],4,4")))
	assert.Equal(t, 0, Compare(Read("[4,4]"), Read("[4,4]")))
	assert.Equal(t, -1, Compare(Read("[1,1,3,1,1]"), Read("[1,1,5,1,1]")))
	assert.Equal(t, -1, Compare(Read("[1,1,3,1]"), Read("[1,1,5,1]")))
	assert.Equal(t, -1, Compare(Read("[1,1,3]"), Read("[1,1,5]")))
	assert.Equal(t, 0, Compare(Read("[1,1]"), Read("[1,1]")))
	assert.Equal(t, -1, Compare(Read("[3,1,3]"), Read("[5,1,5]")))
}
