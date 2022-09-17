//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:

package main

import (
	"fmt"
)

//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array

type Products struct{
	price int
	name string 

}

//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items

func printState(list [4]Products){
	var cost,totalItemes int

	for i:=0 ; i < len(list); i++ {
		item	:=	list[i]
		cost	+=	item.price
		if item.name != "" {
			totalItemes	+=	1
		}
	}
	fmt.Println("Last item on the list:",list[totalItemes-1])
	fmt.Println("Total itemes",totalItemes)
	fmt.Println("Total cost",cost)
}


func main() {
	//* Add a fourth product to the list and print out the
	//  information again
	shippingList := [4]Products{
		{1,"food"},
		{2,"milk"},
		{3,"water"},
	}
	printState(shippingList)

	shippingList[3]=Products{2,"bread"}

	printState(shippingList)
}
