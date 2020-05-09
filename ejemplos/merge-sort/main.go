package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenLargeArray genera un slice de tamaño n con valores pseudo-aleatorios
func GenLargeArray(n int) []int {
	temp := make([]int, 0, n)
	for i := 0; i < n; i++ {
		temp = append(temp, rand.Int())
	}
	return temp
}

func main() {
	test := GenLargeArray(10000000)

	fmt.Println("Mergesort Clásico")
	st1 := time.Now()
	_ = MergeSort(test)
	e1 := time.Now()
	// fmt.Println(r)
	tiempo1 := e1.UnixNano() - st1.UnixNano()
	fmt.Printf("Tiempo: %dns\n", tiempo1)

	fmt.Println("Mergesort con Goroutines")
	st2 := time.Now()
	resultChan := make(chan []int, 1)
	ConcurrentMergeSort(test, resultChan)
	_ = <-resultChan
	e2 := time.Now()
	// fmt.Println(k)
	tiempo2 := e2.UnixNano() - st2.UnixNano()
	fmt.Printf("Time: %dns\n", tiempo2)
	fmt.Printf("Diferencia: %dns\n", tiempo1-tiempo2)
	fmt.Printf("La versión concurrente tomó %f%% del tiempo de la versión clásica\n", float32(tiempo2)*100/float32(tiempo1))
}
