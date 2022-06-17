package factory

import "fmt"

func (p Person) PrintInfo() {
	fmt.Printf("My name is %s and age is %d\n", p.Name, p.Age)
}
