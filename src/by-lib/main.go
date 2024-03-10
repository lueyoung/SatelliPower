package main

/*
#cgo LDFLAGS: -L${SRCDIR}/lib ./lib/libsatellite.so
#cgo CFLAGS: -I ./
#include "satellite_code.h"
*/
import "C"

import "fmt"

func main() {
	C.satellite()
	fmt.Println("Hello")
	//var x, y float64
	//x = 0.1
	//y = 0.2
	//var z = C.add(x, y)
	//fmt.Println(z)
	x := C.test()
	fmt.Printf("%T\n", x)
	a := C.double(0.01)
	b := C.double(0.1)
	c := C.add(a, b)
	fmt.Println(c)

}
