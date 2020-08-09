package main

import "testing"

//go test -gcflags=-N -benchmem -test.count=3 -test.cpu=1 -test.benchtime=1s -bench=.

func BenchmarkDirectCall(b *testing.B) {
	c := &Cat{Name: "grooming"}
	for n := 0; n < b.N; n++ {
		// MOVQ	AX, "".c+24(SP)
		// MOVQ	AX, (SP)
		// CALL	"".(*Cat).Quack(SB)
		c.Quack()
	}
}

func BenchmarkDynamicDispatch(b *testing.B) {
	c := Duck(&Cat{Name: "grooming"})
	for n := 0; n < b.N; n++ {
		// MOVQ	"".d+56(SP), AX
		// MOVQ	24(AX), AX
		// MOVQ	"".d+64(SP), CX
		// MOVQ	CX, (SP)
		// CALL	AX
		c.Quack()
	}
}

func BenchmarkDirectSCall(b *testing.B) {
	c := Cat{Name: "grooming"}
	for n := 0; n < b.N; n++ {
		// MOVQ	AX, (SP)
		// MOVQ	$8, 8(SP)
		// CALL	"".Cat.Quack(SB)
		c.Quack()
	}
}

//func BenchmarkDynamicSDispatch(b *testing.B) {
//	c := Duck(Cat{Name: "grooming"})
//	for n := 0; n < b.N; n++ {
//		// MOVQ	16(SP), AX
//		// MOVQ	24(SP), CX
//		// MOVQ	AX, "".d+32(SP)
//		// MOVQ	CX, "".d+40(SP)
//		// MOVQ	"".d+32(SP), AX
//		// MOVQ	24(AX), AX
//		// MOVQ	"".d+40(SP), CX
//		// MOVQ	CX, (SP)
//		// CALL	AX
//		c.Quack()
//	}
//}