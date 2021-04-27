package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"sort"
	"unsafe"
)

func varSlice() {
	var (
		a []int            // nil切片 与nil相等
		b = []int{}        //空切片  与nil不相等
		c = []int{1, 2, 3} //
		d = c[:2]          //len=2 cap=3
		e = c[0:2:cap(c)]  //len=2 cap=3 保证最小的连续空间大小

		f = c[:0]          // len=0 cap=3
		g = make([]int, 3) //len=cap=3
		h = make([]int, 2, 3)
		//i = make([]int, 0, 10)
		//zz = i[0:11:10] Invalid index values, must be low <= high <= max
	)
	for i := range a {
		fmt.Printf("a[%d]： %d\n", i, a[i])
	}

	for i, v := range b {
		fmt.Printf("b[%d]： %d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]： %d\n", i, c[i])
	}

	a = append(a, 1)
	a = append(a, 1, 2, 3)
	a = append(a, []int{1, 2, 3}...)

	//在开头增加
	a = append([]int{}, a...)
	a = append([]int{-3, 2, 1}, a...)

	//x位 中间增加
	var i = 2
	a = append(a[:i], append([]int{1}, a[:i]...)...)
	a = append(a[:i], append([]int{-3, 2, 1}, a[:i]...)...) // i位 插入一个切片
	// 第二个append 会创建一个临时slice

	// copy + append 可避免创建临时slice
	a = append(a, 0)     //slice 增加一个空间
	copy(a[i+1:], a[i:]) // a[i:]向后移一位
	a[i] = 0             //添加新元素

	x := []int{11, 12}
	a = append(a, x...)       //slice 增加一个空间
	copy(a[i+len(x):], a[i:]) //移len位
	copy(a[i:], x)

	a = a[:len(a)-1] // 删尾部一个
	//a = a[:len(a)-N] 删N个
	a = a[1:] //删开头一个
	//a = a[N:]

	// 原地完成
	a = append(a[:0], a[1:]...)
	//a = append(a[:0],a[N:]...)

	//删除中间一个元素
	a = append(a[:i], a[i+1:]...)
	//a = append(a[:i],a[i+N:]...) //N个

	//删除中间一个元素
	a = a[:i+copy(a[i:], a[i+1:])]
	//a = a[:i+copy(a[i:],a[i+N:])]

	_ = a
	_ = b
	_ = c
	_ = d
	_ = e
	_ = f
	_ = g
	_ = h
	_ = i

}

//len=0,cap!=0 useful
//example
func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}
func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}
func main() {
	// 判断 slice是否为空 使用 len
	// del l
	fn := func(l byte) bool {
		if l == 'l' {
			return true
		}
		return false
	}
	fmt.Println(string(Filter([]byte("hello"), fn)))
	fmt.Println(string(TrimSpace([]byte("he      llo"))))

	//避免内存渗漏
	//var a []*int{}
	//a = a[:len(a)-1]//被删除的最后一个元素依然被引用，可能导致垃圾回收器被阻碍

	//var a []*int{}
	//a[len(a)-1] = nil //垃圾回收器回收最后一个元素内存
	//a = a[:len(a)-1]/

}

//避免内存渗漏
func FindPhoneNumber(filename string) []byte {
	//一个小小的内存引用导致整个底层数组处于被使用状态，延迟了垃圾回收器对底层数组的回收
	b, _ := ioutil.ReadFile(filename)
	return regexp.MustCompile("[0-9]+").Find(b)
}
func FindPhoneNumberB(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	//将感兴趣的数组复制到一个新的切片，切断原始数据的依赖
	return append([]byte{}, b...)
}
func SortFloat64(a []float64) {
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	sort.Ints(b)
}
func SortFloat63FastV2(a []float64) {
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr
	sort.Ints(c)
}
