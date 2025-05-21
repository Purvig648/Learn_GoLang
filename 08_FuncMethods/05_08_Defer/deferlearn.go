package deferlearn

import "fmt"

func Mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func DeferFunc() {
	fmt.Println("Start to learn Defer")
	fmt.Println("Start")

	// Multiple defer statements
	// Executes in LIFO order
	defer fmt.Println("End")
	defer Mul(34, 56)
	defer Mul(10, 10)
}
