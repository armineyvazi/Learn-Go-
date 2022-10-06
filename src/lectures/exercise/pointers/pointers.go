//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)

//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array

package main

import (
	"fmt"
)

//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

const (
	Activate = true
	Inactive = false
)

//  - Security tags have two states: active (true) and inactive (false)
type securityTag bool

//* Create a structure to store items and their security tag state
type Item struct {
	name string
	tag  securityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activate(tag *securityTag) {
	*tag = Activate
}

//* Create functions to activate and deactivate security tags using pointers
func deactivate(tag *securityTag) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {

	fmt.Println("Cheking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	shirt := Item{"Shirt", Activate}
	pants := Item{"Pants", Activate}
	purse := Item{"purse", Activate}
	watch := Item{"watch", Activate}

	//- Store them in a slice or array
	item := []Item{shirt, pants, purse, watch}
	fmt.Println("Iniatial", item)
	//- Deactivate any one security tag in the array/slice
	deactivate(&item[0].tag)
	fmt.Println("Item 0 is deactivate", item)
	//  - Call the checkout() function to deactivate all tags
	checkout(item)
	//  - Print out the array/slice after each change
	fmt.Println("check out", item)

}
