package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addhealth(amount uint) {
	if player.health+amount < player.maxHealth {
		player.health += amount
		fmt.Println(player.name, "Add amount", amount, "health->", player.health)
	} else {
		fmt.Println("Your Health is max")
	}

}
func (player *Player) addDamage(amount uint) {
	if player.health-amount < player.health {
		player.health -= amount
		fmt.Println(player.name, "damage", amount, "damage->", player.health)
	} else {
		player.health = 0
		fmt.Println(player.name, "dead:)")
	}
}
func (player *Player) lessEnergy(amount uint) {
	if player.energy-amount < player.energy {
		player.energy -= amount
		fmt.Print(player.name, "lessEnergy", amount, "essEnergy", player.energy)
	} else {
		player.energy = 0
		fmt.Println("Not Enogh Energy")
	}
}
func (player *Player) addEnergy(amount uint) {
	if player.energy+amount < player.maxEnergy {
		player.energy += amount
		fmt.Println(player.name, "Add amount", amount, "energy->", player.energy)
	} else {
		fmt.Println("Your Energy is max")
	}

}
func (player *Player) check(players Player) {
	fmt.Println()
	fmt.Println(players)
}
