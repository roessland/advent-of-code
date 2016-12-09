package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Input    string
	Expected string
}

func TestDecompress(t *testing.T) {

	table := []TestCase{
		TestCase{"ADVENT", "ADVENT"},
		TestCase{"A(1x5)BC", "ABBBBBC"},
		TestCase{"(3x3)XYZ", "XYZXYZXYZ"},
		TestCase{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
		TestCase{"(6x1)(1x3)A", "(1x3)A"},
		TestCase{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
		TestCase{"X(8x2)(3x3)ABCY(2x2)HA", "X(3x3)ABC(3x3)ABCYHAHA"},
	}
	for _, testCase := range table {
		var buf bytes.Buffer
		f := bufio.NewWriter(&buf)
		Decompress([]byte(testCase.Input), f)
		f.Flush()
		assert.Equal(t, testCase.Expected, buf.String())
	}
}
