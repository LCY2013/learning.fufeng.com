package main

import "fmt"

//go参数传递方式 - 值传递

func myFunction(i int, arr *[2]int) {
	i = 40
	arr[0] = 55
	fmt.Printf("in my_funciton - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}

func main() {
	i := 30
	arr := &[2]int{66, 77}
	fmt.Printf("before calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
	myFunction(i, arr)
	fmt.Printf("after  calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}
