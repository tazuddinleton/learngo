package greetings

import "fmt"

func Hello(name string) (greetings string) {
	greetings = fmt.Sprintf("Hi, %v. Welcome!", name)
	return
}
