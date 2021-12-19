package variables

import (
	"fmt"
	"math/cmplx"
)

func Run() {
	var a = "hello"
	var b = "world"
	c := a + " " + b
	fmt.Println(c)

	var j, k = "J", "K"
	print(j, k)

	var x, y, z int = 10, 10, 10
	fmt.Println(x + y + z)

	p := 10
	q := 10
	r := 20
	fmt.Println(p + q + r)

	isOk := false
	isOk = !isOk
	print(isOk)

	// block declarations
	var (
		toBe   bool       = false
		maxInt uint64     = 1<<64 - 1
		comp   complex128 = cmplx.Sqrt(-5 + 12i)
	)
	print(toBe, maxInt, comp)

	// const
	const s = "this is constant string"
	const f = 10
	print(s, f)

	var (
		ii int
		jj bool
		kk string
		ll complex128
	)
	print(ii, jj, kk, ll)

	// var iii int = 42
	// var fff float64 = float64(iii)

	iii := 42
	fff := float64(iii)
	print(fff)
	print(int(fff) / iii)
	print(fff / float64(iii/2))

	var xx int = 0
	jjj := xx + 10
	print(jjj)

	var abc = "abc"
	abcd := abc
	print(abcd)

	v := 42.42 + 1 + 1
	print(v)

	const (
		BigInt   = 1 << 100
		SmallInt = BigInt >> 99
	)

	print(BigInt>>38, SmallInt)
}


func print(a ...interface{}) {
	fmt.Println(a...)
}