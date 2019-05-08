package sorts

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arr)
	BubbleSort(arr)
	fmt.Println("  Sorted array: ", arr)

	arr = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arr)
	BubbleSort2(arr)
	fmt.Println("  Sorted array: ", arr)
}

func BubbleSort2(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				Swap(arr, i, j)
			}
		}
	}
}

func BubbleSort(arr []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				Swap(arr, i, i+1)
				swapped = true
			}
		}
	}
}
func Swap(arr []int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}
