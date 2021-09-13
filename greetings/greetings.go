package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Hello returns greeting for the named person, error when name is empty
func Hello(name string) (greetings string, err error) {
	if strings.Trim(name, " ") == "" {
		return "", errors.New("emtpy name")
	}
	greetings = fmt.Sprintf(randomFormat(), name)
	return
}

func randomFormat() string {

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hello Mr. %v, nice to see you.",
	}
	idx := rand.Intn(len(formats))
	return formats[idx]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
