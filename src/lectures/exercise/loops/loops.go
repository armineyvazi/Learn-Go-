//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	//* Print integers 1 to 50, except:
	for i:=1;i<=50;i++{
		divisibilityBy3:=i%3==0
		divisibilityBy5:=i%5==0		
		if divisibilityBy3 && divisibilityBy5{
			//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("FizzBuzz:",i)
		}else if divisibilityBy3{
			//  - Print "Fizz" if the integer is divisible by 3
			fmt.Println(" divisible by 3 Fizz:",i)
		}else if divisibilityBy5{
			//  - Print "Buzz" if the integer is divisible by 5
			fmt.Println(" divisible by 5 Buzz:",i)
		}else{
			fmt.Println(i)
		}	
	}
}
