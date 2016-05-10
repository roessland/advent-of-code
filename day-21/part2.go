package main

import "fmt"

type Equip struct {
	Cost, Damage, Armor int
}

type Character struct {
	Cost, HP, Damage, Armor int
}

func (c *Character) Equip(equips []Equip) {
	c.Cost = 0
	c.Damage = 0
	c.Armor = 0
	for _, e := range equips {
		c.Cost += e.Cost
		c.Damage += e.Damage
		c.Armor += e.Armor
	}
}

func (c Character) WinsAgainst(b Character) bool {
	for {
		// character attacks boss
		dmg := c.Damage - b.Armor
		if dmg < 1 {
			dmg = 1
		}
		b.HP -= dmg
		if b.HP <= 0 {
			return true
		}

		// boss attacks character
		dmg = b.Damage - c.Armor
		if dmg < 1 {
			dmg = 1
		}
		c.HP -= dmg
		if c.HP <= 0 {
			return false
		}
	}
}

func main() {
	weapons := []Equip{
		Equip{8, 4, 0},
		Equip{10, 5, 0},
		Equip{25, 6, 0},
		Equip{40, 7, 0},
		Equip{74, 8, 0},
	}
	armors := []Equip{
		Equip{13, 0, 1},
		Equip{31, 0, 2},
		Equip{53, 0, 3},
		Equip{75, 0, 4},
		Equip{102, 0, 5},
	}
	rings := []Equip{
		Equip{25, 1, 0},
		Equip{50, 2, 0},
		Equip{100, 3, 0},
		Equip{20, 0, 1},
		Equip{40, 0, 2},
		Equip{80, 0, 3},
	}
	var player Character
	boss := Character{0, 100, 8, 2}

	maxGold := 0
	for _, wpn := range weapons {
		for _, usearmor := range []bool{true, false} {
			for _, armor := range armors {
				for numrings := range []int{0, 1, 2} {
					for ring1i := range rings {
						for ring2i := range rings {

							// Weapon
							equips := []Equip{}
							equips = append(equips, wpn)

							// Armor
							if usearmor {
								equips = append(equips, armor)
							}

							// Rings
							if numrings == 1 {
								equips = append(equips, rings[ring1i])
							} else if numrings == 2 {
								if ring1i == ring2i {
									continue
								}
								equips = append(equips, rings[ring1i], rings[ring2i])
							}

							player.Equip(equips)
							player.HP = 100
							if !player.WinsAgainst(boss) {
								if player.Cost > maxGold {
									maxGold = player.Cost
									fmt.Printf("%v\n", maxGold)
									fmt.Printf("	Equps: %v\n", equips)
								}
							}
						}
					}

				}
			}
		}
	}
}

// 8 too low
