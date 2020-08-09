package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int { return a + b }

/*
通过 reflect.ValueOf 获取函数 Add 对应的反射对象；
根据反射对象 reflect.rtype.NumIn 方法返回的参数个数创建 argv 数组；
多次调用 reflect.ValueOf 函数逐一设置 argv 数组中的各个参数；
调用反射对象 Add 的 reflect.Value.Call 方法并传入参数列表；
获取返回值数组、验证数组的长度以及类型并打印其中的数据；
 */
func main() {
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // #=> 1
}
