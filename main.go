package main

import (
	"fmt"
)

func main() {
	arr := []int{6, 2, 1, 3, 7, 5, 4, 8}
	n := len(arr)
	input1 := make([]int, n)

	_ = copy(input1, arr)

	fmt.Printf("Before Sort : \n%v\n", input1)

	fmt.Println("Performing Merge Sort...")
	performMergeSort(input1)

	fmt.Printf("After sort: \n%v\n", input1)

	// performBubbleSort(input1, true)

	// fmt.Println(input1)
	//performInsertionSort(input1)

	// fmt.Println(input1)

	// fmt.Println("----------------")

	// input2 := make([]int, n)

	// _ = copy(input2, arr)
	// insertionSorted := performSelectionSort(input2)

	// fmt.Println(insertionSorted)

}
