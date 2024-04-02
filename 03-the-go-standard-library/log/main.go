package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("We did not create a trace file! %v\n", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close trace file %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("We failed to start a trace: %d\n", err)
	}

	defer trace.Stop()

	AddRandomNumbers()
}

func AddRandomNumbers() {
	firstNumber := rand.Intn(100)
	secondNumber := rand.Intn(100)

	time.Sleep(2 * time.Second)

	var result = firstNumber * secondNumber

	fmt.Printf("Result of 2 numbers is %d\n", result)
}
