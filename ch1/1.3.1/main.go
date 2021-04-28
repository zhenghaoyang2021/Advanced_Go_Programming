package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// 定义数组的几种方式
func varArr1() {
	var a [3]int
	var b = [...]int{2: 3, 1: 2}
	var c = [...]int{1, 2, 4: 5, 6}
	_, _, _ = a, b, c
}

// 遍历数组
func varArr2() {
	var a = [...]int{1, 2, 3}
	var b = &a
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])
	for i, v := range b {
		fmt.Println(i, v)
	}

	for i := range a {
		fmt.Printf("a[%d]： %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]： %d\n", i, v)
	}
	var c = [...]int{1, 2}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]： %d\n", i, c[i])
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}
}

// 定义不同类型的数组
func varString1() {
	// string
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}
	_, _, _ = s1, s2, s3

	//struct
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{image.Point{0, 0}, image.Point{1, 1}}
	_, _, _ = line1, line2, line3

	//func
	var decoder1 [2]func(reader io.Reader) (image.Image, error)
	var decoder2 = [...]func(reader io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	_, _ = decoder1, decoder2

	//interface
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}
	_, _ = unknown1, unknown2

	//chan
	var chanList = [2]chan int{}
	_ = chanList

	//nil
	var a [0]int
	var b = [0]int{}
	var c = [...]int{}
	_, _, _ = a, b, c

	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1 //阻塞 等待传值

	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{} //struct{} 类型 {}结构体值
	}()
	<-c2

	var bb = [...]int{2: 3, 1: 2}
	// %T 打印类型  %#v 详细
	fmt.Printf("b: %T\n", bb)  //[3]int
	fmt.Printf("b: %#v\n", bb) //[3]int{0, 2, 3}
}

func main() {
	varArr1()
	varArr2()
	varString1()
}
