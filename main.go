package main

import (
	"container/list"
	"fmt"

	"github.com/ranpariyachetan/goalgos/advent2021"
)

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

	// rootNode := NewNode(1)
	// rootNode.Left = NewNode(2)
	// rootNode.Right = NewNode(3)
	// // rootNode.Right.Left = NewNode(25)
	// rootNode.Right.Right = NewNode(6)

	// result := getLeftView(rootNode)

	// fmt.Println(result)

	// testFindMedian()

	// vals := []int{60, 210, 120}
	// weights := []int{10, 20, 30}

	// w := 50

	// a :=

	// knapsack.TestKnapSack1TopDown()
	// knapsack.TestKnapSack1TopDownMemo()
	// knapsack.TestKnapSack1BottomUp()
	// knapsack.TestKnapSack2TopDown()
	// knapsack.TestKnapSack2BottomUp()
	// knapsack.TestKnapSack2WithDP()
	// knapsack.TestKnapSack3TopDown()
	// knapsack.TestKnapSack3BottomUP()
	// knapsack.TestKnapSack4TopDown()
	// knapsack.TestKnapSack4BottomUp()
	// knapsack.TestKnapSack5BottomUp()
	// knapsack.TestKnapSack6BottomUp()
	// knapsack.TestCoinChangeProblem1WithMemo()
	// knapsack.TestCoinChangeProblem2()
	// knapsack.TestCoinChangeProblem1()
	// knapsack.TestCoinChangeProblem2WithMemo()
	// lcs.TestFindMaxLengthOfLcsWithMemo()
	// lcs.TestFindMaxLengthOfLcs()
	// lcs.TestMinInsertionForPalindrome()
	// lcs.TestFindShortestCommonSuperSequence()
	// lcs.TestPalindromePartitioning()
	// lcs.TestIsScramble()

	// arrays.TestBinarySearch()
	// mcm.TestEggDrop()
	// mcm.TestEggDropBinary()
	// trees.TestDiameterOfBinaryTree()
	// trees.TestHasPathSum()

	// stackTest()

	// trees.TestPathSumII()
	// advent2021.TestGetDepthIncreaseCountByWindow()
	// advent2021.TestGetPowerConsumption()
	advent2021.TestGetLifeSupportRating()
	// 	fmt.Println(a)

	fmt.Println("Hello world")
}

func stackTest() {
	// var stack *list.List

	stack := list.New()

	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)
	stack.PushBack(4)

	for stack.Len() > 0 {
		v := stack.Front()
		fmt.Println(v.Value)
		stack.Remove(v)
	}
}
