//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:

// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var myFavoriteColorName = "red"
	fmt.Println("My favorit color name is:",myFavoriteColorName)
	//* Store your birth year and age (in years) in two variables usings
	 birthYears,ageInYears :=2001,21
	 fmt.Println("birth years is:",birthYears,"Age years:",ageInYears)
	//  compound assignment
	var (
		firstInitial='A'
		lastInitial='I'
	)
	fmt.Println("First Initial :",firstInitial,"Last Initial:",lastInitial)

	//* Store your first & last initials in two variables using block assignment
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	var ageInDay int
	ageInDay = 365*ageInYears
	fmt.Println("I am :",ageInDay)
}
