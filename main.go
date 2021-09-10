package main

import (
	"fmt"
)

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 0
	for j := false; j != true; {
		sum2 += 1
		if sum2 > 10 {
			j = true
		}
	}
	fmt.Println(sum2)

	var sum3 = 0

	k := false
	for k != true {
		sum3 += 1

		if sum3 >= 12 {
			k = true
		}
	}
	fmt.Println(sum3)

	sum = 1
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)

	// for {
	// 	fmt.Println("blah blah")
	// }

	// IF
	if value := 10; value < 11 {
		fmt.Println(value)
	}

	if value := 100; value == 10 {
		fmt.Println("value is 10")
	} else {
		fmt.Println("value is not 10, it's", value)
	}

	fmt.Println(Sqrt(25))

}

func Sqrt(x float64) float64 {

	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / 2 * z
		fmt.Println(z)
	}
	return z
}
