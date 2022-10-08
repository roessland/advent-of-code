package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"strings"
	"testing"
)

const exampleInput = `
#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.##....#
#G..G..G#
#########`

func TestReadMap(t *testing.T) {
	m := ReadMap(strings.NewReader(exampleInput))
	require.Equal(t, 9, len(m))
	require.Equal(t, 9, len(m[1]))
	require.EqualValues(t, Wall, m[4][0].Type)
	require.EqualValues(t, Open, m[4][1].Type)
	require.EqualValues(t, Goblin, m[4][1].Unit.Type)
	require.EqualValues(t, 200, m[4][1].Unit.HitPoints)
	require.EqualValues(t, 3, m[4][1].Unit.AttackPower)
}

func TestDiffusionNoSources(t *testing.T) {
	m := ReadMap(strings.NewReader(exampleInput))
	dm := m.Diffusion(nil)
	require.Equal(t, len(m), len(dm))
	require.Equal(t, len(m[0]), len(dm[0]))
	require.EqualValues(t, math.MaxInt32, dm[0][0])
}

func TestDiffusion1Source(t *testing.T) {
	m := ReadMap(strings.NewReader(exampleInput))
	dm := m.Diffusion([]Pos{Pos{2, 1}})
	require.EqualValues(t, 1, dm[1][3])
	require.EqualValues(t, 9, dm[7][5])
	require.EqualValues(t, 5, dm[1][5])
	fmt.Println(dm)
}

func TestDiffusion2Sources(t *testing.T) {
	m := ReadMap(strings.NewReader(exampleInput))
	dm := m.Diffusion([]Pos{Pos{2, 1}, Pos{4, 6}})
	require.EqualValues(t, 1, dm[1][3])
	require.EqualValues(t, 2, dm[7][5])
	require.EqualValues(t, 5, dm[1][5])
}

func TestRound23(t *testing.T) {
	m := ReadMap(strings.NewReader(`
		#######
		#...G.#
		#..G.G#
		#.#.#G#
		#...#E#
		#.....#
		#######`))

	m.DoRound()
	m.Print()

	m.IterateUnits(func(u *Unit) {
		fmt.Println(u)
	})
}

func TestRoundSample1(t *testing.T) {
	m := ReadMap(strings.NewReader(`
		#######   
		#.G...#
		#...EG#
		#.#.#G#
		#..G#E#
		#.....#
		#######`))

	for i := 0; i < 23; i++ {
		m.DoRound()
	}
	m.Print()

	m.IterateUnits(func(u *Unit) {
		fmt.Println(u)
	})

	require.EqualValues(t, `#######
#...G.#
#..G.G#
#.#.#G#
#...#E#
#.....#
#######
`, m.String())

	m.DoRound()
}

func TestSampleRound26Diffusion(t *testing.T) {
	m := ReadMap(strings.NewReader(`#######
		#G....#
		#.G...#
		#.#.#G#
		#...#E#
		#..G..#
		#######`))

	dm := m.Diffusion([]Pos{{5, 5}})
	PrintDiffusion(dm)
}
