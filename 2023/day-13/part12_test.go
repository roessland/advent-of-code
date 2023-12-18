package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflect(t *testing.T) {
	assert.Equal(t, 5, Reflect(0, 2))
	assert.Equal(t, 0, Reflect(5, 2))
	assert.Equal(t, 1, Reflect(0, 0))
	assert.Equal(t, 0, Reflect(1, 0))
	assert.Equal(t, -1, Reflect(0, -1))
}
