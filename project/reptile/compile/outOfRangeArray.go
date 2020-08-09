package main

//GOSSAFUNC=outOfRange go build outOfRangeArray.go
func outOfRange() int {
	arr := [...]int {1,2,3}
	idx := 3
	elem := arr[idx]
	return elem
}
