package main

import (
	"fmt"
	"learngo/src/calc"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("Random num", rand.Intn(10))
	fmt.Println("Random num", rand.Intn(10))

	fmt.Println("Random num after seed", rand.Intn(10))

	fmt.Println(calc.Add(10, 32))
	fmt.Println(calc.Sub(10, 3))
}
