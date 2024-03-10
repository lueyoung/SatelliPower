package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Repeat("0", 5)
	fmt.Println(s)
	fmt.Println(len([]rune(s)))
}
