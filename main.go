package main

import "fmt"

func main() {
	// var sortType string
	// flag.StringVar(&sortType, "st", "bubble", "Specify sort type. Default is bubble")

	// flag.Usage = func() {
	// 	fmt.Println("Accepted values for sort type are:")
	// 	fmt.Println("- bubble\n- insertion\n- merge\n- selection")
	// 	fmt.Println("Example .\\goalgos.exe -st bubble")
	// }

	// flag.Parse()

	// arr := []int{6, 2, 1, 3, 7, 5, 4, 8}
	// n := len(arr)
	// input1 := make([]int, n)

	// _ = copy(input1, arr)

	// fmt.Printf("Performing %s sort\n", sortType)
	// fmt.Printf("Before Sort : \n%v\n", input1)
	// switch sortType {
	// case "bubble":
	// 	performBubbleSort(input1, true)
	// case "insertion":
	// 	performInsertionSort(input1)
	// case "merge":
	// 	performMergeSort(input1)
	// case "selection":
	// 	performSelectionSort(input1)
	// }

	// fmt.Printf("After sort: \n%v\n", input1)

	// performBubbleSort(input1, true)

	// fmt.Println(input1)
	//performInsertionSort(input1)

	// fmt.Println(input1)

	// fmt.Println("----------------")

	// input2 := make([]int, n)

	// _ = copy(input2, arr)
	// insertionSorted := performSelectionSort(input2)

	// fmt.Println(insertionSorted)

	rootNode := NewNode(3)
	rootNode.Left = NewNode(9)
	rootNode.Right = NewNode(20)
	rootNode.Right.Left = NewNode(25)
	rootNode.Right.Right = NewNode(40)

	result := findMinimumDepth(rootNode)

	fmt.Println(result)

}
