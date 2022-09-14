package main

import "fmt"

func double(x int )int{
	return x+x 
}

func add(lhs , rhs int )int{
	return lhs+rhs
}

func great() {
	fmt.Println("Hello from the great function!")
}

func main() {

	great()

	dozen:=double(2)
	fmt.Println("Double function:",dozen)

	bakersDozen:=add(dozen,1)
	fmt.Println("A baker's dozen is :",bakersDozen)

	anotherBakerDozen:=add(double(10),1)
	fmt.Println("Another baker's doze is::",anotherBakerDozen)


	fmt.Println("for test",add(double(10),1),add(2,2))


}
