package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	noCondition()
	usingIf()
}

func usingIf() {
	t := time.Now()
	if t.Hour() < 12 {
		fmt.Println("Good morning")
	} else if t.Hour() < 17 {
		fmt.Println("Good afternoon")
	} else {
		fmt.Println("Good evening")
	}

}

func noCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func whensSaturdayExample() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away")
	}
}

func casesNotRun(i int) {
	switch i {
	case 0:
		fmt.Println("0")
	case returnsOne():
		fmt.Println("1")
	}
}

func returnsOne() int {
	return 1
}

func osExample() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s \n", os)
	}
}
