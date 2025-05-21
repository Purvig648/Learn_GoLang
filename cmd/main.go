package main

import (
	methods "Learn_Golang/08_FuncMethods/06_08_Methods"
	"fmt"
)

//HelloWorld "Learn_Golang/01_Hello_world"
//identifiers "Learn_Golang/02_Identifiers"
//keywords "Learn_Golang/03_KeyWords"
//datatypes "Learn_Golang/04_DataTypes"
//constants "Learn_Golang/05_Constants"
//operators "Learn_Golang/06_Operators"
//functions "Learn_Golang/08_FuncMethods/01_08_functions"

func main() {

	//HelloWorld.Start()        //printing hello world
	//identifiers.Identifiers() //identifiers
	//keywords.Keywords()   //keywords
	//datatypes.DataTypes()  //datatypes
	//constants.Constants()  //constatnts
	//operators.Operators()  //Go Operators

	// num1 := 10
	// num2 := 20
	// sum := functions.Addition(num1, num2) //call by value
	// fmt.Println(sum)
	// sub := functions.Subtraction(&num2, &num1)
	// fmt.Println(sub) //call by reference

	//methods
	a := methods.Person{
		Name: "Purvi",
		Id:   1,
	}

	a.Display()

	a1 := methods.Number(8)
	mulData := a1.Multiply()
	fmt.Println(mulData)

	newName := "Nithin"
	newData := methods.Person{
		Name: "Purvi",
	}
	newData.ChangeName(newName)

}
