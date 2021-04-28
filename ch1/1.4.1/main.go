package main

import "fmt"

//具名
func Add(a, b int) int {
	return a + b
}

// 多参数，多返回
func Swap(a, b int) (int, int) {
	return b, a
}

//可变参数
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}
func main() {
	//匿名
	var Add = func(a, b int) int {
		return a + b
	}
	_ = Add
	var a = []interface{}{123, "abc"}
	fmt.Println(a...) //123 abc a... => 解包
	fmt.Println(a)    //[123 abc] =>a => 未解包

	//闭包隐藏问题
	for i := 0; i < 3; i++ {
		defer func() {
			println("A:", i)
		}()
	}
	//output:  3 3 3 why ? 闭包引用同一个变量
	//fix 1
	for i := 0; i < 3; i++ {
		i := i
		defer func() {
			println("B:", i)
		}()
	}

	//fix 1
	for i := 0; i < 3; i++ {
		defer func(i int) { // 需命名
			println("C:", i)
		}(i)
	}
	// for 循环 不建议defer
}

// 命名返回值
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// 返回值命名了,可通过defer后的语句修改返回值
func Inc() (v int) {
	defer func() {
		v++
	}()
	return 42
}

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type IntSliceHeader struct {
	Data []int //引用
	Len  int
	Cap  int
}

func Twice(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2 //改变了值
	}
	//修改 Len,Cap 无用
	//所以append 需要 重新赋值
}
