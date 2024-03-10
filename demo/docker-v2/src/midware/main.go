package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
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

	socket_dest := get_dest()
	dest, err := net.ResolveUDPAddr("udp", socket_dest)
	if err != nil {
		panic(err)
	}
	//sender, err := net.DialUDP("udp", nil, &net.UDPAddr{
	//	IP:   net.IPv4(127, 0, 0, 1),
	//	Port: 8081,
	//})
	sender, err := net.DialUDP("udp", nil, dest)

	if err != nil {
		fmt.Println("sender failed: ", err)
		return
	}

	defer sender.Close()

	for {
		var data [2048]byte
		n, _, err := recver.ReadFromUDP(data[:])
		//n, _, err := recver.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed: ", err)
			continue
		}
		//fmt.Printf("data: %v, addr: %v, count: %v\n ", string(data[:]), addr, n)

		//go handler(string(data[:]), n)
		go func() {
			m := make(map[string]interface{})
			err := json.Unmarshal(data[:n], &m)
			if err != nil {
				fmt.Println("read json failed: ", err)
				return
			}
			fmt.Printf("recv frame no. : %v\n", m["ID"])
			apid, ok := m["apid_0060"]
			if !ok {
				apid, ok = m["apid_0064"]
			}
			if ok == false {
				fmt.Println("wrong format")
				return
			}
			fmt.Printf("apid: %v\n", apid)
			msg := pack(&m)
			fmt.Printf("msg: %s\n", msg)
			sender.Write([]byte(msg))
			//fmt.Println("send to : ", get_dest())
		}()
	}
}

func pack(m *map[string]interface{}) string {
	_, ok := (*m)["apid_0060"]
	if ok {
		return pack_0060(m)
	}
	_, ok = (*m)["apid_0064"]
	if !ok {
		return "error"
	}
	return pack_0064(m)
}

func pack_0060(m *map[string]interface{}) string {
	//	fmt.Println("this is a 0060 frame")
	var buf string
	buf = ""

	var items = make_item60()
	for _, item := range items {
		seg := add(m, item)
		buf = fmt.Sprintf("%v%v", buf, seg)
	}
	//fmt.Println(len([]rune(buf)))
	h := conv_bin_to_hex(buf)
	fmt.Println(h)
	return h
}

func pack_0064(m *map[string]interface{}) string {
	//	fmt.Println("this is a 0064 frame")
	return "test64"
}

func add(m *map[string]interface{}, item *Item) string {
	// packetLength 16
	name := item.name
	length := item.len
	str := fmt.Sprintf("%v", (*m)[name])
	// dec
	d, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("err: ", err)
	}
	// bin
	b := fmt.Sprintf("%b", d)
	//fmt.Printf("str: %v, dec: %v, bin: %b. hex: %x\n", str, d, d, d)
	//fmt.Printf("str: %v, dec: %v, bin: %v, hex: %x\n", str, d, b, d)
	n := length - len([]rune(b))
	if n > 0 {
		b = fmt.Sprintf("%v%b", strings.Repeat("0", n), d)
	} else {
		b = fmt.Sprintf("%b", d)
	}
	return b
}
