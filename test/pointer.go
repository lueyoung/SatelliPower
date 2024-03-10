package main

import "fmt"

func main() {
	var x *int
	x = new(int)
	*x = 1
	fmt.Println(*x)
}
