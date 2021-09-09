package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("Random num", rand.Intn(10))
	fmt.Println("Random num", rand.Intn(10))

	fmt.Println("Random num after seed", rand.Intn(10))
}
