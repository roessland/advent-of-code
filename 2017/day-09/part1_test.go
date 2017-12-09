package main

import "testing"
import "fmt"
import "strings"
import "github.com/stretchr/testify/assert"

type TestCase struct {
	Expected int
	Input    string
}

func TestParser(t *testing.T) {
	cases := []TestCase{
		{1, "{}"},
		{1, "{<a>,<a>,<a>,<a>}"},
		{9, "{{<ab>},{<ab>},{<ab>},{<ab>}}"},
		{9, "{{<!!>},{<!!>},{<!!>},{<!!>}}"},
	}
	for _, c := range cases {
		fmt.Printf("testcase: '%s'\n", c.Input)
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			sp := NewStreamParser(strings.NewReader(c.Input))
			assert.Equal(t, c.Expected, sp.FindTotalScore())
		})
	}
}
