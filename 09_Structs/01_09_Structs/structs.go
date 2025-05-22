package structs

import "fmt"

//struct declaration with struct keyword
type Address struct {
	Name    string
	Street  string
	City    string
	State   string
	Pincode int
}

//Nested Struct
type Employee struct {
	Name    string
	Emp_ID  int
	Address Address
}

func StructLearn() {
	var a Address //declaring a struct with zero values
	fmt.Println(a)

	a1 := Address{ //initialzing a struct
		Name:    "#78",
		Street:  "Unknown",
		City:    "Bangalore",
		State:   "Karnataka",
		Pincode: 89,
	}
	fmt.Println(a1)

	//accessing a struct
	fmt.Println("Prinring the street", a.Street)

	//pointeers to a struct
	add := &Address{
		Name:   "CityTest",
		Street: "HUMAN CITY",
		City:   "CITY",
		State:  "KVDTB",
	}
	fmt.Println(add)

	//accessing nested struct
	EmpData := Employee{
		Name:   "XYZ",
		Emp_ID: 12,
		Address: Address{
			Name:    "Unknow city",
			Street:  "unknown street",
			City:    "unknown city1",
			State:   "unknown state",
			Pincode: 123450,
		},
	}
	fmt.Println(EmpData)
}
