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
	time.Sleep(3 * time.Second) //simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	numGoroutines := 4
	dones := make([]chan bool, numGoroutines)
	// for i := 0; i < numGoroutines; i++ {
	// 	dones[i] = make(chan bool)
	// 	go func(i int) {
	// 		switch i {
	// 		case 0:
	// 			greet("Nice to meet you!", dones[i])
	// 		case 1:
	// 			greet("How are you?", dones[i])
	// 		case 2:
	// 			slowGreet("How ...are ...you ...?", dones[i])
	// 		case 3:
	// 			greet("I hope you're liking the course!", dones[i])
	// 		default:
	// 			fmt.Println("Invalid index")
	// 			dones[i] <- true // Enviar sinal mesmo para índices inválidos
	// 		}
	// 	}(i)
	// }
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you!", dones[0])
	// dones[1] = make(chan bool)
	// go greet("How are you?", dones[1])
	// dones[2] = make(chan bool)
	// go slowGreet("How ...are ...you ...?", dones[2])
	// dones[3] = make(chan bool)
	// go greet("I hope you're liking the course!", dones[3])

	// for _, done := range dones {
	// 	<-done
	// }

	done := make(chan bool)
	go greet("Nice to meet you!", done)
	dones[1] = make(chan bool)
	go greet("How are you?", done)
	dones[2] = make(chan bool)
	go slowGreet("How ...are ...you ...?", done)
	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	for range done {

	}
}
