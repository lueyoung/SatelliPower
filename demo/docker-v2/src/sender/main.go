package main

import (
	//"encoding/json"
	"fmt"
	//"net"
	//"strconv"
	//"strings"
)

func main() {
	m := make_item60()
	m[0].value = "hi"
	fmt.Println(m[0].key, m[0].value, m[0].len)
}
