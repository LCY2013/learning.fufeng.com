package main

import "fmt"

//GOSSAFUNC=hello go build hello.go
//GOOS=linux GOARCH=amd64 go tool compile -S hello.go
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(3,8))
}
