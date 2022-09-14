//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//


package main

import "fmt"

//--Requirements:
//* write a function a returns any to numbers
func twoTwos()(int,int){
	return 2,2
}
//* write a function that return any numbers
func five()	int	{
	return 5
}
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThree(a,b,c int)int{
	 return a+b+c
}

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greetPerson(name string){
	fmt.Println("My name is:",name)
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
func hiThere()string{
	return "Hi there"
}


func main() {
	greetPerson("Armin")
	fmt.Println(hiThere())

	a,b:=twoTwos()
	fmt.Println("A is :",a,"B is:",b)

	answer:=addThree(five(),a,b)
	fmt.Println("AddThree function five + a + b equal to:",answer)

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	//* Call every function at least once

}
