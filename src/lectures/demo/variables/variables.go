package main

import "fmt"

func main() {

	var myName = "Armin"

	fmt.Println("My name is:", myName)
	//-------------------------------//
	var name string = "Armin"

	fmt.Println("Name is :",name)
	//-------------------------------//
	userName := "Admin" 
	
	fmt.Println("User name :",userName)
	//-------------------------------//
	var sum int 

	fmt.Println("Sum is:",sum)
	//-------------------------------//
	part1,other:=1,5

	fmt.Println("Part is:",part1,"other is:",other)
	//-------------------------------//
	part2,other:=1,0
	fmt.Println("Part is:",part2,"other is:",other)

	//-------------------------------//
	sum=part1+part2
	fmt.Println("sum=",sum)
	//-------------------------------//
	var(
		lessonName="Variables"
		lessonType="Demo"
	)
	fmt.Println("Lesson Name :",lessonName)
	fmt.Println("Lesson tyoe :",lessonType)
	//-------------------------------//

	word1,word2,_:="hello","word","!"
	fmt.Println(word1,word2)
	
	var word3 int=2
	
	fmt.Println(word3)

}
