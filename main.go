package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Pic(10, 10))
}

func Pic(dx, dy int) [][]uint8 {

	var x = make([][]uint8, dy)

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			x[i] = append(x[i], uint8(j*5))
		}
	}
	return x
}

func basics() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	b := [10]int{1, 2, 3}

	for i := 0; i < 10; i++ {
		if b[i] == 0 {
			fmt.Println("False")
		} else {
			fmt.Println("True")
		}
	}
	fmt.Println(a, b)
}

func slice() {
	nums := [6]int{1, 2, 3, 43, 4, 4}

	s1 := nums[1:3]
	s1[0] = 100
	s1 = append(s1, 10)
	s1 = append(s1, 10)
	s1 = append(s1, 10)
	s1 = append(s1, 10)
	fmt.Println(s1)
	fmt.Println(nums)

	s2 := []int{}
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	s2 = append(s2, 4000)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	s3 := s2[:15]
	fmt.Println(len(s3))
}

func makeSlice() {
	a := make([]int, 0, 100)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 0)
	a = append(a, 1)
	fmt.Println(len(a), cap(a), a)

	b := a[:cap(a)]
	fmt.Println(len(b), cap(b), b)

}

func ticTacToeBoard() {
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	board[0][0] = "x"
	board[0][1] = "0"
	board[0][2] = "x"
	board[1][0] = "x"
	board[1][1] = "0"
	board[1][2] = "0"
	board[2][0] = "0"
	board[2][1] = "x"
	board[2][2] = "x"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendToSlice() {
	s := []string{"hello"}
	s = append(s, "world", "This is another text", "and this is another one")

	h := strings.Join(s[:2], " ")

	fmt.Printf("%s\n", strings.Trim(h, " ")+", "+strings.Join(s[2:], ", "))
}

func sliceRange() {
	var nums = []int{}
	for i := 0; i < 10; i++ {
		nums = append(nums, i*2)
	}
	fmt.Println(nums)

	// for i, v := range nums {
	// 	fmt.Println(i, v)
	// }

	for _, v := range nums {
		fmt.Println(v)
	}
}
