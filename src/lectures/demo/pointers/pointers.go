package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}
func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {

	counter := Counter{}

	hello := "Hello"
	world := "world"
	replace(&hello, "Hi", &counter)

	fmt.Println(hello, world)

	fmt.Println("---------------")

	phrase:=[]string{hello,world}
	fmt.Println(phrase)
	
	fmt.Println("---------------")

	replace(&phrase[1],"Go",&counter)
	fmt.Println(phrase)

}
