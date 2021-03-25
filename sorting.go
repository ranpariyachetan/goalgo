package main

import "fmt"

func performBubbleSort(arr []int, popTop bool) {

	isDone := false

	if popTop {
		// Approach where smallest element is popped up at the top
		for i := 0; i < len(arr)-1; i++ {
			isDone = true
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					arr[i] += arr[j]
					arr[j] = arr[i] - arr[j]
					arr[i] = arr[i] - arr[j]
					isDone = false
				}
			}

			// fmt.Println(arr)
			if isDone {
				break
			}
		}
	} else {
		// Approach where largest element is pushed to the bottom
		for i := 0; i < len(arr); i++ {
			isDone = true
			for j := 0; j < len(arr)-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j] += arr[j+1]
					arr[j+1] = arr[j] - arr[j+1]
					arr[j] = arr[j] - arr[j+1]
					isDone = false
				}
			}
			if isDone {
				break
			}
		}
	}
}

func performInsertionSort(arr []int) {
	var i, j, key int
	n := len(arr)
	for i = 1; i < n; i++ {
		key = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j -= 1
		}

		arr[j+1] = key

	}
}

func performSelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[minIndex] += arr[i]
			arr[i] = arr[minIndex] - arr[i]
			arr[minIndex] = arr[minIndex] - arr[i]
		}

		fmt.Println(arr)
	}

	return arr
}

func performMergeSort(arr []int) {
	n := len(arr)
	low := 0
	high := n - 1

	if n == 1 {
		return
	}

	sort(arr, low, high)
}

func sort(arr []int, low int, high int) {
	// fmt.Println(arr[low : high+1])
	if low < high {
		mid := low + (high-low)/2
		sort(arr, low, mid)
		sort(arr, mid+1, high)
		merge(arr, low, mid, high)
	} else {
		return
	}
}

func merge(arr []int, low, mid, high int) {
	var temp = make([]int, len(arr))
	var l1, l2, i int

	l1 = low
	l2 = mid + 1
	i = low
	for ; l1 <= mid && l2 <= high; i++ {
		if arr[l1] <= arr[l2] {
			temp[i] = arr[l1]
			l1++
		} else {
			temp[i] = arr[l2]
			l2++
		}
	}

	for l1 <= mid {
		temp[i] = arr[l1]
		i++
		l1++
	}

	for l2 <= high {
		temp[i] = arr[l2]
		i++
		l2++
	}

	for i = low; i <= high; i++ {
		arr[i] = temp[i]
	}
}
