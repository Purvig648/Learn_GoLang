package operators

import "fmt"

func Operators() {
	//Arithmetic Operators
	num1, num2 := 10, 2 //assignment
	add := num1 + num2  //addition
	fmt.Println("addition", add)

	sub := num1 - num2
	fmt.Println("subtraction", sub) //subtraction

	mul := num1 * num2
	fmt.Println("multiplication", mul) //multiplication

	div := num1 / num2
	fmt.Println("division", div) //division

	mod := num1 % num2
	fmt.Println("modulos", mod) //modulos

	//Relational Operators
	if num1 == num2 { //equals
		fmt.Println("Success")
	} else if num1 != num2 { //Not Equal
		fmt.Println("Failure")
	} else if num1 > num2 { //greater
		fmt.Println("num1 greater than num2")
	} else if num1 < num2 { //lesser than
		fmt.Println("num1 is lesser than num2")
	} else if num1 >= num2 { //greater than equal to
		fmt.Println("num1 is greater than or equal to num2")
	} else if num1 <= num2 { //lesser than equal to
		fmt.Println("num1 is lesser than or equal to num2")
	}

	// logical operators
	if num1 != num2 && num1 <= num2 { //AND
		fmt.Println("True")
	}

	if num1 != num2 || num1 <= num2 { //OR
		fmt.Println("True")
	}

	if !(num1 == num2) { //NOT
		fmt.Println("True")
	}
}
