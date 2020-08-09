package main

import "fmt"

/*
	通过make初始化内置的数据结构
slice 是一个包含 data、cap 和 len 的私有结构体 internal/reflectlite.sliceHeader；
hash 是一个指向 runtime.hmap 结构体的指针；
ch 是一个指向 runtime.hchan 结构体的指针；
*/
func makeInGo() {
	slices := make([]int, 0, 100)
	hash := make(map[int]string, 10)
	ch := make(chan int, 3)
	fmt.Println(slices)
	fmt.Println(hash)
	fmt.Println(ch)
}

/*
new 的功能就很简单了，它只能接收一个类型作为参数然后返回一个指向该类型的指针
*/
func newInGo() {
	i := new(int)
	fmt.Println(i)
	fmt.Println(*i)

	//上面的代码等价于下面的代码
	var ix int
	ixPointer := &ix
	fmt.Println(ixPointer)
	fmt.Println(ix)
}

/*
当我们想要在 Go 语言中初始化一个结构时，可能会用到两个不同的关键字 — make 和 new。
因为它们的功能相似，所以初学者可能会对这两个关键字的作用感到困惑，但是它们两者能够初始化的却有较大的不同。
	make 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel；
	new 的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针；

通过 var 或者 new 创建的变量不需要在当前作用域外生存，例如不用作为返回值返回给调用方，那么就不需要初始化在堆上。
runtime.newobject 函数会是获取传入类型占用空间的大小，调用 runtime.mallocgc 在堆上申请一片内存空间并返回指向这片内存空间的指针
https://github.com/golang/go/blob/5042317d6919d4c84557e04be35130430e8d1dd4/src/runtime/malloc.go#L1162-L1164
https://github.com/golang/go/blob/5042317d6919d4c84557e04be35130430e8d1dd4/src/runtime/malloc.go#L889-L1132
*/
func main() {
	//示例一
	makeInGo()
	//示例二
	newInGo()
}
