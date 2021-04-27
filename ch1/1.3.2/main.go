package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// 字符串遍历与转换
func varString() {
	// 字符串 底层数据为byte数组
	var data = [...]byte{
		'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd',
	}
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	_, _, _ = data, hello, world
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println("len(s)=:", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s1)=:", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)
	fmt.Println("len(s2)=:", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len)

	fmt.Printf("%#v\n", []byte("hello,世界"))

	fmt.Println("\xe4\xb8\x96")
	fmt.Println("\xe7\x95\x8c")
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc")

	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	const ss = "\xe4\x00\x00\xe7\x95\x8cabc"
	for i := 0; i < len(ss); i++ {
		fmt.Printf("%d %x\n", i, ss[i])
	}

	fmt.Printf("%#v\n", []rune("世界"))
	fmt.Printf("%#v\n", string([]rune{'世', '界'}))
}

//for range 对 string的迭代模拟
func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; i < len(s); i++ {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

//[]byte(s) 转换模拟
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

//string(bytes)转换模拟
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	return p
}

// []rune(s) 转换模拟
func str2runes(s []byte) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

// string(runes) 转换模拟
func runes2string(s []rune) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

func main() {
	//varString()
	for i, i2 := range str2bytes("hello") {
		fmt.Printf("%d %#v %v %d %c\n", i, i2, i2, i2, i2)
	}
	fmt.Println([]byte("hello"))
	fmt.Println(bytes2str([]byte("hello")))
}
