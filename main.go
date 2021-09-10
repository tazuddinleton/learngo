package main

import (
	"fmt"
)

func main() {

	i, j := 42, 2701

	p := &i

	*p = 21
	pp := &p
	ppp := &pp

	fmt.Println(j)
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(pp)
	fmt.Println(ppp)
	fmt.Println(***ppp)
	fmt.Println(**pp)
	fmt.Println(*p)
}
