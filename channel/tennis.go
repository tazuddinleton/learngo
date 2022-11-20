package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for a program to finish
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Run() {
	// Create an unbuffered channel
	court := make(chan int)

	// Add count 2, one for each goroutine
	wg.Add(2)

	// Launch two players
	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set
	court <- 1

	// Wait for the game to finihsh
	wg.Wait()
}

// player simulates the player of tennish game
func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for {
		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			// If channel was closed we wone.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we missed the ball.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// Close the channel to signal that we are lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player
		time.Sleep(time.Second)
		court <- ball
	}

}
