package methods

import "fmt"

//method with struct reciever
type Person struct {
	Name string
	Id   int
}

func (p Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Id)
}

//method with non struct reciever
type Number int

func (n Number) Multiply() Number {
	multiplicationData := n * n
	return multiplicationData
}

//method with pointer reciever
func (p *Person) ChangeName(newName string) {
	p.Name = newName
}
