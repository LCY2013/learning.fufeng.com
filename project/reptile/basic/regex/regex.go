package main

import (
	"fmt"
	"regexp"
)

//const text  = "My email is 1475653689@qq.com"
const text  = `
	My email is 1475653689@qq.com
	isdjf@gmail.com
	is0djf@gmail.com
	is1djf@gmail.com
	isdjf3@email.com
`
func main() {
	re := regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?")
	//go 语言``语意是里面的东西都不转译，原样输入
	//re := regexp.MustCompile(`[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?`)
	findString := re.FindAllString(text,-1)
	fmt.Printf("%s\n",findString)

	re = regexp.MustCompile("([\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*)@((?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?)")
	submatch := re.FindAllStringSubmatch(text, -1)
	fmt.Println(submatch)
}
