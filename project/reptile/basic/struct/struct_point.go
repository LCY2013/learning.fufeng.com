package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	i int
	j int
}

/*
将 MyStruct 指针修改成 int 类型的，那么访问新指针就会返回整型变量 i，将指针移动 8 个字节之后就能获取下一个成员变量 j
 */
func myFunction(ms *MyStruct) {
	ptr := unsafe.Pointer(ms)
	for i := 0; i < 2; i++ {
		c := (*int)(unsafe.Pointer((uintptr(ptr) + uintptr(8*i))))
		*c += i + 1
		fmt.Printf("[%p] %d\n", c, *c)
	}
}

func main() {
	a := &MyStruct{i: 40, j: 50}
	myFunction(a)
	fmt.Printf("[%p] %v\n", a, a)
}
