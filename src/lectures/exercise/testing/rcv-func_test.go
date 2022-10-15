//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		energy:    100,
		maxHealth: 100,
		maxEnergy: 100,
	}
}
func TestAddHealth(t *testing.T) {
	player := newPlayer()
	player.addhealth(100)
	if player.health > player.maxHealth {
		t.Fatalf("health can not more than maxhealth:%v and health:%v", player.maxHealth, player.health)
	}

}
func TestAddEnrgy(t *testing.T) {
	player := newPlayer()
	player.addEnergy(100)
	if player.energy > player.maxEnergy {
		t.Errorf("enrgy can not more than maxenergy:%v and energy:%v", player.maxEnergy, player.energy)
	}
}

func TestDamageHealth(t *testing.T) {
	player := newPlayer()
	player.addDamage(10)
	if player.health < 0 {
		t.Errorf("health can not less 0 health is :%v and damage value is%v", player.health, 100)

	}
}
func TestLessEnergy(t *testing.T) {
	player := newPlayer()
	player.lessEnergy(100)
	if player.energy < 0 {
		t.Errorf("energy can not less 0 energy is:%v", player.energy)
	}

}
