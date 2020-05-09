package main

import "fmt"

func sendNumber(n int, numChan chan int) {
	numChan <- n
}

func main() {
	numberChannel := make(chan int)

	go sendNumber(9)

	received := <-numberChannel

	fmt.Println(received)
}
