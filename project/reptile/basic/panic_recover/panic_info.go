package main

import (
	"fmt"
	"time"
)

/*
	跨协程失效
	 panic 只会触发当前 Goroutine 的延迟函数调用
多个 Goroutine 之间没有太多的关联，一个 Goroutine 在 panic 时也不应该执行其他 Goroutine 的延迟函数。
*/
func overGoroutineFail() {
	defer fmt.Println("defer in overGoroutineFail")
	go func() {
		defer fmt.Println("defer in goroutine")
		panic("panic")
	}()
	time.Sleep(time.Second)
}

/*
	失效的崩溃恢复
在主程序中调用 recover 试图中止程序的崩溃，但是从运行的结果中我们也能看出，如下所示的程序依然没有正常退出。
背后的原因，recover 只有在发生 panic 之后调用才会生效。然而在上面的控制流中，recover 是在 panic 之前调用的，并不满足生效的条件，所以我们需要在 defer 中使用 recover 关键字
*/
func failRecovery() {
	/*defer fmt.Println("in failRecovery")
	if err := recover(); err != nil {
		fmt.Println(err)
	}*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("unknown error")
}

/*
	嵌套崩溃
Go 语言中的 panic 是可以多次嵌套调用的。如下所示的代码就展示了如何在 defer 函数中多次调用 panic
我们可以确定程序多次调用 panic 也不会影响 defer 函数的正常执行。所以使用 defer 进行收尾的工作一般来说都是安全的
 */
func nestingCrash()  {
	defer fmt.Println("in nestingCrash")
	defer func() {
		defer func() {
			panic("crash two again")
		}()
		panic("crash again")
	}()
	panic("crash once")
}

/*
panic 只会触发当前 Goroutine 的延迟函数调用；
recover 只有在 defer 函数中调用才会生效；
panic 允许在 defer 中嵌套多次调用；

panic引起的退出来之于runtime.exit
https://github.com/golang/go/blob/cbaa666682386fe5350bf87d7d70171704c90fe4/src/runtime/sys_darwin.go#L231-L233
defer 中的 recover 是如何中止程序崩溃的。编译器会将关键字 recover 转换成 runtime.gorecover：
*/
func main() {
	//示例一
	//overGoroutineFail()
	//示例二
	//failRecovery()
	//示例三
	//nestingCrash()

}

/*
 源码地址
https://github.com/golang/go/blob/cfe3cd903f018dec3cb5997d53b1744df4e53909/src/runtime/runtime2.go#L891-L900
	argp 是指向 defer 调用时参数的指针；
	arg 是调用 panic 时传入的参数；
	link 指向了更早调用的 runtime._panic 结构；
	recovered 表示当前 runtime._panic 是否被 recover 恢复；
	aborted 表示当前的 panic 是否被强行终止；
从数据结构中的 link 字段我们就可以推测出以下的结论 — panic 函数可以被连续多次调用，它们之间通过 link 的关联形成一个链表。
结构体中的 pc、sp 和 goexit 三个字段都是为了修复 runtime.Goexit 的问题引入的2。该函数能够只结束调用该函数的 Goroutine 而不影响其他的 Goroutine，但是该函数会被 defer 中的 panic 和 recover 取消3，引入这三个字段的目的就是为了解决这个问题。

编译器会负责做转换关键字的工作；
	将 panic 和 recover 分别转换成 runtime.gopanic 和 runtime.gorecover；
	将 defer 转换成 deferproc 函数；
	在调用 defer 的函数末尾调用 deferreturn 函数；
在运行过程中遇到 gopanic 方法时，会从 Goroutine 的链表依次取出 _defer 结构体并执行；
如果调用延迟执行函数时遇到了 gorecover 就会将 _panic.recovered 标记成 true 并返回 panic 的参数；
	在这次调用结束之后，gopanic 会从 _defer 结构体中取出程序计数器 pc 和栈指针 sp 并调用 recovery 函数进行恢复程序；
	recovery 会根据传入的 pc 和 sp 跳转回 deferproc；
	编译器自动生成的代码会发现 deferproc 的返回值不为 0，这时会跳回 deferreturn 并恢复到正常的执行流程；
如果没有遇到 gorecover 就会依次遍历所有的 _defer 结构，并在最后调用 fatalpanic 中止程序、打印 panic 的参数并返回错误码 2；


这些情况都不能成功 recover ：

// 相隔 0 个
defer recover()
panic(1)

// 相隔 2 个
defer func() {
    defer func() {
	    recover()
    }()
}()
panic(1)
而这些可以成功 recover：

// 相隔 1 个
defer func() {
	recover()
}()
panic(1)

// 相隔 1 个
defer func() { // 只有这个在 gopanic 中调用
	recover()
}()

defer func() {
	panic(1)
}()
 */