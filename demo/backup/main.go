package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})

	if err != nil {
		fmt.Println("listen failed: ", err)
		return
	}

	defer conn.Close()

	for {
		var data [2048]byte
		n, _, err := conn.ReadFromUDP(data[:])
		//n, _, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed: ", err)
			continue
		}
		//fmt.Printf("data: %v, addr: %v, count: %v\n ", string(data[:]), addr, n)

		//go handler(string(data[:]), n)
		go func() {
			var str string
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
			pack(&m, &str)
		}()
	}
}

func pack(m *map[string]interface{}, ret *string) {
	_, ok := (*m)["apid_0060"]
	if ok {
		pack_0060(m, ret)
		return
	}
	_, ok = (*m)["apid_0064"]
	if !ok {
		return
	}
	pack_0064(m, ret)
}

func pack_0060(m *map[string]interface{}, ret *string) {
	//	fmt.Println("this is a 0060 frame")
	var b string //bin
	var d int    //dec
	//var h string //hex
	var str string
	var n int
	var buf string

	var name string
	var length int

	//str = string((*m)["apid_0060"])

	// apid 16
	str = fmt.Sprintf("%v", (*m)["apid_0060"])
	d, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("err: ", err)
	}
	b = fmt.Sprintf("%b", d)
	//fmt.Printf("str: %v, dec: %v, bin: %b. hex: %x\n", str, d, d, d)
	//fmt.Printf("str: %v, dec: %v, bin: %v. hex: %x\n", str, d, b, d)
	n = 16 - len([]rune(b))
	if n > 0 {
		b = fmt.Sprintf("%v%b", strings.Repeat("0", n), d)
	} else {
		b = fmt.Sprintf("%b", d)
	}
	buf = fmt.Sprintf("%v", b)
	//fmt.Println("res: ", buf)

	// packetSeqId 2
	str = fmt.Sprintf("%v", (*m)["packetSeqId_0060"])
	d, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println("err: ", err)
	}
	b = fmt.Sprintf("%b", d)
	//fmt.Printf("str: %v, dec: %v, bin: %b. hex: %x\n", str, d, d, d)
	//fmt.Printf("str: %v, dec: %v, bin: %v, hex: %x\n", str, d, b, d)
	n = 2 - len([]rune(b))
	if n > 0 {
		b = fmt.Sprintf("%v%b", strings.Repeat("0", n), d)
	} else {
		b = fmt.Sprintf("%b", d)
	}
	buf = fmt.Sprintf("%v%v", buf, b)

	// packetCount 14
	str = fmt.Sprintf("%v", (*m)["packetCount_0060"])
	d, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println("err: ", err)
	}
	b = fmt.Sprintf("%b", d)
	//fmt.Printf("str: %v, dec: %v, bin: %b. hex: %x\n", str, d, d, d)
	//fmt.Printf("str: %v, dec: %v, bin: %v, hex: %x\n", str, d, b, d)
	n = 14 - len([]rune(b))
	if n > 0 {
		b = fmt.Sprintf("%v%b", strings.Repeat("0", n), d)
	} else {
		b = fmt.Sprintf("%b", d)
	}
	buf = fmt.Sprintf("%v%v", buf, b)

	// packetLength 16
	name = "packetLength_0060"
	length = 16
	str = fmt.Sprintf("%v", (*m)[name])
	d, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println("err: ", err)
	}
	b = fmt.Sprintf("%b", d)
	//fmt.Printf("str: %v, dec: %v, bin: %b. hex: %x\n", str, d, d, d)
	//fmt.Printf("str: %v, dec: %v, bin: %v, hex: %x\n", str, d, b, d)
	n = length - len([]rune(b))
	if n > 0 {
		b = fmt.Sprintf("%v%b", strings.Repeat("0", n), d)
	} else {
		b = fmt.Sprintf("%b", d)
	}
	buf = fmt.Sprintf("%v%v", buf, b)

	fmt.Println("res: ", buf)

	var t = make_item60()
	for _, item := range t {
		fmt.Println(item.name)
	}

}

func pack_0064(m *map[string]interface{}, ret *string) {
	//	fmt.Println("this is a 0064 frame")
}
