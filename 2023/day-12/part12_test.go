package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func c(t *testing.T, expec int, s string, cont bool, groups ...int) {
	t.Run(fmt.Sprintf("%s_%v", s, groups), func(t *testing.T) {
		got := Arrs([]byte(s), groups, cont)
		assert.Equal(t, expec, got, s)
	})
}

func TestArrs(t *testing.T) {
	c(t, 1, "", false)
	c(t, 0, "#", false)
	c(t, 1, "#", false, 1)
	c(t, 1, "?", false, 1)
	c(t, 0, ".", false, 1)
	c(t, 1, "##", false, 2)
	c(t, 1, "#?", false, 2)
	c(t, 1, "?#", false, 2)
	c(t, 1, "??", false, 2)
	c(t, 1, "???", false, 3)
	c(t, 0, "???", false, 4)
	c(t, 0, "???", false, 5)
	c(t, 2, "??", false, 1)
	c(t, 3, "???", false, 1)
	c(t, 2, "???", false, 2)
	c(t, 0, "?", false, 1, 1, 1, 1, 1, 1)

	c(t, 1, "???", false, 1, 1)
	{
		c(t, 1, "??", true, 0, 1)
		c(t, 0, "??", false, 1, 1)
		{
			c(t, 1, "?", true, 1)
			c(t, 0, "?", false, 1, 1)
		}
	}

	c(t, 10, "?###????????", false, 3, 2, 1)

	c(t, 0, "#????#?????", false)
}
