package main

import "fmt"

func saludar(name string) {
	fmt.Println("Hola,", name)
}

func main() {
	go saludar("David")

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
