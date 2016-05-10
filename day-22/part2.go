package main

import "math/rand"
import "fmt"

var spells = [...]Spell{
	Spell{
		Name:   "Magic Missile",
		Cost:   53,
		Damage: 4,
	},
	Spell{
		Name:   "Drain",
		Cost:   73,
		Damage: 2,
		Heal:   2,
	},
	Spell{
		Name: "Shield",
		Cost: 113,
		Effect: Effect{
			Name:      "Shield",
			TurnsLeft: 6,
			Shield:    7,
		},
	},
	Spell{
		Name: "Poison",
		Cost: 173,
		Effect: Effect{
			Name:      "Poison",
			TurnsLeft: 6,
			Damage:    3,
		},
	},
	Spell{
		Name: "Recharge",
		Cost: 229,
		Effect: Effect{
			Name:      "Recharge",
			TurnsLeft: 5,
			Recharge:  101,
		},
	},
}

type Spell struct {
	Name   string
	Cost   int
	Damage int
	Heal   int
	Effect Effect
}

type Effect struct {
	Name      string
	TurnsLeft int
	Damage    int
	Heal      int
	Recharge  int
	Shield    int
}

type Character struct {
	HP, Mana, Armor, Damage int
	Type                    string
}

func RandomSpell() Spell {
	return spells[rand.Intn(len(spells))]
}

func Log(msg string, params ...interface{}) {
	if false {
		fmt.Printf(msg, params...)
	}
}

func Battle() (bool, int) {
	manaSpent := 0
	boss := Character{HP: 51, Damage: 9, Type: "physical"}
	player := Character{HP: 50, Mana: 500, Type: "magic"}
	playerEffects := map[string]*Effect{}
	for {
		Log("-- Player turn --\n")
		Log("- Player has %v hit points, %v armor, %v mana\n",
			player.HP, player.Armor, player.Mana)
		Log("- Boss has %v hit points\n", boss.HP)

		player.HP--
		if player.HP <= 0 {
			Log("- Player died because of hard mode.\n")
			return false, manaSpent
		}
		// Apply effects
		for name, _ := range playerEffects {
			playerEffects[name].TurnsLeft--
			Log("- Applies %v; its timer is now %v.\n", name, playerEffects[name].TurnsLeft)
			boss.HP -= playerEffects[name].Damage
			player.HP += playerEffects[name].Heal
			player.Mana += playerEffects[name].Recharge
			if playerEffects[name].TurnsLeft == 0 {
				delete(playerEffects, name)
			}
		}
		// effects can kill the boss
		if boss.HP <= 0 {
			Log("- Boss killed by effects.\n")
			return true, manaSpent
		}

		// Choose a random available spell which the player has enough mana for
		availableSpellCount := 0
		for _, spell := range spells {
			if _, busy := playerEffects[spell.Effect.Name]; spell.Cost <= player.Mana && !busy {
				availableSpellCount++
			}
		}
		if availableSpellCount == 0 {
			Log("Player does not have enough mana. Annihilated.\n")
			return false, manaSpent
		}
		var spell Spell
		for {
			spell = RandomSpell()
			if _, busy := playerEffects[spell.Effect.Name]; spell.Cost <= player.Mana && !busy {
				break
			}
		}

		// Apply damage and heals and effects
		Log("- Player casts %v\n", spell.Name)
		boss.HP -= spell.Damage
		player.HP += spell.Heal
		if len(spell.Effect.Name) != 0 {
			// Dereference to get a copy of the effect, not just a reference
			playerEffects[spell.Effect.Name] = &spell.Effect
		}
		player.Mana -= spell.Cost
		manaSpent += spell.Cost

		Log("-- Boss turn --\n")
		Log("- Player has %v hit points, %v armor, %v mana\n",
			player.HP, player.Armor, player.Mana)
		Log("- Boss has %v hit points\n", boss.HP)

		// Apply effects
		for name, _ := range playerEffects {
			playerEffects[name].TurnsLeft--
			Log("- Applies %v; its timer is now %v.\n", name, playerEffects[name].TurnsLeft)
			boss.HP -= playerEffects[name].Damage
			player.HP += playerEffects[name].Heal
			player.Mana += playerEffects[name].Recharge
			if playerEffects[name].TurnsLeft == 0 {
				delete(playerEffects, name)
			}
		}

		if boss.HP <= 0 {
			Log("Player wins\n")
			return true, manaSpent
		}

		shield := 0
		if eff, ok := playerEffects["Shield"]; ok {
			shield = eff.Shield
		}
		dmg := boss.Damage - shield
		if dmg < 1 {
			dmg = 1
		}
		player.HP -= dmg
		if player.HP <= 0 {
			Log("Player died.\n")
			return false, manaSpent
		}
	}
}

func main() {
	minSpent := 99999999
	for {
		//Battle()
		//break
		win, spent := Battle()
		if win && spent < minSpent {
			minSpent = spent
			fmt.Printf("Spent: %v\n", minSpent)
		}
	}
}
