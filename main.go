package main

import (
	"fmt"
	"sync"
)

const (
	length = 12
)

func main() {
	//read array of ints
	intSequence := getInts()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("\nGot Q1: %v Sorting...\n", intSequence[:3])
		bubbleSort(intSequence[:3])
		fmt.Printf("Sorted Q1: %v\n", intSequence[:3])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("\nGot Q2: %v Sorting...\n", intSequence[3:6])
		bubbleSort(intSequence[3:6])
		fmt.Printf("Sorted Q2: %v\n", intSequence[3:6])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("\nGot Q3: %v Sorting...\n", intSequence[6:9])
		bubbleSort(intSequence[6:9])
		fmt.Printf("Sorted Q3: %v\n", intSequence[6:9])

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("\nGot Q4: %v Sorting...\n", intSequence[9:])
		bubbleSort(intSequence[9:])
		fmt.Printf("Sorted Q4: %v\n", intSequence[9:])
	}()

	wg.Wait()
	bubbleSort(intSequence[:])
	fmt.Printf("\nSorted all: %v\n", intSequence[:])
}

func getInts() [length]int {
	//declare array of ints
	intSequence := [length]int{}

	// prompt the user to enter a sequence of (length) ints
	fmt.Printf("Please enter a sequence of %d integers.\n", length)

	for i := 0; i < length; i++ {
		fmt.Printf("Int #%d: ", i+1)
		fmt.Scan(&intSequence[i])
	}

	return intSequence
}

func bubbleSort(q []int) {
	lengthLess1 := len(q) - 1

	for j := 0; j < lengthLess1; j++ {
		for i := 0; i < lengthLess1; i++ {
			if q[i] > q[i+1] {
				hold := q[i]
				q[i] = q[i+1]
				q[i+1] = hold
			}
		}
	}
}
