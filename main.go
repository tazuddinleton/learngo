package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

func main() {

	fmt.Println(wordCount("hello world, hello world, hello again world, hello world again."))

}

func wordCount(s string) map[string]int {

	sp := strings.Fields(s)
	wc := make(map[string]int)

	for i := 0; i < len(sp); i++ {
		_, ok := wc[sp[i]]
		if !ok {
			wc[sp[i]] = 1
		} else {
			wc[sp[i]]++
		}
	}
	return wc
}

func mapMake2() {
	var points = map[string]Point{
		"north": {10001, 10002},
		"south": {10002, 1000032},
	}
	fmt.Println(points)

	delete(points, "south")
	fmt.Println(points)
	points["south"] = Point{x: 10003, y: -11000}

	east, ok := points["east"]

	if !ok {
		fmt.Println(east)
	}

	printPoint(east)
}

func printPoint(point Point) {
	fmt.Printf("x: %v, y: %v", point.x, point.y)
}

func mapMake() {
	var points = make(map[string]Point)
	fmt.Println(points)
	points["NorthPole"] = Point{
		x: 1000001, y: 10000002}
	fmt.Println(points)
	fmt.Println(points["NorthPole"])
}
