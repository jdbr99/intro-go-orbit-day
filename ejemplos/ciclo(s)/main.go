package main

import "fmt"

func main() {
	fmt.Println("Ciclo `for` clásico")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Utilizando `for` como `while`")
	i := 0
	for i < 9 {
		fmt.Println(i * i)
		i++
	}
}
