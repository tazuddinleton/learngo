package main

import (
	"fmt"

	"github.com/tazuddinleton/learngo/arrays"
	err "github.com/tazuddinleton/learngo/errors"
	"github.com/tazuddinleton/learngo/fn"
	"github.com/tazuddinleton/learngo/variables"
)

func main() {
	rn := []func(){}
	rn = append(rn, arrays.Run, err.Run, fn.Run, variables.Run)

	for _, f := range rn {
		fmt.Println("")
		f()
	}
}
