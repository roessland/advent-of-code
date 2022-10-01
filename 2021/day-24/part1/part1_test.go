package main

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

import _ "embed"

//go:embed input.txt
var fs embed.FS

func TestComputeNegate(t *testing.T) {
	instrs := []Instruction{
		{Type: Inp, A: 'x'},
		{Type: MulC, A: 'x', B: -1},
	}
	tcs := []struct {
		Name  string
		X     int
		Input []int
	}{
		{"-1 is -1", -1, []int{1}},
		{"-5 is -5", -5, []int{5}},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			reg := Compute(instrs, tc.Input, Reg{})
			require.EqualValues(t, tc.X, reg.X)
		})
	}
}

func TestComputeIsTriple(t *testing.T) {
	instrs := []Instruction{
		{Type: Inp, A: 'z'},
		{Type: Inp, A: 'x'},
		{Type: MulC, A: 'z', B: 3},
		{Type: EqlP, A: 'z', B: 'x'},
	}
	tcs := []struct {
		Name   string
		Input  []int
		Expect int
	}{
		{
			"3*0 is 9",
			[]int{0, 0},
			1,
		},
		{
			"3*1 is 3",
			[]int{1, 3},
			1,
		},
		{
			"3*2 is not 5",
			[]int{2, 5},
			0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			reg := Compute(instrs, tc.Input, Reg{})
			assert.Equal(t, tc.Expect, reg.Z)
		})
	}
}

func TestComputeBinary(t *testing.T) {
	instrs := []Instruction{
		{Type: Inp, A: 'w'},
		{Type: AddP, A: 'z', B: 'w'},
		{Type: ModC, A: 'z', B: 2},
		{Type: DivC, A: 'w', B: 2},
		{Type: AddP, A: 'y', B: 'w'},
		{Type: ModC, A: 'y', B: 2},
		{Type: DivC, A: 'w', B: 2},
		{Type: AddP, A: 'x', B: 'w'},
		{Type: ModC, A: 'x', B: 2},
		{Type: DivC, A: 'w', B: 2},
		{Type: ModC, A: 'w', B: 2},
	}
	tcs := []struct {
		Name   string
		Input  []int
		Expect string
	}{
		{
			"1 is 0b001",
			[]int{1},
			"0001",
		},
		{
			"2 is 0b010",
			[]int{2},
			"0010",
		},
		{
			"15 is 0b1111",
			[]int{15},
			"1111",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			reg := Compute(instrs, tc.Input, Reg{})
			w, x, y, z := reg.W, reg.X, reg.Y, reg.Z
			got := fmt.Sprintf("%d%d%d%d", w, x, y, z)
			assert.Equal(t, tc.Expect, got)
		})
	}
}

func TestComputeInpOut(t *testing.T) {
	instrs := []Instruction{
		{Type: Inp, A: 'w'},
		{Type: MulC, A: 'w', B: 2},
		{Type: Inp, A: 'x'},
		{Type: MulC, A: 'x', B: 3},
		{Type: Inp, A: 'y'},
		{Type: MulC, A: 'y', B: 4},
		{Type: Inp, A: 'z'},
		{Type: MulC, A: 'z', B: 5},
	}
	reg := Compute(instrs, []int{1, 2, 3, 4}, Reg{})
	require.EqualValues(t, Reg{1 * 2, 2 * 3, 3 * 4, 4 * 5}, reg)
}

func TestComputePtrs(t *testing.T) {
	// Inp 2
	//	AddP 2 + 3 = 5
	//	MulP (2+3)*4 = 20
	//	DivP (2+3)*4 / 2 = 10
	//	ModP (2+3)*4 / 2 % 3 = 1
	// Inp 1
	//	EqlP
	instrs := []Instruction{
		{Type: Inp, A: 'w'}, // w=2
		{Type: Inp, A: 'x'}, // x=3

		{Type: AddP, A: 'w', B: 'x'}, // w=5

		{Type: Inp, A: 'x'},          // x=4
		{Type: MulP, A: 'w', B: 'x'}, // w=20
		{Type: AddP, A: 'z', B: 'w'}, // z=20

		{Type: Inp, A: 'x'},          // x=2
		{Type: DivP, A: 'w', B: 'x'}, // w=10
		{Type: AddP, A: 'y', B: 'w'}, // y=10

		{Type: Inp, A: 'x'},          // x=3
		{Type: ModP, A: 'w', B: 'x'}, // w=1

		{Type: Inp, A: 'x'},          // x=1
		{Type: EqlP, A: 'x', B: 'w'}, // x=1
	}
	reg := Compute(instrs, []int{2, 3, 4, 2, 3, 1}, Reg{})
	require.EqualValues(t, Reg{1, 1, 10, 20}, reg)
}

