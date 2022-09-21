package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCuboid_Intersect_Simple(t *testing.T) {
	a := Cuboid{
		10, 13,
		10, 13,
		10, 14,
	}
	b := Cuboid{
		11, 14,
		12, 14,
		11, 15,
	}
	c := Cuboid{
		11, 13,
		12, 13,
		11, 14,
	}
	require.EqualValues(t, c, a.Intersect(b))
}

func TestCuboid_Intersect_NoIntersection(t *testing.T) {
	a := Cuboid{
		10, 13,
		10, 13,
		10, 14,
	}
	b := Cuboid{
		50, 54,
		12, 14,
		11, 15,
	}
	c := Cuboid{
		0, 0, 0, 0, 0, 0,
	}
	require.EqualValues(t, c, a.Intersect(b))
	require.False(t, a.Intersect(b).HasVolume())
	require.EqualValues(t, 0, a.Intersect(b).Volume())
}

func TestCuboid_Subtract1(t *testing.T) {
	b := Cuboid{
		11, 14,
		11, 14,
		11, 14,
	}
	a := Cuboid{
		10, 13,
		10, 13,
		10, 13,
	}

	require.EqualValues(t, 19, Volume(b.Subtract(a)))
}

func TestCuboid_Subtract2(t *testing.T) {
	c := Cuboid{
		9, 12,
		9, 12,
		9, 12,
	}
	b := Cuboid{
		11, 14,
		11, 14,
		11, 14,
	}
	a := Cuboid{
		10, 13,
		10, 13,
		10, 13,
	}

	require.EqualValues(t, 38,
		Volume(
			Subtract(
				append([]Cuboid{a}, b.Subtract(a)...),
				c,
			),
		),
	)
}

func TestCuboid_Subtract3(t *testing.T) {
	d := Cuboid{
		10, 11, 10, 11, 10, 11,
	}
	c := Cuboid{
		9, 12,
		9, 12,
		9, 12,
	}
	b := Cuboid{
		11, 14,
		11, 14,
		11, 14,
	}
	a := Cuboid{
		10, 13,
		10, 13,
		10, 13,
	}

	require.EqualValues(t, 39,
		Volume(
			append([]Cuboid{d},
				Subtract(
					append([]Cuboid{a}, b.Subtract(a)...),
					c,
				)...),
		),
	)
}
