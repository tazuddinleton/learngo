package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(v int) {
			ch <- v
			fmt.Println("value written", v)
			wg.Done()
			fmt.Println("done called", v)
		}(i)
	}

	go func() {
		fmt.Println("waiting...")
		wg.Wait()
		fmt.Println("wait over closing channel...")
		close(ch)
	}()

	Display(ch)
}

func Display(ch chan int) {
	fmt.Println("Display got called")
	for i := range ch {
		fmt.Println("received", i)
	}
	fmt.Println("Display exiting...")

}
