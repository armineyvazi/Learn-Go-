package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chiken string
type Salad string

func (c Chiken) PrepareDish() {
	fmt.Println("cook chicken")
}
func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}
func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish:%v", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}
func main() {
	dishes := []Preparer{Chiken("Chicken wings"), Salad("Salad")}
	prepareDishes(dishes)

}
