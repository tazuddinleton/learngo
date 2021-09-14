package main

import (
	"fmt"
	"math"
)

func main() {
	p := Person{}

	v := &p

	v.SetFirstName("asdfasdf")

	fmt.Println(v)
}

func (v *Vertex) Scale(s float64) {
	v.x = v.x * s
	v.y = v.y * s
}

func abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

type Vertex struct {
	x, y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Value reciever
func (p Person) SetFirstName(s string) {
	p.FirstName = s
}

func SetFirstName(p Person, s string) {
	p.FirstName = s
}