func TestTrimSpaceTrimsTabs(t *testing.T) {
	input := "\tbanana"
	require.EqualValues(t, "banana", strings.TrimSpace(input))
}

func TestComputeParseString(t *testing.T) {
	input := `
		inp w // get input // yo
		inp x
		add w x
		inp x
		mul w x
		add z w
		inp x
		div w x
		add y w
		inp x
		mod w x
		inp x
		eql x w`
	actualInstrs := ParseString(input)
	expectedInstrs := []Instruction{
		{Type: Inp, A: 'w', Text: "inp w // get input // yo"},
		{Type: Inp, A: 'x', Text: "inp x"},
		{Type: AddP, A: 'w', B: 'x', Text: "add w x"},
		{Type: Inp, A: 'x', Text: "inp x"},
		{Type: MulP, A: 'w', B: 'x', Text: "mul w x"},
		{Type: AddP, A: 'z', B: 'w', Text: "add z w"},
		{Type: Inp, A: 'x', Text: "inp x"},
		{Type: DivP, A: 'w', B: 'x', Text: "div w x"},
		{Type: AddP, A: 'y', B: 'w', Text: "add y w"},
		{Type: Inp, A: 'x', Text: "inp x"},
		{Type: ModP, A: 'w', B: 'x', Text: "mod w x"},
		{Type: Inp, A: 'x', Text: "inp x"},
		{Type: EqlP, A: 'x', B: 'w', Text: "eql x w"},
	}
	require.EqualValues(t, expectedInstrs, actualInstrs)
}

func BenchmarkComputeInput(b *testing.B) {
	f, err := fs.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	instrs := ParseInput(f)
	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}

	for j := 0; j < b.N; j++ {
		if j%1000 == 0 {
			RandomizeInput(inps)
		}
		Compute(instrs, inps, Reg{})
	}
}

func BenchmarkGeneratedCode(b *testing.B) {
	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}

	for j := 0; j < b.N; j++ {
		if j%1000 == 0 {
			RandomizeInput(inps)
		}
		F13(inps)
	}
}

func BenchmarkCodegen(B *testing.B) {
	f, err := fs.Open("input.txt")
	if err != nil {
		B.Fatal(err)
	}
	instrs := ParseInput(f)
	a, b, c := ExtractABC(instrs)

	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	for j := 0; j < B.N; j++ {
		if j%1000 == 0 {
			RandomizeInput(inps)
		}
		Codegen(a, b, c, inps)
	}
}

func BenchmarkModuleMemoization(b *testing.B) {
	f, err := fs.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	instrs := ParseInput(f)
	_, F := GetModulesWithMemoization(instrs)

	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	for j := 0; j < b.N; j++ {
		if j%100 == 0 {
			RandomizeInput(inps)
		}
		F[13](inps)
	}
}

func TestGetModulesWithMemoization(t *testing.T) {
	f, err := fs.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	instrs := ParseInput(f)

	_, F := GetModulesWithMemoization(instrs)

	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	require.Equal(t, Compute(instrs, inps, Reg{}).Z, F[13](inps))
}

func TestCodegen(t *testing.T) {
	f, err := fs.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	instrs := ParseInput(f)
	a, b, c := ExtractABC(instrs)

	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	require.Equal(t, Compute(instrs, inps, Reg{}).Z, Codegen(a, b, c, inps))
}

func TestF13(t *testing.T) {
	f, err := fs.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	instrs := ParseInput(f)

	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	for i := 0; i < 100; i++ {
		RandomizeInput(inps)
		require.Equal(t, Compute(instrs, inps, Reg{}).Z, F13(inps))
	}
}

func BenchmarkRandomizeInput(b *testing.B) {
	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	for j := 0; j < b.N; j++ {
		RandomizeInput(inps)
	}
}

func BenchmarkPartialRandomizeInput(b *testing.B) {
	inps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	for j := 0; j < b.N; j++ {
		PartialRandomizeInput(inps)
	}
}
