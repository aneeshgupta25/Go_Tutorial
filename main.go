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
}

func main() {
	// doneChan := make(chan bool)
	doneChans := make([]chan bool, 4)

	doneChans[0] = make(chan bool)
	go greet("Nice to meet you!", doneChans[0])
	doneChans[1] = make(chan bool)
	go greet("How are you?", doneChans[1])
	doneChans[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", doneChans[2])
	doneChans[3] = make(chan bool)
	go greet("I hope you're liking the course!", doneChans[3])

	for _, doneChan := range doneChans {
		<- doneChan
	}
}

// channel can be used to wait for the completion of a goroutine
