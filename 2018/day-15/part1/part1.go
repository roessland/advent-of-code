package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type TileType rune
type UnitType rune

const (
	Wall   TileType = '#'
	Open   TileType = '.'
	Elf    UnitType = 'E'
	Goblin UnitType = 'G'
)

type Unit struct {
	Type        UnitType
	HitPoints   int
	AttackPower int
	Pos         Pos
}

func (u *Unit) IsAlive() bool {
	return u.HitPoints > 0
}

func (u *Unit) IsDead() bool {
	return u.HitPoints <= 0
}

func (u *Unit) String() string {
	return fmt.Sprintf("%c(%d)", u.Type, u.HitPoints)
}

type Tile struct {
	Type TileType
	Unit *Unit
}

type Map [][]*Tile

func (m Map) DeepCloneWithElfAttackPower(elfAttackPower int) Map {
	c := make(Map, len(m))
	for y := 0; y < len(m); y++ {
		c[y] = make([]*Tile, len(m[y]))
		for x := 0; x < len(m[y]); x++ {
			c[y][x] = &Tile{Type: m[y][x].Type}
			if m[y][x].Unit != nil {
				c[y][x].Unit = &Unit{
					Type:      m[y][x].Unit.Type,
					HitPoints: m[y][x].Unit.HitPoints,
					Pos:       m[y][x].Unit.Pos,
				}
				if c[y][x].Unit.Type == Goblin {
					c[y][x].Unit.AttackPower = 3
				} else if c[y][x].Unit.Type == Elf {
					c[y][x].Unit.AttackPower = elfAttackPower
				}
			}
		}
	}
	return c
}

func (m Map) Print() {
	fmt.Print(m)
}

func (m Map) String() string {
	var str strings.Builder
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			var s string
			t := m[y][x]
			if t.Unit == nil {
				s = fmt.Sprintf("%c", t.Type)
			} else {
				s = fmt.Sprintf("%c", t.Unit.Type)
			}
			str.WriteString(s)
		}
		str.WriteString("\n")
	}
	return str.String()
}

func (m Map) Iterate(do func(*Tile)) {
	for _, row := range m {
		for _, tile := range row {
			do(tile)
		}
	}
}

func (m Map) IterateUnits(do func(*Unit)) {
	m.Iterate(func(t *Tile) {
		if t.Unit != nil {
			do(t.Unit)
		}
	})
}

func (m Map) AliveUnits() []*Unit {
	var aliveUnits []*Unit
	m.IterateUnits(func(unit *Unit) {
		if unit.IsAlive() {
			aliveUnits = append(aliveUnits, unit)
		} else {
			panic("dead unit should have been removed")
		}
	})
	return aliveUnits
}

func (m Map) UnitAt(p Pos) *Unit {
	return m[p.Y][p.X].Unit
}

func (m Map) TileTypeAt(p Pos) TileType {
	return m[p.Y][p.X].Type
}

func EnemyOf(unitType UnitType) UnitType {
	switch unitType {
	case Goblin:
		return Elf
	case Elf:
		return Goblin
	default:
		panic("unknown unit type")
	}
}

func (m Map) IsInRangeOfEnemy(attacker *Unit) bool {
	return m.GetEnemyInRange(attacker) != nil
}

func (m Map) GetEnemyInRange(attacker *Unit) *Unit {
	p := attacker.Pos
	for _, adj := range p.Adjacent() {
		unit := m.UnitAt(adj)
		if unit == nil || unit.Type == attacker.Type || unit.IsDead() {
			continue
		}
		return unit
	}
	return nil
}

func (m Map) FindUnitsOfType(unitType UnitType) []*Unit {
	var units []*Unit
	m.IterateUnits(func(u *Unit) {
		if u.Type == unitType {
			units = append(units, u)
		}
	})
	return units
}

func UnitPositions(units []*Unit) []Pos {
	var ps []Pos
	for _, unit := range units {
		ps = append(ps, unit.Pos)
	}
	return ps
}

func AdjacentPositions(ps []Pos) []Pos {
	alreadyAdded := make(map[Pos]bool)
	var adjs []Pos
	for _, p := range ps {
		for _, adj := range p.Adjacent() {
			if !alreadyAdded[adj] {
				adjs = append(adjs, adj)
			}
		}
	}
	return adjs
}

