package main

import (
	"fmt"
)

type Logger struct {
	Info  func(string) int
	Debug func(string) int
	Error func(string) int
}

func main() {

	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	fib := []int{0}
	return func() int {
		if len(fib) == 1 {
			fib = append(fib, 1)
			return 0
		}

		val := fib[len(fib)-1]
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
		return val
	}
}

func runLogExample() {
	log := Logger{
		Info:  GetLogger(0),
		Debug: GetLogger(1),
		Error: GetLogger(-1),
	}

	log.Info("hello world")
	log.Debug("hello world")
	log.Debug("hello world")
	log.Info("hello world")
	log.Error("hello world")
	log.Info("hello world")
	log.Info("hello world")
	log.Info("hello default world")
}

// GetLogger returns log function based on level
func GetLogger(level int) func(string) int {
	logCount := 0
	switch level {
	case 0:
		return func(s string) int {
			logCount += 1
			fmt.Printf("info: (%v) %v\n", logCount, s)
			return logCount
		}
	case 1:
		return func(s string) int {
			logCount += 1
			fmt.Printf("debug: (%v) %v\n", logCount, s)
			return logCount
		}
	default:
		return func(s string) int {
			logCount += 1
			fmt.Printf("error: (%v) %v\n", logCount, s)
			return logCount
		}
	}
}

func multiPlyEach() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ForEach(s, func(x int) int {
		return x * 2
	}))
}

// ForEach takes a slice and a func to apply to each element, returns
func ForEach(s []int, fn func(int) int) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = fn(s[i])
	}
	return res
}
