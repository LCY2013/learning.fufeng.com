package main

import "fmt"

//GOOS=linux GOARCH=amd64 go tool compile -S stringOp.go
func main() {
	//字符串转换
	str := "hello"
	byteStr := []byte(str)
	fmt.Println(byteStr)
	strByte := string(byteStr)
	println(strByte)
}

func show1()  {
	str := `
		1
		2
	`
	println(str)
}

func show2()  {
	str := "this"
	println(str)
}

func show3()  {
	str := 's'
	println(str)
}

//字符串拼接
func show4()  {
	//str = runtime.concatstrings("",)

}
