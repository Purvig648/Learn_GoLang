package datatypes

import "fmt"

type User struct { // struct data type
	Name string
	UID  int
}

func DataTypes() { // func is a reference type
	var number int // variable declaration
	var floatnumber float32
	var unumber uint

	number = 10
	floatnumber = 3.6
	unumber = 10

	fmt.Println(number, floatnumber, unumber) // Number types

	StringData := "Hello Everyone, Great Learning!" //shortend declaration

	fmt.Println(StringData) // String Type

	myBooleanLearning := true

	fmt.Println(myBooleanLearning)

	//Aggreagte Types
	var NumberArray = [5]int{} //Arrays
	fmt.Println(NumberArray)

	//Reference Types
	var NumberSlice = []int{} //slices
	fmt.Println(NumberSlice)

	mapData := make(map[int]bool) //maps
	fmt.Println(mapData)

}
