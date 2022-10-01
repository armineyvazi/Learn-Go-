//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string
//* Create a function to print out the contents of the assembly line
func showLine(line []Part) {
	for i:=0;i<len(line); i++{
		part:=line[i]
		fmt.Println(part)
	}
}
func main() {
	//* Using a slice, create an assembly line that contains type Part
	assemblyLine:=make([]Part, 3)
	//- Create an assembly line having any three parts
	assemblyLine[0]="pipe"
	assemblyLine[1]="Bolt"
	assemblyLine[2]="Sheet"

	fmt.Println("3 Parts")
	showLine(assemblyLine)
	//  - Add two new parts to the line
	assemblyLine =  append(assemblyLine, "washer","wheel")
	fmt.Println("\nAdded two parts:")
	showLine(assemblyLine)
	
	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine=assemblyLine[2:]
	fmt.Println("\nSlice:")
	showLine(assemblyLine)
}
