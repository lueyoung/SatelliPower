package main

import (
	"fmt"
	"net"
)

func main() {
	socket_src := get_src()
	src, err := net.ResolveUDPAddr("udp", socket_src)
	if err != nil {
		panic(err)
	}
	//recver, err := net.ListenUDP("udp", &net.UDPAddr{
	//	IP:   net.IPv4(127, 0, 0, 1),
	//	Port: 8080,
	//})

	recver, err := net.ListenUDP("udp", src)
	if err != nil {
		fmt.Println("recver failed: ", err)
		return
	}

	defer recver.Close()

	for {
		var data [2048]byte
		_, _, err := recver.ReadFromUDP(data[:])
		//n, _, err := recver.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed: ", err)
			continue
		}
		//fmt.Printf("data: %v, addr: %v, count: %v\n ", string(data[:]), addr, n)

		//go handler(string(data[:]), n)
		go func() {
			fmt.Printf("%v\n ", string(data[:]))
		}()
	}
}
