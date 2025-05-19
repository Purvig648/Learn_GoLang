package keywords

import "fmt"

// keywords are reserverd words
// there are 25 keywords in golang

func Keywords() { // func is a keyword
	arr := []int{10, 20, 30}
	for i := 0; i < len(arr); i++ {
		if arr[i]+arr[i+1] > 10 {
			continue
		}
		break
	}
	fmt.Println("Learning Keywords")
}

// for , if, break , continue , package ,import are al keywords
