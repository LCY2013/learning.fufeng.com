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
	"golang.org/x/sync/errgroup"
	"net/http"
)

/*

 */
func main() {
	errGroupUse()
}

/*
	x/sync/errgroup.Group 就为我们在一组 Goroutine 中提供了同步、错误传播以及上下文取消的功能，
	我们可以使用如下所示的方式并行获取网页的数据.
*/
func errGroupUse() {
	var eg errgroup.Group
	var urls = []string{
		"http://q.cnblogs.com/q/103554/",
		"http://www.baidu.com/",
		"http://www.somestupidname.com/",
	}
	for i := range urls {
		url := urls[i]
		eg.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := eg.Wait(); err == nil {
		fmt.Println("fetch all urls successful")
	}
}
/*
x/sync/errgroup.Group.Go 方法能够创建一个 Goroutine 并在其中执行传入的函数，
而 x/sync/errgroup.Group.Wait 会等待所有 Goroutine 全部返回，该方法的不同返回结果也有不同的含义：
	如果返回错误 — 这一组 Goroutine 最少返回一个错误；
	如果返回空值 — 所有 Goroutine 都成功执行；

x/sync/errgroup.Group 结构体同时由三个比较重要的部分组成：
	cancel — 创建 context.Context 时返回的取消函数，用于在多个 Goroutine 之间同步取消信号；
	wg — 用于等待一组 Goroutine 完成子任务的同步原语；
	errOnce — 用于保证只接收一个子任务返回的错误；

x/sync/errgroup.Group 的实现没有涉及底层和运行时包中的 API，它只是对基本同步语义进行了封装以提供更加复杂的功能。在使用时，我们也需要注意以下的几个问题：
	x/sync/errgroup.Group 在出现错误或者等待结束后都会调用 context.Context 的 cancel 方法同步取消信号；
	只有第一个出现的错误才会被返回，剩余的错误都会被直接抛弃；
 */
