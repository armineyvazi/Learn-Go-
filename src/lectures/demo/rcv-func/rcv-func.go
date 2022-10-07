package main

import "fmt"

type space struct {
	occupied bool
}
type ParkingLot struct {
	sapces []space
}

func occupySpace(lot *ParkingLot, spaceNumber int) {
	lot.sapces[spaceNumber-1].occupied = true
}
//receiver function
func (lot *ParkingLot)occupySpace(spaceNumber int){
	lot.sapces[spaceNumber-1].occupied = true
}
func (lot *ParkingLot)vacateSpace(spaceNumber int){
	lot.sapces[spaceNumber-1].occupied = false
}

func main() {
	lot:=ParkingLot{sapces:make([]space, 4)}
	fmt.Println("initial:",lot)
	lot.occupySpace(1)
	occupySpace(&lot,2)
	fmt.Println("after occpuied",lot)
	lot.vacateSpace(2)
	fmt.Println("after vaated:",lot)

	



	
}
