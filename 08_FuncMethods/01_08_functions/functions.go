package functions

//functions are declared with func keyword
func Addition(num1, num2 int) int { //num1 and num2 are parameters
	sum := num1 + num2
	return sum //return type
}

func Subtraction(num1, num2 *int) int { //referencing
	sub := *num1 - *num2
	return sub
}
