package decisionmaking

import "fmt"

//decision making statements are if , if- else , Nested if ,if -else if ladder

func DecisionMaking() {
	Number := 10
	if Number >= 1 && Number < 10 { //if statement
		fmt.Println("it is a single digit")
	} else if Number >= 10 && Number < 100 { //if else if ladder
		fmt.Println("It is double digit")
	} else { //else statement
		if Number == 0 { // Nested if
			fmt.Println("It is a zero")
		}
		fmt.Println("It is not considred digit for the ")
	}
}
