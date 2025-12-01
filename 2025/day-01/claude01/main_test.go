package main

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		line      string
		direction byte
		distance  int
	}{
		{"L68", 'L', 68},
		{"R48", 'R', 48},
		{"L5", 'L', 5},
	}

	for _, tt := range tests {
		dir, dist := parseLine(tt.line)
		if dir != tt.direction || dist != tt.distance {
			t.Errorf("parseLine(%q) = (%c, %d), want (%c, %d)", tt.line, dir, dist, tt.direction, tt.distance)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		start    int
		dir      byte
		dist     int
		expected int
	}{
		{50, 'L', 68, 82},  // 50 - 68 = -18 -> 82
		{82, 'L', 30, 52},  // 82 - 30 = 52
		{52, 'R', 48, 0},   // 52 + 48 = 100 -> 0
		{0, 'L', 5, 95},    // 0 - 5 = -5 -> 95
		{95, 'R', 60, 55},  // 95 + 60 = 155 -> 55
		{55, 'L', 55, 0},   // 55 - 55 = 0
		{0, 'L', 1, 99},    // 0 - 1 = -1 -> 99
		{99, 'L', 99, 0},   // 99 - 99 = 0
		{0, 'R', 14, 14},   // 0 + 14 = 14
		{14, 'L', 82, 32},  // 14 - 82 = -68 -> 32
		{11, 'R', 8, 19},   // from problem description
		{19, 'L', 19, 0},   // from problem description
		{5, 'L', 10, 95},   // from problem description
		{95, 'R', 5, 0},    // from problem description
	}

	for _, tt := range tests {
		result := rotate(tt.start, tt.dir, tt.dist)
		if result != tt.expected {
			t.Errorf("rotate(%d, '%c', %d) = %d, want %d", tt.start, tt.dir, tt.dist, result, tt.expected)
		}
	}
}

func TestSolve(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	result := solve(input)
	if result != 3 {
		t.Errorf("solve(example) = %d, want 3", result)
	}
}

func TestCountZeroCrossings(t *testing.T) {
	tests := []struct {
		pos      int
		dir      byte
		dist     int
		expected int
	}{
		// From part 2 example
		{50, 'L', 68, 1},  // L68 from 50 passes through 0 once
		{82, 'L', 30, 0},  // L30 from 82 doesn't pass 0
		{52, 'R', 48, 1},  // R48 from 52 hits 0 (lands on it)
		{0, 'L', 5, 0},    // L5 from 0 doesn't pass 0 again
		{95, 'R', 60, 1},  // R60 from 95 passes through 0 once
		{55, 'L', 55, 1},  // L55 from 55 hits 0 (lands on it)
		{0, 'L', 1, 0},    // L1 from 0 doesn't pass 0
		{99, 'L', 99, 1},  // L99 from 99 hits 0 (lands on it)
		{0, 'R', 14, 0},   // R14 from 0 doesn't pass 0
		{14, 'L', 82, 1},  // L82 from 14 passes through 0 once
		// Special case from problem: R1000 from 50 hits 0 ten times
		{50, 'R', 1000, 10},
	}

	for _, tt := range tests {
		result := countZeroCrossings(tt.pos, tt.dir, tt.dist)
		if result != tt.expected {
			t.Errorf("countZeroCrossings(%d, '%c', %d) = %d, want %d", tt.pos, tt.dir, tt.dist, result, tt.expected)
		}
	}
}

func TestSolve2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	result := solve2(input)
	if result != 6 {
		t.Errorf("solve2(example) = %d, want 6", result)
	}
}
