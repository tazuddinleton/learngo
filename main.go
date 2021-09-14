package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	i, err := doubleIt(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	x := 25.0
	j, err := sqrt(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt of %v is %v", x, j)
	}

	fmt.Println(ErrNegativeSqrt(-2).Error())
}

func doubleIt(x int64) (int64, error) {
	if x < 0 {
		return 0, NegativeNumberError{Message: "We don't support doubling negative numbers", TimeStamp: time.Now()}
	}
	return x * 2, nil
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

type NegativeNumberError struct {
	TimeStamp time.Time
	Message   string
}

func (err NegativeNumberError) Error() string {
	return fmt.Sprintf("%v. Error happen at %s", err.Message, err.TimeStamp)
}

type ErrNegativeSqrt float64

func (v ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("We don't Sqrt negative numbers: %v", int64(v))
}
