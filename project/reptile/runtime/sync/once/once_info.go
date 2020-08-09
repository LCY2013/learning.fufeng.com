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
)

/*
	Go 语言标准库中 sync.Once 可以保证在 Go 程序运行期间的某段代码只会执行一次.
	作为用于保证函数执行次数的 sync.Once 结构体，它使用互斥锁和 sync/atomic 包提供的方法实现了某个函数在程序运行期间只能执行一次的语义。在使用该结构体时，我们也需要注意以下的问题：
		sync.Once.Do 方法中传入的函数只会被执行一次，哪怕函数中发生了 panic；
		两次调用 sync.Once.Do 方法传入不同的函数也只会执行第一次调用的函数；
*/
func main() {
	//示例一
	onceUse()
}

/*
	https://github.com/golang/go/blob/bc593eac2dc63d979a575eccb16c7369a5ff81e0/src/sync/once.go#L12-L20
 */
func onceUse() {
	one := sync.Once{}
	for i := 0; i < 10; i++ {
		one.Do(func() {
			fmt.Println(i)
		})
	}
}
