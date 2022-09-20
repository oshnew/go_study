package main

import (
	"fmt"
)

//binary search를 Github Copilot으로 만들어자
func main() {
	fmt.Print("\n\n\nhello world")
	fmt.Println(" welcome!")

	//call binarySearch function
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 5
	index := binarySearch(arr, key)
	println("index of key is ", index)

	//call binarySearchRecursive function
	index = binarySearchRecursive(arr, key, 0, len(arr)-1)
	println("index of key is ", index)

}

func binarySearchRecursive(arr []int, key int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == key {
		return mid
	} else if arr[mid] > key {
		return binarySearchRecursive(arr, key, low, mid-1)
	} else {
		return binarySearchRecursive(arr, key, mid+1, high)
	}
}

// Language: go
//binary search by golang
//binary search with sorted array
func binarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
