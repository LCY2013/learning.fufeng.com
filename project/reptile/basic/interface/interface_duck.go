package main

//GOSSAFUNC=Quack go build interface_duck.go
//GOOS=linux GOARCH=amd64 go tool compile -S interface_duck.go
type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

func (cat *Cat) Quack()  {
	println(cat.Name , " -> running")
}

func main() {
	/*
	我们将生成的汇编指令拆分成三部分分析：
	结构体 Cat 的初始化；
	赋值触发的类型转换过程；
	调用接口的方法 Quack()；
	 */
	var cat Duck = &Cat{"cat"}
	cat.Quack()
	switch cat.(type) {
	case *Cat:
		c := cat.(*Cat)
		c.Quack()
	}
}
