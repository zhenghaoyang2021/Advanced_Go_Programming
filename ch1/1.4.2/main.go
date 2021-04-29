package main

import (
	"fmt"
	"image/color"
	"sync"
)

// C 语言风格 函数
type File struct {
	fd int
}

// 类似构造函数
func OpenFile(name string) (f *File, err error) {
	//
	return nil, err
}

func CloseFile(f *File) error {
	//
	return nil
}
func ReadFile(f *File, offset int64, data []byte) int {
	return 0
}

// Go风格1
func (f *File) CloseFile() (err error) {
	//
	return err
}

func (f *File) ReadFile(offset int64, data []byte) int {
	//
	return 0
}

// Go风格 简化  只属于File的函数、方法
func (f *File) Open() (err error) {
	//
	return err
}

func (f *File) Read(offset int64, data []byte) int {
	//
	return 0
}

//继承 通过组合实现  嵌入匿名成员
type Point struct {
	X, Y float64
}
type ColoredPoint struct {
	//通过嵌入匿名成员，不仅继承成员的内部成员，还继承方法
	Point
	Color color.RGBA
	sync.Mutex
}

// oop 方法 面向对象
//
func main() {
	var cp ColoredPoint
	cp.X = 1  // 内部成员 xｙ
	cp.Lock() //内部方法 Lock 编译时会展开 p.Mutex.Lock(),无运行时代价
	cp.Unlock()
	fmt.Println(cp.Point.X)

}
