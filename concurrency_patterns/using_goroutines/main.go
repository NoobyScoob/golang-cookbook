package main

// This program illustates how to use goroutines
// and how it may sometimes cause non-deterministic outputs
// outputs 10 different values every time, enter anything to exit

import "fmt"

var counter int = 0		// shared variable

func incrementCounter() {
	// looping, huge number of times takes time
	for i := 0; i < 100000; i++ {
		counter++
	}
}

func displayCounter() {
	for i := 0; i < 10; i++ {
		fmt.Println(counter)
	}
}

func main() {

	// async execution
	// go keyword is used to spawn goroutines
	go incrementCounter()
	go displayCounter()

	// stopping main to complete exection
	var input string
	fmt.Scanln(&input)

}