package day06

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Vec struct {
	X, Y int
}

type VecPair struct {
	Pos, Dir Vec
}

func (v Vec) TurnRight() Vec {
	switch v {
	case Vec{1, 0}:
		return Vec{0, 1}
	case Vec{0, -1}:
		return Vec{1, 0}
	case Vec{-1, 0}:
		return Vec{0, -1}
	case Vec{0, 1}:
		return Vec{-1, 0}
	}
	panic("invalid direction")
}

func (v Vec) Add(dir Vec) Vec {
	return Vec{v.X + dir.X, v.Y + dir.Y}
}

func ReadInput(inputName string) ([][]byte, Vec, Vec) {
	m := aocutil.FSReadLinesAsBytes(Input, inputName)
	for y := range m {
		for x := range m[y] {
			if m[y][x] == '^' {
				m[y][x] = '.'
				return m, Vec{x, y}, Vec{0, -1}
			}
		}
	}
	panic("no ^ found")
}

func Next(m [][]byte, currPos, currDir, obstacle Vec) (nextPos, nextDir Vec, hitWall bool) {
	nextPos = currPos.Add(currDir)
	if IsOutOfBounds(m, nextPos) {
		return nextPos, currDir, false
	}
	if m[nextPos.Y][nextPos.X] == '#' || nextPos == obstacle {
		return currPos, currDir.TurnRight(), true
	}
	return nextPos, currDir, false
}

func IsOutOfBounds(m [][]byte, pos Vec) bool {
	return pos.X < 0 || pos.Y < 0 || pos.Y >= len(m) || pos.X >= len(m[pos.Y])
}

func Part12(inputName string) (int, int) {
	visitedWithoutObstacle := make(map[Vec]bool)
	m, startPos, startDir := ReadInput(inputName)

	currPos, currDir := startPos, startDir
	for !IsOutOfBounds(m, currPos) {
		visitedWithoutObstacle[currPos] = true
		currPos, currDir, _ = Next(m, currPos, currDir, Vec{-1, -1})
	}

	possiblePositions := 0
	for obstaclePosition := range visitedWithoutObstacle {
		currPos, currDir := startPos, startDir
		visited := map[VecPair]bool{}

		// Optimization: Only store positions where we crashed into a wall.
		// Improves runtime by ~10x.
		hitWall := false
		for !IsOutOfBounds(m, currPos) {
			if hitWall {
				if visited[VecPair{currPos, currDir}] {
					possiblePositions++
					break
				}
				visited[VecPair{currPos, currDir}] = true
			}
			currPos, currDir, hitWall = Next(m, currPos, currDir, obstaclePosition)
		}
	}

	return len(visitedWithoutObstacle), possiblePositions
}
