package main

import (
	"fmt"
	"learngo/greetings"
	"log"
)

func main() {
	msg, err := greetings.Hello("Tahmid")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(msg)
}

func init() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
}
