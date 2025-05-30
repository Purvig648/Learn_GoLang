package switchStatements

import "fmt"

func LearnSwitch() {
	//Expression switch
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a proper data")
	}

	//type switch
	var day1 interface{} = 4
	switch v := day1.(type) {
	case int:
		switch v {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		default:
			fmt.Println("Invalid day")
		}
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
