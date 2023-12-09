package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	input := ReadInput()
	part1(input)
	part2(input)
}

type Entity struct {
	Cat string
	Num int
}

// MapFunc maps a number to to an entity
type MapFunc func(int) Entity

// makeMapFunc creates a MapFunc for a given mapping input section.
// For example, a seed-to-soil map func.
func makeMapFunc(mapInput MapInput) MapFunc {
	return func(num int) (ent Entity) {
		ent.Cat = mapInput.DstCat
		for _, mapping := range mapInput.Mappings {
			if mapping.SrcStart <= num && num < mapping.SrcStart+mapping.RangeLen {
				ent.Num = mapping.DstStart + (num - mapping.SrcStart)
				return ent
			}
		}
		ent.Num = num
		return ent
	}
}

func part1(input Input) {
	mapForwards := make(map[string]MapFunc)
	for _, mapInput := range input.Maps {
		mapInput := mapInput // Loopvar: Remove in Go 1.22
		mapForwards[mapInput.SrcCat] = makeMapFunc(mapInput)
	}

	locationNums := []int{}
	for _, seedNum := range input.Seeds {
		ent := Entity{"seed", seedNum}
		for ent.Cat != "location" {
			ent = mapForwards[ent.Cat](ent.Num)
		}
		locationNums = append(locationNums, ent.Num)
	}
	sort.Ints(locationNums)
	fmt.Println("Part 1:", locationNums[0])
}

// Range is a half-open interval [Start, End)
type Range struct {
	Start, End int
}

func (r Range) Size() int {
	return r.End - r.Start
}

// Entity2 is like Entity, but is represented with ranges instead of numbers.
type Entity2 struct {
	Cat    string
	Ranges []Range
}

// ExtractIntersection returns the intersection between two ranges (if any),
// and the remaining input after removing the intersection.
// There can be 0, 1 or 2 remaining ranges.
func ExtractIntersection(in, srcRange Range) (overlapping Range, remaining []Range) {
	defer func() {
		// Remove empty ranges after returning
		if len(remaining) == 2 && remaining[1].Size() == 0 {
			remaining = []Range{remaining[0]}
		}
		if len(remaining) == 1 && remaining[0].Size() == 0 {
			remaining = []Range{}
		}
	}()

	// No intersection
	if in.End <= srcRange.Start || srcRange.End <= in.Start {
		overlapping = Range{0, 0}
		remaining = []Range{in}
		return
	}

	// Input contained in map source range
	if srcRange.Start <= in.Start && in.End <= srcRange.End {
		overlapping = in
		remaining = nil
		return
	}

	// Map source contained in input range
	if in.Start <= srcRange.Start && srcRange.End <= in.End {
		overlapping = srcRange
		remaining = []Range{
			{in.Start, srcRange.Start},
			{srcRange.End, in.End},
		}
		return
	}

	// Input ends in source range
	if in.End <= srcRange.End {
		overlapping = Range{srcRange.Start, in.End}
		remaining = []Range{
			{in.Start, srcRange.Start},
		}
		return
	}

	// Input starts in source range
	if srcRange.Start <= in.Start {
		overlapping = Range{in.Start, srcRange.End}
		remaining = []Range{
			{srcRange.End, in.End},
		}
		return
	}

	panic("Should not reach here")
}

// MapFunc2 is like MapFunc, but maps an entire list of ranges at once.
type MapFunc2 func([]Range) Entity2

// makeMapFunc2 creates a MapFunc2 for a given mapping input section.
func makeMapFunc2(mapInput MapInput) MapFunc2 {
	return func(rngs []Range) (ent Entity2) {
		ent.Cat = mapInput.DstCat

		// While there are unmapped ranges that can be mapped, map them
		for len(rngs) > 0 {
			// Pop the last unmapped range
			rng := rngs[len(rngs)-1]
			rngs = rngs[:len(rngs)-1]

			// Check if part of the range has a mapping
			mappedSomething := false
			for _, mapping := range mapInput.Mappings {
				intersection, remaining := ExtractIntersection(rng, mapping.SrcRange)

				// Mapped the intersection
				if intersection.Size() > 0 {
					mappedSomething = true
					offset := -mapping.SrcStart + mapping.DstStart
					ent.Ranges = append(ent.Ranges, Range{
						Start: intersection.Start + offset,
						End:   intersection.End + offset,
					})
					// Add remaining back in the stack
					rngs = append(rngs, remaining...)
				}
			}

			// Use default mapping otherwise
			if !mappedSomething {
				ent.Ranges = append(ent.Ranges, rng)
			}
		}

		return ent
	}
}

func part2(input Input) {
	// MapFuncs for each source category
	mapForwards := make(map[string]MapFunc2)
	for _, mapInput := range input.Maps {
		mapInput := mapInput // Loopvar: Remove in Go 1.22
		mapForwards[mapInput.SrcCat] = makeMapFunc2(mapInput)
	}

	// Make initial entity
	ent := Entity2{"seed", nil}
	for i := 0; i < len(input.Seeds); i += 2 {
		ent.Ranges = append(ent.Ranges, Range{input.Seeds[i], input.Seeds[i] + input.Seeds[i+1]})
	}

	// Map until we reach a location
	for ent.Cat != "location" {
		ent = mapForwards[ent.Cat](ent.Ranges)
	}

	// Find minimum location number
	minLocationNum := math.MaxInt32
	for _, rng := range ent.Ranges {
		if rng.Start < minLocationNum {
			minLocationNum = rng.Start
		}
	}

	fmt.Println("Part 2:", minLocationNum)
}

type Input struct {
	Seeds []int
	Maps  []MapInput
}

type MapInput struct {
	SrcCat, DstCat string
	Mappings       []MappingsInput
}

type MappingsInput struct {
	DstStart, SrcStart, RangeLen int
	SrcRange, DstRange           Range
}

func ReadInput() Input {
	inputStr := aocutil.ReadFile("input.txt")

	input := Input{}

	sections := strings.Split(inputStr, "\n\n")

	input.Seeds = aocutil.GetIntsInString(sections[0])

	for i := 1; i < len(sections); i++ {
		mapInput := MapInput{}
		lines := strings.Split(sections[i], "\n")
		words := strings.Split(strings.TrimSuffix(lines[0], " map:"), "-to-")
		mapInput.SrcCat = words[0]
		mapInput.DstCat = words[1]

		for _, line := range lines[1:] {
			nums := aocutil.GetIntsInString(line)
			mapInput.Mappings = append(mapInput.Mappings,
				MappingsInput{
					DstStart: nums[0],
					SrcStart: nums[1],
					RangeLen: nums[2],
					DstRange: Range{nums[0], nums[0] + nums[2]},
					SrcRange: Range{nums[1], nums[1] + nums[2]},
				})
		}

		input.Maps = append(input.Maps, mapInput)
	}

	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
