package main

import "fmt"

func varArr1() {
	var a [3]int
	var b = [...]int{2: 3, 1: 2}
	var c = [...]int{1, 2, 4: 5, 6}
	_, _, _ = a, b, c
}

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
	var c = [...]int{1, 2, 4: 5, 6}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]： %d\n", i, c[i])
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}
}

func main() {
	varArr2()
}
