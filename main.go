package main

import (
	"fmt"
	"time"
)

func main() {
	runSelectExample()
}

func runSelectExample() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd, even := make(chan int), make(chan int)

	go numberPicker(nums, odd, even)

	for {
		select {
		case v, ok := <-odd:
			if ok {
				fmt.Println("Odd number", v)
			} else {
				return
			}
		case v, ok := <-even:
			if ok {
				fmt.Println("Even number", v)
			} else {
				return
			}
		default:
			fmt.Println("Not ready yet!")
		}
	}
}

func numberPicker(nums []int, odd chan int, even chan int) {
	defer close(odd)
	defer close(even)

	for i := 1; i <= len(nums); i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
}

func runSumExample() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func shootBoomerang(speed int, name string) {
	fmt.Println(name, "Swaaaaa!")
	time.Sleep(time.Millisecond * time.Duration(speed))
	fmt.Println(name, "Swooop!")
}
