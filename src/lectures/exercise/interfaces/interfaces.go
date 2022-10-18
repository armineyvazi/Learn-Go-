//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
const (
	small = iota
	standard
	large
)
type Lift int

type park interface {
	autopark() Lift
}
type Motorcycles string
type Cars string
type Trucks string

//* Write a single function to handle all of the vehicles
func (m Motorcycles) String() string {
	return fmt.Sprintf("Motorcycles:%v", string(m))
}
func (c Cars) String() string {
	return fmt.Sprintf("Cars:%v", string(c))

}
func (t Trucks) String() string {
	return fmt.Sprintf("Trucks:%v", string(t))
}
//* The shop has lifts for multiple vehicle sizes/types:
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
func (m Motorcycles) autopark() Lift {
	return small
}
func (c Cars) autopark() Lift {
	return standard
}
func (t Trucks) autopark() Lift {
	return large
}
func autolift(l park) {
	switch l.autopark() {
		case small:
			fmt.Printf("send %v to small lift\n",l)
			break
		case standard:
			fmt.Printf("send %v to standard lift\n",l)
			break
		case large:
			fmt.Printf("send %v to large lift\n",l)
			break
	}
}

func main() {
	motor 	:= Motorcycles("ktm")
	cars 	:= Cars("sporto")
	trucks 	:= Trucks("trucks")

	autolift(motor)
	autolift(cars)
	autolift(trucks)
}
