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
	"sync"
	"time"
)

/*
	Go 语言标准库中的 sync.Cond 一个条件变量，它可以让一系列的 Goroutine 都在满足特定条件时被唤醒。每一个 sync.Cond 结构体在初始化时都需要传入一个互斥锁

	sync.Cond 不是一个常用的同步机制，在遇到长时间条件无法满足时，与使用 for {} 进行忙碌等待相比，sync.Cond 能够让出处理器的使用权。在使用的过程中我们需要注意以下问题：
		sync.Cond.Wait 方法在调用之前一定要使用获取互斥锁，否则会触发程序崩溃；
		sync.Cond.Signal 方法唤醒的 Goroutine 都是队列最前面、等待最久的 Goroutine；
		sync.Cond.Broadcast 会按照一定顺序广播通知等待的全部 Goroutine；
*/
func main() {
	//示例一
	condUse()
}

/*
上述代码同时运行了 11 个 Goroutine，这 11 个 Goroutine 分别做了不同事情：
	10 个 Goroutine 通过 sync.Cond.Wait 等待特定条件的满足；
	1 个 Goroutine 会调用 sync.Cond.Broadcast 方法通知所有陷入等待的 Goroutine；
调用 sync.Cond.Broadcast 方法后，上述代码会打印出 10 次 “listen” 并结束调用。
 */
func condUse() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go lister(c)
	}
	time.Sleep(time.Second)
	go broadcast(c)

	time.Sleep(time.Second*3)
	/*ch := make(chan os.Signal,1)
	signal.Notify(ch,os.Interrupt)
	<-ch*/
}

/*
sync.Cond.Signal 方法会唤醒队列最前面的 Goroutine；
sync.Cond.Broadcast 方法会唤醒队列中全部的 Goroutine；
 */
func broadcast(c *sync.Cond) {
	c.L.Lock()
	c.Broadcast()
	c.L.Unlock()
}

func lister(c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println("lister")
	c.L.Unlock()
}

/*
查询按type分类的数据每类数据在10000以内
A表
id,name,type
select id,name,type from A a where 10000 > (select count(id) from A where a.type= type and a.id > id)
 */