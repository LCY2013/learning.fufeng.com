package main

import (
	"context"
	"fmt"
	"time"
)

/*
	context.Context:
	https://github.com/golang/go/blob/df2999ef43ea49ce1578137017949c0ee660608a/src/context/context.go#L62-L154
	context.Context 是 Go 语言在 1.7 版本中引入标准库的接口，该接口定义了四个需要实现的方法，其中包括：
		Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
		Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个 Channel；
		Err — 返回 context.Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值；
			如果 context.Context 被取消，会返回 Canceled 错误；
			如果 context.Context 超时，会返回 DeadlineExceeded 错误；
		Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；
	context 包中提供的 context.Background、context.TODO、context.WithDeadline 和 context.WithValue 函数会返回实现该接口的私有结构体，我们会在后面详细介绍它们的工作原理。

	我们可能会创建多个 Goroutine 来处理一次请求，而 context.Context 的作用就是在不同 Goroutine 之间同步请求特定数据、取消信号以及处理请求的截止日期。
	每一个 context.Context 都会从最顶层的 Goroutine 一层一层传递到最下层。context.Context 可以在上层 Goroutine 执行出现错误时，将信号及时同步给下层。
	当最上层的 Goroutine 因为某些原因执行失败时，下层的 Goroutine 由于没有接收到这个信号所以会继续工作；但是当我们正确地使用 context.Context 时，就可以在下层及时停掉无用的工作以减少额外资源的消耗
	两个私有变量都是通过 new(emptyCtx) 语句初始化的，它们是指向私有结构体 context.emptyCtx 的指针，这是最简单、最常用的上下文类型

	从源代码来看，context.Background 和 context.TODO 函数其实也只是互为别名，没有太大的差别。它们只是在使用和语义上稍有不同：
		context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来；
		context.TODO 应该只在不确定应该使用哪种上下文时使用；

	context.WithCancel 函数能够从 context.Context 中衍生出一个新的子上下文并返回用于取消该上下文的函数（CancelFunc）。一旦我们执行返回的取消函数，当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号
 */
func main() {
	//示例一
	contextUseCase()
}

/*
 context 使用例子
 */
func contextUseCase()  {
	// timeout 取消上下文
	ctx , cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	//因为过期时间大于处理时间，所以我们有足够的时间处理该『请求』，运行上述代码会打印出如下所示的内容
	//handle 函数没有进入超时的 select 分支，但是 main 函数的 select 却会等待 context.Context 的超时并打印出 main context deadline exceeded。
	//context.Context 的使用方法和设计原理 — 多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作
	//goroutine 执行handle
	go handle(ctx,500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("contextUseCase",ctx.Err())
	}
}

/*
 ctx : 上下文信息
 duration : 时间戳
 */
func handle(ctx context.Context,duration time.Duration)  {
	select {
	case <- ctx.Done():
		fmt.Println("handle",ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with :",duration)
	}
}


