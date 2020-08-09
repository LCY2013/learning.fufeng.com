package main

//GOOS=linux GOARCH=amd64 go tool compile -S hello.go
//go tool compile -S -N -l add.go  注意:如果编译时不使用 -N -l 参数，编译器会对汇编代码进行优化，编译结果会有较大差别
func div(a int,b int) int {
	return a - b
}

func main() {
	div(8,1)
}
