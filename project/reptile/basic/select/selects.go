package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	下面控制结构会等待 c <- x 或者 <-quit 两个表达式中任意一个的返回。
	无论哪一个表达式返回都会立刻执行 case 中的代码，当 select 中的两个 case 同时被触发时，
	就会随机选择一个 case 执行
 */
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			//fmt.Println("default")
		}
	}
}

//定义task
type task interface {
	Run() error
}

type taskImpl struct {
	taskName string
}

type SelectError struct {
	code int
	msg string
}

func (se SelectError) Error() string {
	return fmt.Sprintf("%d -> %s",se.code,se.msg)
}

func (ti taskImpl) Run() error {
	fmt.Println(ti.taskName," -> Running")
	if rand.Int() % 2 == 0 {
		return nil
	}
	return SelectError{500,"server inner error"}
}

/*
非阻塞的 Channel 发送和接收操作还是很有必要的，在很多场景下我们不希望向 Channel 发送消息
或者从 Channel 中接收消息会阻塞当前 Goroutine，我们只是想看看 Channel 的可读或者可写状态
 */
func tasks(tasks []task) error {
	errCh := make(chan error, len(tasks))
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for i := range tasks {
		go func(i int) {
			defer wg.Done()
			if err := tasks[i].Run(); err != nil {
				errCh <- err
			}
		}(i)
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}

/*
使用 select 遇到的情况是同时有多个 case 就绪时，select 会选择那个 case 执行的问题

select 在遇到多个 <-ch 同时满足可读或者可写条件时会随机选择一个 case 执行其中的代码

select 提交5引入并一直保留到现在的，虽然中间经历过一些修改，但是语义一直都没有改变。
在下面的代码中，两个 case 都是同时满足执行条件的，如果我们按照顺序依次判断，
那么后面的条件永远都会得不到执行，而随机的引入就是为了避免饥饿问题的发生
 */
func manySelect()  {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		case <-ch:
			return
		}
	}
}

/*
select 在 Go 语言的源代码中不存在对应的结构体，但是 select 控制结构中的 case 却使用
runtime.scase https://github.com/golang/go/blob/d1969015b4ac29be4f518b94817d3f525380639d/src/runtime/select.go#L28-L34
结构体来表示
runtime.scase 结构体中也包含一个 runtime.hchan 类型的字段存储 case 中使用的 Channel
https://github.com/golang/go/blob/d1969015b4ac29be4f518b94817d3f525380639d/src/runtime/chan.go#L32-L51
*/
func main() {
	//示例一
	c := make(chan int)
	quit := make(chan int)
	go fibonacci(c, quit)
	go func() {
		for i:=0;i<10;i++ {
			ret := <- c
			fmt.Println(ret)
		}
		quit <- -1
	}()
	time.Sleep(time.Millisecond * 100)

	//示例二
	ta := []task{
		taskImpl{"task1"},
		taskImpl{"task2"},
		taskImpl{"task3"},
		taskImpl{"task4"},
	}
	err := tasks(ta)
	fmt.Println(err)

	//示例三
	manySelect()

}

/*
caseNil：当前 case 不包含 Channel；
	这种 case 会被跳过；
caseRecv：当前 case 会从 Channel 中接收数据；
	如果当前 Channel 的 sendq 上有等待的 Goroutine，就会跳到 recv 标签并从缓冲区读取数据后将等待 Goroutine 中的数据放入到缓冲区中相同的位置；
	如果当前 Channel 的缓冲区不为空，就会跳到 bufrecv 标签处从缓冲区获取数据；
	如果当前 Channel 已经被关闭，就会跳到 rclose 做一些清除的收尾工作；
caseSend：当前 case 会向 Channel 发送数据；
	如果当前 Channel 已经被关，闭就会直接跳到 sclose 标签，触发 panic 尝试中止程序；
	如果当前 Channel 的 recvq 上有等待的 Goroutine，就会跳到 send 标签向 Channel 发送数据；
	如果当前 Channel 的缓冲区存在空闲位置，就会将待发送的数据存入缓冲区；
caseDefault：当前 case 为 default 语句；
	表示前面的所有 case 都没有被执行，这里会解锁所有 Channel 并返回，意味着当前 select 结构中的收发都是非阻塞的；
 */


/*
发送
首先是 Channel 的发送过程，当 case 中表达式的类型是 OSEND 时，编译器会使用 if/else 语句和 runtime.selectnbsend 函数改写代码：
select {
case ch <- i:
    ...
default:
    ...
}
https://github.com/golang/go/blob/d1969015b4ac29be4f518b94817d3f525380639d/src/runtime/chan.go#L662-L664
https://github.com/golang/go/blob/d1969015b4ac29be4f518b94817d3f525380639d/src/runtime/chan.go#L157-L278
if selectnbsend(ch, i) {
    ...
} else {
    ...
}
这段代码中最重要的就是 runtime.selectnbsend 函数，它为我们提供了向 Channel 非阻塞地发送数据的能力。
我们在 Channel 一节介绍了向 Channel 发送数据的 runtime.chansend 函数包含一个 block 参数，该参数会决定这一次的发送是不是阻塞的：
func selectnbsend(c *hchan, elem unsafe.Pointer) (selected bool) {
	return chansend(c, elem, false, getcallerpc())
}
由于我们向 runtime.chansend 函数传入了 false，所以哪怕是不存在接收方或者缓冲区空间不足都不会阻塞当前 Goroutine 而是会直接返回。



接收
由于从 Channel 中接收数据可能会返回一个或者两个值，所以接受数据的情况会比发送稍显复杂，不过改写的套路是差不多的：
// 改写前
select {
case v <- ch: // case v, ok <- ch:
    ......
default:
    ......
}
https://github.com/golang/go/blob/d1969015b4ac29be4f518b94817d3f525380639d/src/runtime/chan.go#L705-L709
// 改写后
if selectnbrecv(&v, ch) { // if selectnbrecv2(&v, &ok, ch) {
    ...
} else {
    ...
}
返回值数量不同会导致使用函数的不同，两个用于非阻塞接收消息的函数 runtime.selectnbrecv 和 runtime.selectnbrecv2 只是对 runtime.chanrecv 返回值的处理稍有不同：
func selectnbrecv(elem unsafe.Pointer, c *hchan) (selected bool) {
	selected, _ = chanrecv(c, elem, false)
	return
}
func selectnbrecv2(elem unsafe.Pointer, received *bool, c *hchan) (selected bool) {
	selected, *received = chanrecv(c, elem, false)
	return
}
因为接收方不需要，所以 runtime.selectnbrecv 会直接忽略返回的布尔值，而 runtime.selectnbrecv2 会将布尔值回传给调用方。与 runtime.chansend 一样，runtime.chanrecv 也提供了一个 block 参数用于控制这一次接收是否阻塞。
 */