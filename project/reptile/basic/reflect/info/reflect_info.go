package main

import (
	"fmt"
	"reflect"
)

/*
调用 reflect.ValueOf 函数获取变量指针；
调用 reflect.Value.Elem 方法获取指针指向的变量；
调用 reflect.Value.SetInt 方法更新变量的值：

go build -gcflags="-S -N" main.go
 */
func main() {
	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)

	// ==
	j := 2
	k := &j
	*k = 20
	fmt.Println(j)
}
