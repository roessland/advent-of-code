package day05

import "testing"

func BenchmarkPart1(b *testing.B) {
	ranges, ids := ReadInput("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part1(ranges, ids)
	}
}

func BenchmarkPart2(b *testing.B) {
	ranges, _ := ReadInput("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part2(ranges)
	}
}

func BenchmarkIsFresh(b *testing.B) {
	ranges, ids := ReadInput("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsFresh(ranges, ids[i%len(ids)])
	}
}

func BenchmarkRangeSub(b *testing.B) {
	r1 := Range{0, 10000000}
	r2 := Range{5000000, 15000000}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r1.Sub(r2)
	}
}
