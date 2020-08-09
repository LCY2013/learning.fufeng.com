/*
 * The MIT License (MIT)
 * ------------------------------------------------------------------
 * Copyright © 2019 Ramostear.All Rights Reserved.
 *
 * ProjectName: learning.fufeng.com
 * @Author : <a href="https://github.com/lcy2013">MagicLuo</a>
 * @date : 2020-07-08
 * @version : 1.0.0-RELEASE
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */
package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

/*
	信号量是在并发编程中常见的一种同步机制，在需要控制访问资源的进程数量时就会用到信号量，它会保证持有的计数器在 0 到初始化的权重之间波动。
		每次获取资源时都会将信号量中的计数器减去对应的数值，在释放时重新加回来；
		当遇到计数器大于信号量大小时就会进入休眠等待其他线程释放信号；
	Go 语言的扩展包中就提供了带权重的信号量 x/sync/semaphore.Weighted，我们可以按照不同的权重对资源的访问进行管理，这个结构体对外也只暴露了四个方法：
		x/sync/semaphore.NewWeighted 用于创建新的信号量；
		x/sync/semaphore.Weighted.Acquire 阻塞地获取指定权重的资源，如果当前没有空闲资源，就会陷入休眠等待；
		x/sync/semaphore.Weighted.TryAcquire 非阻塞地获取指定权重的资源，如果当前没有空闲资源，就会直接返回 false；
		x/sync/semaphore.Weighted.Release 用于释放指定权重的资源；
*/
func main() {
	//示例一
	semaphoreUse()

	time.Sleep(time.Second*5)
}

func semaphoreUse() {
	weighted := semaphore.NewWeighted(10)
	for i := 0; i < 20; i++ {
		go func(i int) {
			weighted.Acquire(context.Background(), 1)
			fmt.Println(i)
			time.Sleep(time.Second*2)
			weighted.Release(1)
		}(i)
	}
}

/*
带权重的信号量确实有着更多的应用场景，这也是 Go 语言对外提供的唯一一种信号量实现，在使用的过程中我们需要注意以下的几个问题：
	x/sync/semaphore.Weighted.Acquire 和 x/sync/semaphore.Weighted.TryAcquire 方法都可以用于获取资源，前者会阻塞地获取信号量，后者会非阻塞地获取信号量；
	x/sync/semaphore.Weighted.Release 方法会按照 FIFO 的顺序唤醒可以被唤醒的 Goroutine；
	如果一个 Goroutine 获取了较多地资源，由于 x/sync/semaphore.Weighted.Release 的释放策略可能会等待比较长的时间；
*/
