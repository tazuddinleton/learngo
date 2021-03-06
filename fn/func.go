//Package func shows how to use func in golang, functions are first class citizens in golang.
package fn

import "fmt"

type Logger struct {
	Info  func(string) int
	Debug func(string) int
	Error func(string) int
}

//Run runs all example of current package
func Run() {
	fmt.Println(">>\nPackage fn")
	f := Fibonacci()
	fmt.Println("Fibonacci upto 10")
	for i := 0; i < 10; i++ {
		fmt.Println("Next fib: ", f())
	}
}

// Fibonacci returns a func that returns the next fibonacci number each time gets called
func Fibonacci() func() int {
	fib := []int{0}
	return func() int {
		l := len(fib)
		if l == 1 {
			fib = append(fib, 1)
			return 0
		}
		val := fib[l-1]
		fib = append(fib, fib[l-1]+fib[l-2])
		return val
	}
}

// RunLogExample runs the logger specific example
func RunLogExample() {
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

// DoubleEach multiplies each element of given slice and prints result
func DoubleEach(s []int) {
	fmt.Println(ForEach(s, func(x int) int {
		return x * 2
	}))
}

// ForEach takes a slice and a func to apply to each element, returns the updated slice
func ForEach(s []int, fn func(int) int) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = fn(s[i])
	}
	return res
}
