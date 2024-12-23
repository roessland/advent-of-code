package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	{
		a, b := Part12("input-ex1.txt")
		assert.Equal(t, a, 7)
		assert.Equal(t, b, "co,de,ka,ta")
	}
	{
		a, b := Part12("input.txt")
		assert.Equal(t, a, 1323)
		assert.Equal(t, b, "er,fh,fi,ir,kk,lo,lp,qi,ti,vb,xf,ys,yu")
	}
}
