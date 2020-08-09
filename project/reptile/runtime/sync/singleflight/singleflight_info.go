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
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

/*
	x/sync/singleflight.Group 是 Go 语言扩展包中提供了另一种同步原语，它能够在一个服务中抑制对下游的多次重复请求。
	一个比较常见的使用场景是 — 我们在使用 Redis 对数据库中的数据进行缓存，发生缓存击穿时，大量的流量都会打到数据库上进而影响服务的尾延时。
*/
func main() {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			go singleFlightUse("o")
		}else {
			go singleFlightUse("j")
		}
	}

	time.Sleep(time.Second*5)
}

var sg singleflight.Group

func singleFlightUse(key string) {
	_, _, _ = sg.Do(key, func() (interface{}, error) {
		fmt.Println("starting ", key)
		return nil, nil
	})
}

/*
https://github.com/golang/sync/blob/cd5d95a43a6e21273425c7ae415d3df9ea832eeb/singleflight/singleflight.go#L33-L36
x/sync/singleflight.Group 提供了两个用于抑制相同请求的方法：
	x/sync/singleflight.Group.Do — 同步等待的方法 Do；
	x/sync/singleflight.Group.DoChan — 返回 Channel 异步等待的方法；

每次调用 x/sync/singleflight.Group.Do 方法时都会获取互斥锁，随后判断是否已经存在 key 对应的 x/sync/singleflight.call 结构体：
	当不存在对应的 x/sync/singleflight.call 时：
		初始化一个新的 x/sync/singleflight.call 结构体指针；
		增加 sync.WaitGroup 持有的计数器；
		将 x/sync/singleflight.call 结构体指针添加到映射表；
		释放持有的互斥锁；
		阻塞地调用 x/sync/singleflight.Group.doCall 方法等待结果的返回；
	当存在对应的 x/sync/singleflight.call 时；
		增加 dups 计数器，它表示当前重复的调用次数；
		释放持有的互斥锁；
		通过 sync.WaitGroup.Wait 等待请求的返回；

运行传入的函数 fn，该函数的返回值就会赋值给 c.val 和 c.err；
调用 sync.WaitGroup.Done 方法通知所有等待结果的 Goroutine — 当前函数已经执行完成，可以从 call 结构体中取出返回值并返回了；
获取持有的互斥锁并通过管道将信息同步给使用 x/sync/singleflight.Group.DoChan 方法的 Goroutine；

当我们需要减少对下游的相同请求时，就可以使用 x/sync/singleflight.Group 来增加吞吐量和服务质量，不过在使用的过程中我们也需要注意以下的几个问题：
x/sync/singleflight.Group.Do 和 x/sync/singleflight.Group.DoChan 一个用于同步阻塞调用传入的函数，一个用于异步调用传入的参数并通过 Channel 接收函数的返回值；
x/sync/singleflight.Group.Forget 方法可以通知 x/sync/singleflight.Group 在持有的映射表中删除某个键，接下来对该键的调用就不会等待前面的函数返回了；
一旦调用的函数返回了错误，所有在等待的 Goroutine 也都会接收到同样的错误；
*/
