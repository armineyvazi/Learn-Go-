//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
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
	if player.energy+amount <= player.maxEnergy {
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
func main() {

	player1 := Player{
		name: "Armin",
		health: 33,
		maxHealth: 100,
		energy: 99,
		maxEnergy: 100,
	}
	fmt.Println("befor move anything", player1)
	player1.addhealth(10)
	player1.addhealth(10)
	player1.check(player1)
	player1.addDamage(50)
	player1.addEnergy(1)
	player1.check(player1)
	player1.lessEnergy(99)
	player1.check(player1)

}
