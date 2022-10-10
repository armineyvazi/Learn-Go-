package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8, 9, 10}
	all := append(a, b...)
	answer := sum(all...)
	fmt.Println("all answer a and b", answer)
	//or use
	answer = sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(answer)
}
