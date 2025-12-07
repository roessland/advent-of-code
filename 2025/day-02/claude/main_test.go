package main

import "testing"

func TestIsInvalidID(t *testing.T) {
	tests := []struct {
		id   int64
		want bool
	}{
		// Part 2: repeated at least twice
		{11, true},           // 1 repeated 2 times
		{22, true},           // 2 repeated 2 times
		{99, true},           // 9 repeated 2 times
		{111, true},          // 1 repeated 3 times
		{999, true},          // 9 repeated 3 times
		{1111111, true},      // 1 repeated 7 times
		{6464, true},         // 64 repeated 2 times
		{123123, true},       // 123 repeated 2 times
		{12341234, true},     // 1234 repeated 2 times
		{123123123, true},    // 123 repeated 3 times
		{1212121212, true},   // 12 repeated 5 times
		{1010, true},         // 10 repeated 2 times
		{1188511885, true},   // 11885 repeated 2 times
		{222222, true},       // 222 repeated 2 times
		{446446, true},       // 446 repeated 2 times
		{38593859, true},     // 3859 repeated 2 times
		{565656, true},       // 5656 repeated 2 times
		{824824824, true},    // 824 repeated 3 times
		{2121212121, true},   // 2121 repeated 2 times, or 21 repeated 5 times

		{101, false},         // not a repeat pattern
		{12, false},          // different digits
		{100, false},         // odd length
		{1234, false},        // not a repeat
		{5, false},           // single digit
		{1698522, false},     // not a repeat
	}

	for _, tt := range tests {
		got := isInvalidID(tt.id)
		if got != tt.want {
			t.Errorf("isInvalidID(%d) = %v, want %v", tt.id, got, tt.want)
		}
	}
}

func TestCountInvalidIDsInRange(t *testing.T) {
	tests := []struct {
		start, end int64
		wantSum    int64
	}{
		// Part 2 ranges from example
		{11, 22, 11 + 22},                        // 11 and 22
		{95, 115, 99 + 111},                      // 99 and 111
		{998, 1012, 999 + 1010},                  // 999 and 1010
		{1188511880, 1188511890, 1188511885},     // just 1188511885
		{222220, 222224, 222222},                 // just 222222
		{1698522, 1698528, 0},                    // none
		{446443, 446449, 446446},                 // just 446446
		{38593856, 38593862, 38593859},           // just 38593859
		{565653, 565659, 565656},                 // just 565656
		{824824821, 824824827, 824824824},        // just 824824824
		{2121212118, 2121212124, 2121212121},     // just 2121212121
	}

	for _, tt := range tests {
		got := sumInvalidIDsInRange(tt.start, tt.end)
		if got != tt.wantSum {
			t.Errorf("sumInvalidIDsInRange(%d, %d) = %d, want %d", tt.start, tt.end, got, tt.wantSum)
		}
	}
}

func TestSolveExample(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	want := int64(4174379265)
	got := solve(input)
	if got != want {
		t.Errorf("solve(example) = %d, want %d", got, want)
	}
}
