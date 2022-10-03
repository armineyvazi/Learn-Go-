package main

import "fmt"

func main() {

	slice:=[]string{"hello","world","!"}

	for i,element:=range slice {
		fmt.Println(i,element,":")
		for _,ch:=range element {
			fmt.Printf("%q\n",ch)
		}
	}
	
	for i:=0;i<len(slice);i++{
		fmt.Println("slice",slice[i],":")
		char:=slice[i]
		for i:=0;i<len(char);i++{
			ch:=char[i]
			fmt.Printf("%q\n",ch)
		}
	}

}
