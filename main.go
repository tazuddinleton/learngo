package main

import "github.com/tazuddinleton/learngo/generics"

func main() {
	// rn := []func(){}
	// rn = append(rn, arrays.Run, err.Run, fn.Run, variables.Run)

	// for _, f := range rn {
	// 	fmt.Println("")
	// 	f()
	// }

	generics.RunNonGenericSum()
	generics.RunGenericSum()
}
