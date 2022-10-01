package main

import "fmt"
func printSlice(title string,slice []string){
	fmt.Println()
	fmt.Println("----",title,"-----")
	for i:=0;i<len(slice);i++{
		element:=slice[i]
		fmt.Println(element)
	}
}

func main() {

	route:=[]string{"Georcy","Department Story","Salon"}
    
	printSlice("armin",route)

	route = append(route, "Home")
	printSlice("Route 2",route)

	fmt.Println()

	fmt.Println(route[0],"visited")
	fmt.Println(route[1],"visited")
	fmt.Println(route[2],"visited")

	route = route[2:]

	printSlice("Remaining location",route)

}
