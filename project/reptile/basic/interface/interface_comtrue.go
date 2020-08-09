package main

import "fmt"

//接口定义
type error interface {
	Error() string
}

type RPCError struct {
	Code int
	Message string
}

//结构体RPCStruct 隐式实现了error接口
func (rpc *RPCError) Error() string {
	return fmt.Sprintf("rpc exception: %v -> %v",rpc.Code,rpc.Message)
}

//创建一个RPCError
func NewRPCError(code int,message string) *RPCError {
	return &RPCError{
		code, message,
	}
}

//返回RPCError
func AsError(err error) error {
	return err
}

/*
将 *RPCError 类型的变量赋值给 error 类型的变量 rpcErr；
将 *RPCError 类型的变量 rpcErr 传递给签名中参数类型为 error 的 AsErr 函数；
将 *RPCError 类型的变量从函数签名的返回值类型为 error 的 NewRPCError 函数中返回；
 */
/*func main() {
	rpcError := NewRPCError(500, "server inner error")
	err := AsError(rpcError)
	fmt.Println(err)
}*/

//打印所有类型
/*
上述函数不接受任意类型的参数，只接受 interface{} 类型的值，在调用 Print 函数时会对参数 v 进行类型转换，
将原来的 Test 类型转换成 interface{} 类型，我们会在本节后面介绍类型转换的过程和原理。
 */
func PrintV(v interface{})  {
	println(v)
}

/*
	结构体实现接口			结构体指针	实现接口
	结构体初始化变量		通过			不通过
	结构体指针初始化变量	通过			通过
 */
/*func main() {
	type user struct {}
	u := user{}
	PrintV(u)
}*/

