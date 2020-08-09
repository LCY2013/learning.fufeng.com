/*
 * The MIT License (MIT)
 * ------------------------------------------------------------------
 * Copyright © 2019 Ramostear.All Rights Reserved.
 *
 * ProjectName: learning.fufeng.com
 * @Author : <a href="https://github.com/lcy2013">MagicLuo</a>
 * @date : 2020-07-07
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
	同步原语与锁
	Go 语言作为一个原生支持用户态进程（Goroutine）的语言，当提到并发编程、多线程编程时，往往都离不开锁这一概念。锁是一种并发编程中的同步原语（Synchronization Primitives），它能保证多个 Goroutine 在访问同一片内存时不会出现竞争条件（Race condition）等问题。
	Go 语言中常见的同步原语 sync.Mutex、sync.RWMutex、sync.WaitGroup、sync.Once 和 sync.Cond 以及扩展原语 errgroup.Group、semaphore.Weighted 和 singleflight.Group 的实现原理，同时也会涉及互斥锁、信号量等并发编程中的常见概念。
 */
func main() {
	//示例一
	go mutexUse(1)
	go mutexUse(2)
	
	//示例二
	go rwMutexUse(3)
	go rwMutexUse(4)

	time.Sleep(time.Second*2)
}

/*
	Mutex: state 表示当前互斥锁的状态，而 sema 是用于控制锁状态的信号量。
	https://github.com/golang/go/blob/71239b4f491698397149868c88d2c851de2cd49b/src/sync/mutex.go#L25-L28
	状态:
	互斥锁的状态比较复杂，如下图所示，最低三位分别表示 mutexLocked、mutexWoken 和 mutexStarving，剩下的位置用来表示当前有多少个 Goroutine 等待互斥锁的释放：
	在默认情况下，互斥锁的所有状态位都是 0，int32 中的不同位分别表示了不同的状态：
		mutexLocked — 表示互斥锁的锁定状态；
		mutexWoken — 表示从正常模式被从唤醒；
		mutexStarving — 当前的互斥锁进入饥饿状态；
		waitersCount — 当前互斥锁上等待的 Goroutine 个数；
	互斥锁的状态不是 0 时就会调用 sync.Mutex.lockSlow 尝试通过自旋（Spinnig）等方式等待锁的释放，该方法的主体是一个非常大 for 循环，这里将该方法分成几个部分介绍获取锁的过程：
	https://github.com/golang/go/blob/71239b4f491698397149868c88d2c851de2cd49b/src/sync/mutex.go#L84-L171
		判断当前 Goroutine 能否进入自旋；
		通过自旋等待互斥锁的释放；
		计算互斥锁的最新状态；
		更新互斥锁的状态并获取锁；
 */
var mu sync.Mutex
func mutexUse(idx int)  {
	defer mu.Unlock()
	mu.Lock()
	fmt.Println("mutexUse",idx)
	time.Sleep(time.Millisecond*500)
}

var rwMu sync.RWMutex
func rwMutexUse(idx int)  {
	//读锁
	rwMu.RLock()
	fmt.Println("rwMutexUse=>",idx)
	time.Sleep(time.Millisecond*500)
	fmt.Println("rwMutexUse=>",idx,"->")
	rwMu.RUnlock()

	//写锁
	rwMu.Lock()
	fmt.Println("rwMutexUse",idx)
	time.Sleep(time.Millisecond*500)
	fmt.Println("rwMutexUse",idx,"->")
	rwMu.Unlock()
}

/*
互斥锁的加锁过程比较复杂，它涉及自旋、信号量以及调度等概念：
	如果互斥锁处于初始化状态，就会直接通过置位 mutexLocked 加锁；
	如果互斥锁处于 mutexLocked 并且在普通模式下工作，就会进入自旋，执行 30 次 PAUSE 指令消耗 CPU 时间等待锁的释放；
	如果当前 Goroutine 等待锁的时间超过了 1ms，互斥锁就会切换到饥饿模式；
	互斥锁在正常情况下会通过 sync.runtime_SemacquireMutex 函数将尝试获取锁的 Goroutine 切换至休眠状态，等待锁的持有者唤醒当前 Goroutine；
	如果当前 Goroutine 是互斥锁上的最后一个等待的协程或者等待的时间小于 1ms，当前 Goroutine 会将互斥锁切换回正常模式；

互斥锁的解锁过程与之相比就比较简单，其代码行数不多、逻辑清晰，也比较容易理解：
	当互斥锁已经被解锁时，那么调用 sync.Mutex.Unlock 会直接抛出异常；
	当互斥锁处于饥饿模式时，会直接将锁的所有权交给队列中的下一个等待者，等待者会负责设置 mutexLocked 标志位；
	当互斥锁处于普通模式时，如果没有 Goroutine 等待锁的释放或者已经有被唤醒的 Goroutine 获得了锁，就会直接返回；在其他情况下会通过 sync.runtime_Semrelease 唤醒对应的 Goroutine；

读写互斥锁 sync.RWMutex 虽然提供的功能非常复杂，不过因为它建立在 sync.Mutex 上，所以整体的实现上会简单很多。我们总结一下读锁和写锁的关系：
	调用 sync.RWMutex.Lock 尝试获取写锁时；
		每次 sync.RWMutex.RUnlock 都会将 readerWait 其减一，当它归零时该 Goroutine 就会获得写锁；
		将 readerCount 减少 rwmutexMaxReaders 个数以阻塞后续的读操作；
	调用 sync.RWMutex.Unlock 释放写锁时，会先通知所有的读操作，然后才会释放持有的互斥锁；
读写互斥锁在互斥锁之上提供了额外的更细粒度的控制，能够在读操作远远多于写操作时提升性能。

 */
