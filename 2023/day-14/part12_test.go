package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	p := &Platform{
		data: [][]byte{
			{0, 1},
			{2, 3},
			{4, 5},
		},
		width:  2,
		height: 3,
	}

	// Transpose
	// 024
	// 125
	//
	// Flip
	// y=newHeight - y -1

	// 01 (0,1) -> (0,1)
	// 23
	// 45
	p.rotation = 0
	assert.EqualValues(t, Pos{0, 0}, p.translate(Pos{0, 0}))
	assert.EqualValues(t, Pos{0, 1}, p.translate(Pos{0, 1}))
	assert.EqualValues(t, Pos{1, 0}, p.translate(Pos{1, 0}))
	assert.EqualValues(t, Pos{1, 1}, p.translate(Pos{1, 1}))
	assert.EqualValues(t, Pos{2, 0}, p.translate(Pos{2, 0}))
	assert.EqualValues(t, Pos{2, 1}, p.translate(Pos{2, 1}))

	// 135 (0,1) -> (1, 1)
	// 024
	p.rotation = 1
	assert.EqualValues(t, Pos{0, 1}, p.translate(Pos{0, 0}))
	assert.EqualValues(t, Pos{1, 1}, p.translate(Pos{0, 1}))
	assert.EqualValues(t, Pos{2, 1}, p.translate(Pos{0, 2}))
	assert.EqualValues(t, Pos{0, 0}, p.translate(Pos{1, 0}))
	assert.EqualValues(t, Pos{1, 0}, p.translate(Pos{1, 1}))
	assert.EqualValues(t, Pos{2, 0}, p.translate(Pos{1, 2}))

	// 54
	// 32
	// 10
	p.rotation = 2
	assert.EqualValues(t, Pos{2, 1}, p.translate(Pos{0, 0}))
	assert.EqualValues(t, Pos{2, 0}, p.translate(Pos{0, 1}))
	assert.EqualValues(t, Pos{1, 1}, p.translate(Pos{1, 0}))
	assert.EqualValues(t, Pos{1, 0}, p.translate(Pos{1, 1}))
	assert.EqualValues(t, Pos{0, 1}, p.translate(Pos{2, 0}))
	assert.EqualValues(t, Pos{0, 0}, p.translate(Pos{2, 1}))

	// 420
	// 531
	p.rotation = 3
	assert.EqualValues(t, Pos{2, 0}, p.translate(Pos{0, 0}))
	assert.EqualValues(t, Pos{1, 0}, p.translate(Pos{0, 1}))
	assert.EqualValues(t, Pos{0, 0}, p.translate(Pos{0, 2}))
	assert.EqualValues(t, Pos{2, 1}, p.translate(Pos{1, 0}))
	assert.EqualValues(t, Pos{1, 1}, p.translate(Pos{1, 1}))
	assert.EqualValues(t, Pos{0, 1}, p.translate(Pos{1, 2}))
}
