package main

import "fmt"

/**
当元素数量小于或者等于 4 个时，会直接将数组中的元素放置在栈上；
当元素数量大于 4 个时，会将数组中的元素放置到静态区并在运行时取出；
 */
func main() {
	arr1 := [3]int{1,2,3}

	arr2 := [...]int{4,5,6}

	arr1[0] = 2

	fmt.Println(arr1)
	fmt.Println(arr2)
}
