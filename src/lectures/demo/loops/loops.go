package main

import "fmt"

func main() {
	sum:=0
	fmt.Println("Sum is",sum)
	//print 1,3,5,7,9
	for i:=1;i<10;i++{
		if i%2==0{
			continue
		}
		fmt.Println(i)
	}
	for i:=0;i<10;i++{
		sum+=i
		fmt.Println("sum is :",sum)
	}
	
	for sum>10{
		sum-=5
		fmt.Println("Decrement .sum is",sum)
	}
}
