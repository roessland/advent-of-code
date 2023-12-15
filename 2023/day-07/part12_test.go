package main

import (
	"testing"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"github.com/stretchr/testify/assert"
)

func h(s string) Hand {
	ms := aocutil.NewImmutableMultiSet[Card]([]Card(s)...)
	return Hand{ImmutableMultiSet: *ms}
}

func TestValues(t *testing.T) {
	assert.Greater(t, h("123QQ").Value(), h("12345").Value())
	assert.Greater(t, h("123AA").Value(), h("123QQ").Value())
}
