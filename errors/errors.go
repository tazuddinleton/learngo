/**
Errors package shows examples of errors in golang
*/

package errors

import (
	"fmt"
	"math"
	"time"
)

// Run runs all the examples of current package
func Run() {
	i, err := DoubleIt(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	x := 25.0
	j, err := Sqrt(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Printf("sqrt of %v is %v", x, j))
	}

	fmt.Println(ErrNegativeSqrt(-2).Error())
}

// DoubleIt doubles the positive integers and returns it, returns errors and zero value if negative
func DoubleIt(x int64) (int64, error) {
	if x < 0 {
		return 0, NegativeNumberError{Message: "We don't support doubling negative numbers", TimeStamp: time.Now()}
	}
	return x * 2, nil
}

// Sqrt returns square root of float64 positive numbers, error with zero value if negative
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

type NegativeNumberError struct {
	// Defines a negative number error object
	TimeStamp time.Time
	Message   string
}

// Error is the implementation of the interface error by the type NegativeNumberError
func (err NegativeNumberError) Error() string {
	return fmt.Sprintf("%v. Error happen at %s", err.Message, err.TimeStamp)
}

// ErrNegativeSqrt is a float64 type
type ErrNegativeSqrt float64

// Error is the implementation of the interface error by the type ErrNegativeSqrt
func (v ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("We don't Sqrt negative numbers: %v", int64(v))
}
