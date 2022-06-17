package factory

import "fmt"

type PersonInterface interface {
	PrintInfo2()
}

func (p Person) PrintInfo2() {
	fmt.Printf("My name is %s and age is %d\n", p.Name, p.Age)
}

func NewPerson2(name string, age int) PersonInterface {
	return Person{
		Name: name,
		Age:  age,
	}
}
