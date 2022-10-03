package main

import "fmt"

func main() {
	
	shoppingList:=make(map[string]int)
	shoppingList["eggs"]	=11
	shoppingList["milk"]	=1
	shoppingList["bread"]	+=1


	shoppingList["eggs"]	+=1
	fmt.Println(shoppingList)

	delete(shoppingList,"milk")
	fmt.Println("Milk deleted","new List:",shoppingList)

	fmt.Println("need",shoppingList["eggs"],"eggs")

	ceral,found:=shoppingList["ceral"]
	if !found{
		fmt.Println("nop")
	}else{
		fmt.Println("yup",ceral,"boxes")
	}
	
	totalItem:=0
	for _,amount:=range shoppingList{
		totalItem+=amount
	}
	fmt.Println("There are:",totalItem,"on the shopping list")

}
