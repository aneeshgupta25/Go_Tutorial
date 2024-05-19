package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) //helps go to identify when to close the channel so that for loop in main
	// can stop waiting for new messages from the channel.
	// But for this approach, we must know the longest time consuming operation to close the channel
}

func main() {
	doneChan := make(chan bool)
	// doneChans := make([]chan bool, 4)

	// doneChans[0] = make(chan bool)
	go greet("Nice to meet you!", doneChan)
	// doneChans[1] = make(chan bool)
	go greet("How are you?", doneChan)
	// doneChans[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", doneChan)
	// doneChans[3] = make(chan bool)
	go greet("I hope you're liking the course!", doneChan)

	for doneChannel := range doneChan {
		fmt.Println(doneChannel)
	}

	// for _, doneChan := range doneChans {
	// 	<- doneChan
	// }
}

// channel can be used to wait for the completion of a goroutine
