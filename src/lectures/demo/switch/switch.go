package main

import (
	"fmt"
	//"golang.org/x/text/cases"
)

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	
	switch  p:=price();{
	case p<2:
		fmt.Println("Cheap team")
	case p<10:
		fmt.Println("Moderately Price item")
	default:
		fmt.Println("Expensive item")
	}

	ticket:=Economy

	switch ticket{
		case Economy:
			fmt.Println("Economy seating")
		case Business:
			fmt.Println("Business seeting")
		case FirstClass:
			fmt.Println("FirstClass seeting")
		default:
			fmt.Println("Other seeting")
	}

}
