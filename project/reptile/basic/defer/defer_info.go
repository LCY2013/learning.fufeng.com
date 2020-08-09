package main

import (
	"fmt"
	"time"
)

/* 结果如下:
4
3
2
1
0
倒序执行所有向 defer 关键字中传入的表达式，最后一次 defer 调用传入了 fmt.Println(4)，所以会这段代码会优先打印 4
 */
func deferRunWithForeach() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/*
结果如下:
block execute
deferOrderOut out
defer execute
defer 传入的函数不是在退出代码块的作用域时执行的，它只会在当前函数和方法返回之前被调用
 */
func deferOrderOut()  {
	{
		defer fmt.Println("defer execute")
		fmt.Println("block execute")
	}
	fmt.Println("deferOrderOut out")
}

/*
Go 语言中所有的函数调用都是传值的，defer 虽然是关键字，但是也继承了这个特性
 */
func calcTime()  {
	startAt := time.Now()
	//下面代码的运行结果并不符合我们的预期
	//我们会发现调用 defer 关键字会立刻对函数中引用的外部参数进行拷贝，所以 time.Since(startedAt) 的结果不是在 main 函数退出之前计算的，而是在 defer 关键字调用时计算的
	//defer fmt.Println(time.Since(startAt))
	//虽然调用 defer 关键字时也使用值传递，但是因为拷贝的是函数指针，所以 time.Since(startedAt) 会在函数返回前被调用并打印出符合预期的结果
	defer func() {fmt.Println(time.Since(startAt))}()
	time.Sleep(time.Second)
}

/*
defer 关键字的调用时机以及多次调用 defer 时执行顺序是如何确定的；
defer 关键字使用传值的方式传递参数时会进行预计算，导致不符合预期的结果；
 */
func main() {
	//示例一
	deferRunWithForeach()
	//示例二
	//函数参数化
	fc := deferOrderOut
	fc()
	//示例三
	calcTime()
}





/*
源码defer结构体定义位置
https://github.com/golang/go/blob/cfe3cd903f018dec3cb5997d53b1744df4e53909/src/runtime/runtime2.go#L853-L878
/usr/local/go/src/runtime/runtime2.go:861
siz 是参数和结果的内存大小；
sp 和 pc 分别代表栈指针和调用方的程序计数器；
fn 是 defer 关键字中传入的函数；
_panic 是触发延迟调用的结构体，可能为空；

https://github.com/golang/go/blob/22d28a24c8b0d99f2ad6da5fe680fa3cfa216651/src/runtime/panic.go#L218-L258
runtime.deferproc 函数负责创建新的延迟调用；
https://github.com/golang/go/blob/22d28a24c8b0d99f2ad6da5fe680fa3cfa216651/src/runtime/panic.go#L526-L571
runtime.deferreturn 函数负责在函数调用结束时执行所有的延迟调用；
这两个函数是 defer 关键字运行时机制的入口，我们从它们开始分别介绍这两个函数的执行过程。

后调用的 defer 函数会先执行：
后调用的 defer 函数会被追加到 Goroutine _defer 链表的最前面；
运行 runtime._defer 时是从前到后依次执行；
函数的参数会被预先计算；
调用 runtime.deferproc 函数创建新的延迟调用时就会立刻拷贝函数的参数，函数的参数不会等到真正执行时计算；
 */