func FilterPositions(ps []Pos, f func(Pos) bool) []Pos {
	var filtered []Pos
	for _, p := range ps {
		if f(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// DoMove does a move and returns false if combat has ended.
func (m Map) DoMove(attacker *Unit) bool {
	enemyType := EnemyOf(attacker.Type)
	targets := m.FindUnitsOfType(enemyType)
	if len(targets) == 0 {
		return false
	}
	targetPositions := UnitPositions(targets)
	targetAdjacentPositions := AdjacentPositions(targetPositions)
	targetAdjacentOpenPositions := FilterPositions(targetAdjacentPositions, func(p Pos) bool {
		return m.TileTypeAt(p) == Open && m.UnitAt(p) == nil
	})
	if len(targetAdjacentOpenPositions) == 0 {
		return true
	}
	// Find distance and direction to move
	distMap := m.Diffusion(targetAdjacentOpenPositions)
	nextPositions := FilterPositions(attacker.Pos.Adjacent(), func(p Pos) bool {
		return m.TileTypeAt(p) == Open
	})
	sort.Slice(nextPositions, func(i, j int) bool {
		pA, pB := nextPositions[i], nextPositions[j]
		dA, dB := distMap[pA.Y][pA.X], distMap[pB.Y][pB.X]
		if dA != dB {
			return dA < dB
		} else if pA.Y != pB.Y {
			return pA.Y < pB.Y
		} else {
			return pA.X < pB.X
		}
	})
	if len(nextPositions) == 0 {
		return true
	}
	if distMap[nextPositions[0].Y][nextPositions[0].X] == math.MaxInt32 {
		return true
	}
	m.DoStep(attacker, nextPositions[0])

	return true
}

// DoTurn does a turn and returns false if combat has ended
func (m Map) DoTurn(attacker *Unit) (bool, int) {
	if attacker.IsDead() {
		return true, 0
	}
	if !m.IsInRangeOfEnemy(attacker) {
		if !m.DoMove(attacker) {
			return false, 0
		}
	}
	return true, m.Attack(attacker)
}

// DoRound does a round and returns false if combat has ended
func (m Map) DoRound() (bool, int) {
	elvesLostTotal := 0
	for _, attacker := range m.AliveUnits() {
		moreRounds, elvesKilled := m.DoTurn(attacker)
		elvesLostTotal += elvesKilled
		if !moreRounds {
			return false, elvesLostTotal
		}
	}
	return true, elvesLostTotal
}

func (m Map) DoStep(attacker *Unit, pos Pos) {
	m[attacker.Pos.Y][attacker.Pos.X].Unit = nil
	m[pos.Y][pos.X].Unit = attacker
	attacker.Pos = pos
}

func (m Map) Attack(attacker *Unit) (elvesLost int) {
	// Find direction to attack
	targetPositions := FilterPositions(attacker.Pos.Adjacent(), func(p Pos) bool {
		return m.UnitAt(p) != nil && m.UnitAt(p).Type == EnemyOf(attacker.Type)
	})
	sort.Slice(targetPositions, func(i, j int) bool {
		pA, pB := targetPositions[i], targetPositions[j]
		hpA, hpB := m.UnitAt(pA).HitPoints, m.UnitAt(pB).HitPoints
		if hpA != hpB {
			return hpA < hpB
		} else if pA.Y != pB.Y {
			return pA.Y < pB.Y
		} else {
			return pA.X < pB.X
		}
	})
	if len(targetPositions) == 0 {
		return 0
	}
	targetPos := targetPositions[0]
	target := m.UnitAt(targetPos)
	target.HitPoints -= attacker.AttackPower
	if target.HitPoints <= 0 {
		if target.Type == Elf {
			elvesLost = 1
		}
		m[targetPos.Y][targetPos.X].Unit = nil
	}
	return elvesLost
}

func SimulateCombat(m Map) (roundsCompleted, totalElfLosses int) {
	for {
		onemore, elfLosses := m.DoRound()
		totalElfLosses += elfLosses
		if !onemore {
			break
		}
		roundsCompleted++
	}

	return roundsCompleted, totalElfLosses
}

func (m Map) Outcome() (outcome int, elfLosses int) {
	roundsCompleted, elfLosses := SimulateCombat(m)
	totalHP := 0
	m.IterateUnits(func(u *Unit) {
		totalHP += u.HitPoints
	})
	return totalHP * roundsCompleted, elfLosses
}

func part1(m Map) int {
	outcome, _ := m.DeepCloneWithElfAttackPower(3).Outcome()
	return outcome
}

func part2(m Map) int {
	for attackPower := 3; ; attackPower++ {
		c := m.DeepCloneWithElfAttackPower(attackPower)
		outcome, elfLosses := c.Outcome()
		if elfLosses == 0 {
			return outcome
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	m := ReadMap(f)

	fmt.Println("Part 1:", part1(m))
	fmt.Println("Part 2:", part2(m))
}

// 54112 too high
// 52688
