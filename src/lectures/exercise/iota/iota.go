//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import (
	"fmt"
)
const 
	(
		Add = iota
		Subtract
		Multiply
		Divide
	)
type Operations int 

func (op Operations)calculate(lhs,rhs int) int{

	switch op{
		case Add:
			return lhs+rhs
		case Subtract:
			return lhs-rhs
		case Multiply:
			return lhs*rhs
		case Divide:
			return lhs/rhs
	}
	panic("Unhandle operation")
}

func main() {
//* The existing function calls in main() represent the API and cannot be changed
	add:=Operations(Add)
	sub:=Operations(Subtract)
	mul:=Operations(Multiply)
	div:=Operations(Divide)
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50

	
}
