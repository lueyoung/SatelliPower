package main

import "fmt"

func main() {
	satellite()
	fmt.Println("Hello")
	//var x, y float64
	//x = 0.1
	//y = 0.2
	//var z = C.add(x, y)
	//fmt.Println(z)
	x := test()
	fmt.Printf("%T\n", x)
	a := 0.01
	b := 0.1
	c := add(a, b)
	fmt.Println(c)

}
