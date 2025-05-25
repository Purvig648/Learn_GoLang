package slices

import (
	"fmt"
	"sort"
)

func Learnslices() {
	arr := [5]int{10, 20, 30, 40} // array
	fmt.Println("Priniting array", arr)
	newArr := arr[1:2] // slice using an array
	fmt.Println("Printing Slice", newArr)

	new2Arr := []int{2, 3, 4}          // slice declaration using short end
	new2Arr = append(new2Arr, 6, 7, 8) // appending data to slice
	fmt.Println("Printing Appended slice", new2Arr)

	fmt.Println("length of slice", len(new2Arr))
	fmt.Println("Capacity of slice", cap(new2Arr))

	var newSlice = []string{"Learn", "Go"} //one more way of slice declaration
	fmt.Println("New Slice decalaration data", newSlice)

	var newMakeSlice = make([]int, 4, 7) //type ,len ,cap
	fmt.Println("Slice using make", newMakeSlice)
	fmt.Println("Len of new slice", len(newMakeSlice)) //4
	fmt.Println("Capacity", cap(newMakeSlice))         //7

	//iterate over a slice
	for i := 0; i < len(newMakeSlice); i++ {
		fmt.Println("priniting slice DAta", newMakeSlice[i])
	}

	//using range
	for index, value := range newMakeSlice {
		fmt.Println("Priniting index and value", index, value)
	}

	//multi-Dimensional slice
	var newSliceDim = [][]int{
		{10, 20},
	}
	fmt.Println("Priniting multidimenional slice", newSliceDim)

	fmt.Println("Before Sorting", new2Arr)
	//sorting a slice
	sort.Ints(new2Arr)
	fmt.Println("After sorting", new2Arr)
}
