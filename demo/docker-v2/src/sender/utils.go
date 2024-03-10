package main

import (
	"fmt"
	"strconv"
)

func conv_bin_to_hex(b string) string {
	l := len([]rune(b))
	r := l / 4
	m := l % 4
	var buf string
	//fmt.Println(l, r, m)

	if m > 0 {

		seg := b[0:m]
		//fmt.Printf("%v\n", seg)
		dec, err := strconv.ParseInt(seg, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Printf("%d\n", dec)
		buf = fmt.Sprintf("%v", dec)
		//fmt.Println(buf)

	}

	for i := 0; i < r; i++ {
		//fmt.Println(i*4+m, (i+1)*4+m)
		seg := b[i*4+m : (i+1)*4+m]
		//fmt.Printf("%v\n", seg)
		dec, err := strconv.ParseInt(seg, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Printf("%d\n", dec)
		buf = fmt.Sprintf("%v%x", buf, dec)
		//fmt.Println(buf)
	}
	return buf
}

func conv_bin_to_hex2(b string) string {
	l := len([]rune(b))
	r := l / 4
	//m := l % 4
	var buf string
	for i := 0; i <= r; i++ {
		end := l - 1 - i*4
		start := end
		if start <= 0 {
			break
		}
		for j := 0; j < 3; j++ {
			start--
			if start == 0 {
				break
			}
		}
		//fmt.Printf("%q\n", b[start:end+1])
		seg := b[start : end+1]
		//fmt.Printf("%v\n", seg)
		dec, err := strconv.ParseInt(seg, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Printf("%d\n", dec)
		buf = fmt.Sprintf("%x%v", dec, buf)
		//fmt.Println(buf)
	}
	return buf
}
