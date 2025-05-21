package anonymousfunc

import "fmt"

func Print() {
	func() {
		fmt.Println("Welcome") //this is a anonymous function
	}()

	value := func() {
		fmt.Println("Hi Everyone") //can assign an anonymous function to a variable.
	}
	value()

	func(str string) { //passing arguments
		fmt.Println(str)
	}("Working")
}

//passing as parameters
func PrintTwoNumbers(i func(a, b int) int) {
	fmt.Println(i(5, 10))
}

//returning anonymous function
func PrintNum() func(a, b int) int {
	myf := func(a, b int) int {
		return a + b
	}
	return myf
}
