package main

import (
	"fmt"
)

//https://github.com/golang/go/blob/be64a19d99918c843f8555aad580221207ea35bc/src/cmd/compile/internal/gc/reflect.go#L82-L187
//https://github.com/golang/go/blob/f07059d949057f414dd0f8303f93ca727d716c62/src/cmd/compile/internal/gc/sinit.go#L768-L873
//https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/runtime/map.go#L303-L336
/**
计算哈希占用的内存是否溢出或者超出能分配的最大值；
调用 fastrand 获取一个随机的哈希种子；
根据传入的 hint 计算出需要的最小需要的桶的数量；
使用 runtime.makeBucketArray 创建用于保存桶的数组；
runtime.makeBucketArray 函数会根据传入的 B 计算出的需要创建的桶数量在内存中分配一片连续的空间用于存储数据

当桶的数量小于 2^4 时，由于数据较少、使用溢出桶的可能性较低，这时就会省略创建的过程以减少额外开销；
当桶的数量多于 2^4 时，就会额外创建 2^𝐵−4 个溢出桶，根据上述代码，我们能确定在正常情况下，
正常桶和溢出桶在内存中的存储空间是连续的，只是被 hmap 中的不同字段引用，
当溢出桶数量较多时会通过 runtime.newobject 创建新的溢出桶。
*/

func main() {
	//初始化
	hash := map[interface{}]interface{}{
		"一": 1,
		"二": 2,
		"三": 3,
		"四": 4,
		5:   "5",
	}

	fmt.Println(hash)

	foreachMap(hash)

	//https://github.com/golang/go/blob/4d5bb9c60905b162da8b767a8a133f6b4edcaa65/src/cmd/compile/internal/gc/walk.go#L439-L1532
	//https://github.com/golang/go/blob/36f30ba289e31df033d100b2adb4eaf557f05a34/src/runtime/map.go#L685-L791
	delete(hash,5)

	foreachMap(hash)

	//https://github.com/golang/go/blob/36f30ba289e31df033d100b2adb4eaf557f05a34/src/runtime/map.go#L394-L450
	//函数会先通过哈希表设置的哈希函数、种子获取当前键对应的哈希，再通过 bucketMask 和 add 函数拿到该键值对所在的桶序号和哈希最上面的 8 位数字。
	//https://github.com/golang/go/blob/36f30ba289e31df033d100b2adb4eaf557f05a34/src/runtime/map.go#L452-L508
	//写入
	//https://github.com/golang/go/blob/36f30ba289e31df033d100b2adb4eaf557f05a34/src/runtime/map.go#L571-L683
	value,ok := hash["一"]
	if ok{
		fmt.Println(value)
	}
	//扩容
	//https://github.com/golang/go/blob/36f30ba289e31df033d100b2adb4eaf557f05a34/src/runtime/map.go#L1017-L1058
}

func foreachMap(hMap map[interface{}]interface{}) {
	for key, value := range hMap {
		fmt.Println(key, "->", value)
		fmt.Printf("key type %T,value type %T\n", key, value)
	}
}
