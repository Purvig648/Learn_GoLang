package arrays

import "fmt"

func CalculateScore(arr [5]int) { // passing array to a function
	//iterating an array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// copying a array
	array := arr
	fmt.Println(array)
}

func Array() {
	arr := [5]int{} //shortend declaration of array
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + i + 1 // assigning vale to an array
	}

	fmt.Println(arr)
}
