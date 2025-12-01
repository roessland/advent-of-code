// Package day01 solves AoC 2025 Day 1
package day01

import (
	"embed"
	"log/slog"
	"strings"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/mathutil"
)

const K = 100

//go:embed input*.txt
var InputFile embed.FS

func ReadInput(inputName string) (rots []int) {
	for _, line := range aocutil.ReadLines(inputName) {
		line = strings.ReplaceAll(line, "R", "")
		line = strings.ReplaceAll(line, "L", "-")
		rots = append(rots, aocutil.Atoi(line))
	}
	return rots
}

func CumSumMod(ns []int, k int) (out []int) {
	sum := 0
	for _, rot := range ns {
		sum = (sum + rot) % k
		out = append(out, sum)
	}
	return out
}

func NumZeros(ns []int) (numZeros int) {
	for _, n := range ns {
		if n == 0 {
			numZeros++
		}
	}
	return numZeros
}

func Part1(rots []int) int {
	rots = append([]int{50}, rots...)
	return NumZeros(CumSumMod(rots, K))
}

func Rot(start, rot int) (end int) {
	return (((start + rot) % K) + K) % K
}

type Expl struct {
	FullRotations, Passthroughs, EndedUp int
}

func (ve Expl) Total() int {
	return ve.FullRotations + ve.Passthroughs + ve.EndedUp
}

func EnsureOrderedInterval(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func ModInterval(a, b int) (int, int) {
	a, b = EnsureOrderedInterval(a, b)
	for a <= -100 {
		a, b = a+100, b+100
	}
	for b >= 100 {
		a, b = a-100, b-100
	}
	return a, b
}

func VisitsVerbose(startRot, rot int) (expl Expl) {
	slog.Debug("VisitsVerbose", slog.Int("startRot", startRot), slog.Int("rot", rot))
	// ++ since we do one full rotation
	expl.FullRotations = mathutil.AbsInt(rot) / K
	remainingRot := rot % K

	finalEndRot := Rot(startRot, rot)

	a, b := ModInterval(startRot, startRot+remainingRot)
	slog.Debug("modInterval", slog.Int("a", a), slog.Int("b", b))
	if a <= 0 && b >= 0 && finalEndRot != 0 {
		expl.Passthroughs++
		// if finalEndRot == 0 {
		// 	expl.Passthroughs--
		// }
	}

	// -- if started at 0
	if startRot == 0 {
		expl.Passthroughs--
	}

	// ++ since we end up at zero
	if finalEndRot == 0 {
		expl.EndedUp++
	}

	return expl
}

func Part2(rots []int) int {
	visits := 0
	// visitsSlow := 0
	curr := 50
	for _, rot := range rots {
		next := Rot(curr, rot)
		expl := VisitsVerbose(curr, rot)

		dVisits := expl.Total()
		visits += dVisits

		// dVisitsSlow := VisitsSlow(curr, rot)
		// visitsSlow += dVisitsSlow

		// slog.Info("moved", slog.Int("from", curr), "rot", rot, "to", next, "visits", dVisits, "visitsSlow", dVisitsSlow)
		// if dVisits != dVisitsSlow {
		// 	slog.Error("WAT")
		// }
		curr = next
	}

	return visits
}

func VisitsSlow(startRot, remainingRot int) (visits int) {
	sign := mathutil.SignInt(remainingRot)
	curr := startRot

	for remainingRot != 0 {
		curr += sign
		remainingRot -= sign

		curr %= K
		if curr == 0 {
			visits++
		}
	}
	return visits
}
