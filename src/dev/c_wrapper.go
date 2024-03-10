package main

/*
#cgo LDFLAGS: -lm -L${SRCDIR}/lib -L./lib -lsatellite
#cgo CFLAGS: -I ./
#include "satellite_code.h"
*/
import "C"

func satellite() {
	C.satellite()
}

func add(x float64, y float64) float64 {
	return float64(C.add(C.double(x), C.double(y)))
}

func test() float64 {
	return float64(C.test())
}

//#cgo LDFLAGS: -L${SRCDIR}/lib ./lib/libsatellite.so
