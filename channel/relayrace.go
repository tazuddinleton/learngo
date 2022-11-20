package channel

import (
	"fmt"
	"time"
)

// Wait group is used to wait for a program to finish.
// var wg sync.WaitGroup

func RelayRace() {

	// Create an unbuffered channel.
	baton := make(chan int)

	// Add a count of one for the last runner
	wg.Add(1)

	// First runnder to his mark
	go Runner(baton)

	// Start the race
	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	// Wait to receive the baton
	runner := <-baton
	// Start running around the track
	fmt.Printf("Runner %d running with baton\n", runner)

	// New runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// Running around the track
	time.Sleep(100 * time.Millisecond)

	// Is the race over
	if runner == 4 {
		fmt.Printf("Runner %d finished, Race over!\n", runner)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner
	fmt.Printf("Runner %d exchange with runner %d\n", runner, newRunner)

	baton <- newRunner
}
