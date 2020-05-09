package main

// Merge junta dos arreglos
func Merge(left, right []int) []int {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		return right
	}
	if rightLength == 0 {
		return left
	}

	result := make([]int, (leftLength + rightLength))

	li := 0
	ri := 0
	resulti := 0
	var rnum, lnum int

	for li < leftLength || ri < rightLength {
		if li < leftLength && ri < rightLength {
			lnum = left[li]
			rnum = right[ri]

			if lnum <= rnum {
				result[resulti] = lnum
				li++
			} else {
				result[resulti] = rnum
				ri++
			}

		} else if li < leftLength {
			lnum = left[li]
			result[resulti] = lnum
			li++
		} else if ri < rightLength {
			rnum = right[ri]
			result[resulti] = rnum
			ri++
		}

		resulti++
	}

	return result
}

// MergeSort es la implementación recursiva habitual de mergesort
func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	m := length / 2

	sortedLeft := MergeSort(arr[0:m])
	sortedRight := MergeSort(arr[m:length])

	return Merge(sortedLeft, sortedRight)
}

// ConcurrentMerge es la implementación concurrente de Merge
func ConcurrentMerge(left, right []int, resultChan chan []int) {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		resultChan <- right
		return
	}
	if rightLength == 0 {
		resultChan <- left
		return
	}

	result := make([]int, (leftLength + rightLength))
	li := 0
	ri := 0
	resulti := 0
	var rnum, lnum int

	for li < leftLength || ri < rightLength {
		if li < leftLength && ri < rightLength {
			lnum = left[li]
			rnum = right[ri]

			if lnum <= rnum {
				result[resulti] = lnum
				li++
			} else {
				result[resulti] = rnum
				ri++
			}

		} else if li < leftLength {
			lnum = left[li]
			result[resulti] = lnum
			li++
		} else if ri < rightLength {
			rnum = right[ri]
			result[resulti] = rnum
			ri++
		}

		resulti++
	}

	resultChan <- result
}

// ConcurrentMergeSort es la implementación concurrente de MergeSort
func ConcurrentMergeSort(arr []int, resultChan chan []int) {
	length := len(arr)
	if length <= 1 {
		resultChan <- arr
		return
	}
	if length <= 10000 {
		resultChan <- MergeSort(arr)
		return
	}

	m := length / 2

	lchan := make(chan []int, 1)
	rchan := make(chan []int, 1)

	go ConcurrentMergeSort(arr[0:m], lchan)
	go ConcurrentMergeSort(arr[m:length], rchan)
	go ConcurrentMerge(<-lchan, <-rchan, resultChan)
}
