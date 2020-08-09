package main

//算术定义结构体
type Arithmetic struct {
	ariOne int
	ariTwo int
}

func (ari *Arithmetic) add() int {
	return ari.ariOne + ari.ariTwo
}

func (ari *Arithmetic) reduce() int {
	return ari.ariOne + ari.ariTwo
}

func (ari *Arithmetic) ride() int {
	return ari.ariOne + ari.ariTwo
}

func (ari *Arithmetic) div() int {
	return ari.ariOne + ari.ariTwo
}

/*func main() {
	ari := Arithmetic{
		1,
		2,
	}
	fmt.Println(ari.add())


}
*/
