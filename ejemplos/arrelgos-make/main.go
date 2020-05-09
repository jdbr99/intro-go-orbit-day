package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	arr2 := [4]int{4, 5, 6, 7}
	slc := make([]int, 10)
	slc2 := make([]float32, 0, 10)
	testChan := make(chan int)
	buffChan := make(chan int, 3)

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(slc)
	fmt.Println(slc2)
	fmt.Println(testChan)
	fmt.Println(buffChan)
}
