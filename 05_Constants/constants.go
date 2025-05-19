package constants

import "fmt"

const PI = 3.14 // can be used globally

func Constants() {
	const Data = "Welcome" // const is used to devlare a constant

	fmt.Println(Data)
	//Data cannot be changed
}
