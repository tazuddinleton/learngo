package main

import "fmt"

func main() {
	var a = "hello"
	var b = "world"
	c := a + " " + b
	fmt.Println(c)

	var x, y, z int = 10, 10, 10
	fmt.Println(x + y + z)

	p := 10
	q := 10
	r := 20

	fmt.Println(p + q + r)

	isOk := false
	isOk = !isOk
	print(isOk)

}

func print(a ...interface{}) {
	fmt.Println(a...)
}
