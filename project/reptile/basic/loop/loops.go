package main

import "fmt"

//遍历一个数组同时向里面追加元素，最终结果还是循环原有数组的次数
func rangeForArray()  {
	arr := []int{1,2,3,4}
	for _,v := range arr{
		arr = append(arr,v)
	}
	fmt.Println(arr)
}

//当我们在遍历一个数组时，如果获取 range 返回变量的地址并保存到另一个数组或者哈希时，
//就会遇到令人困惑的现象
func rangeForArrayCopy()  {
	arr := []int{1,2,3}
	var newArr []*int
	for i,_ := range arr {
		newArr = append(newArr,&arr[i])
	}
	for _,v := range newArr {
		fmt.Println(*v)
	}
}

/*
当我们想要在 Go 语言中清空一个切片或者哈希表时，我们一般都会使用以下的方法将切片中的元素置零，
但是依次去遍历切片和哈希表看起来是非常耗费性能的事情
https://github.com/golang/go/blob/05c02444eb2d8b8d3ecd949c4308d8e2323ae087/src/runtime/memclr_386.s#L12-L16
 */
func memCls()  {
	arr := []int{1,2,3}
	for i := range arr{
		arr[i] = 0
	}
	// 编译器通过 runtime.memclrNoHeapPointers 清空切片中的数据
}

/*
 Go 语言中使用 range 遍历哈希表时，往往都会使用如下的代码结构，但是这段代码在每次运行时都会打印出不同的结果：
  Go 语言故意的设计，它在运行时为哈希表的遍历引入不确定性，也是告诉所有使用 Go 语言的使用者，程序不要依赖于哈希表的稳定遍历
 */
func hashRange()  {
	hash := map[string]int{
		"1":1,
		"2":2,
		"3":3,
	}
	for k,v := range hash {
		fmt.Println(k,"->",v)
	}
}

/*
Go 语言中的经典循环在编译器看来是一个 OFOR 类型的节点，这个节点由以下四个部分组成：
	初始化循环的 Ninit；
	循环的继续条件 Left；
	循环体结束时执行的 Right；
	循环体 NBody：
eg:
	for Ninit; Left; Right {
    	NBody
	}
https://github.com/golang/go/blob/4d5bb9c60905b162da8b767a8a133f6b4edcaa65/src/cmd/compile/internal/gc/ssa.go#L1023-L1502

节点类型的转换过程都发生在 SSA 中间代码生成阶段，所有的 for/range 循环都会被 cmd/compile/internal/gc.walkrange 函数转换成不包含复杂结构、只包含基本表达式的语句。接下来，我们按照循环遍历的元素类型依次介绍遍历数组和切片、哈希表、字符串以及管道时的过程

cmd/compile/internal/gc.arrayClear 是一个非常有趣的优化，这个函数会优化 Go 语言遍历数组或者切片并删除全部元素的逻辑：
 */

func main() {
	rangeForArray()
	rangeForArrayCopy()
	memCls()
	hashRange()
}
