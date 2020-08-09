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

/*
Go 1.9 版本之前，所有的计时器由全局唯一的四叉堆维护1；
Go 1.10 ~ 1.13，全局使用 64 个四叉堆维护全部的计时器，每个处理器（P）创建的计时器会由对应的四叉堆维护2；
Go 1.14 版本之后，每个处理器单独管理计时器并通过网络轮询器触发3；

https://github.com/golang/go/blob/093adeef4004fd029de1a8fd138802607265dc73/src/runtime/time.go#L28-L37

这个结构体中的字段 t 就是最小四叉堆，创建的所有计时器都会加入到四叉堆中。runtime.timerproc#093ade Goroutine 会运行时间驱动的事件，它会在发生以下事件时会被唤醒：
	四叉堆中的计时器到期；
	四叉堆中加入了触发时间更早的新计时器；

    状态						解释
timerNoStatus			还没有设置状态
timerWaiting			等待触发
timerRunning			运行计时器函数
timerDeleted			被删除
timerRemoving			正在被删除
timerRemoved			已经被停止并从堆中删除
timerModifying			正在被修改
timerModifiedEarlier	被修改到了更早的时间
timerModifiedLater		被修改到了更晚的时间
timerMoving				已经被修改正在被移动
 */
func main() {

}

/*
触发调度器
	调度器调度时会检查处理器中的计时器是否准备就绪；
	系统监控会检查是否有未执行的到期计时器；
https://github.com/golang/go/blob/8d7be1e3c9a98191f8c900087025c5e78b73d962/src/runtime/proc.go#L2623
runtime.checkTimers 是调度器用来运行处理器中计时器的函数，它会在发生以下情况时被调用：
	调度器调用 runtime.schedule 执行调度时；
	调度器调用 runtime.findrunnable 获取可执行的 Goroutine 时；
	调度器调用 runtime.findrunnable 从其他处理器窃取计时器时；

 */
func timerUse()  {

}






