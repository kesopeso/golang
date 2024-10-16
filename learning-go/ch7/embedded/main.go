package main

import "fmt"

type Employee struct {
	Name string
	ID   int
}

type Manager struct {
	Employee
	Reports []Employee
}

func (e Employee) Description() string {
	return fmt.Sprintf("Name: %s, ID: %d", e.Name, e.ID)
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bruv",
			ID:   1,
		},
		Reports: []Employee{},
	}

	fmt.Println(m.Description())

	// var eFalse Employee = m // this is incorrect, embedded field must be accessed directly in order to use it as Employee type
	var eTrue Employee = m.Employee
	fmt.Println(eTrue)

	// different way of converting methods to functions
	descFunc := m.Employee.Description
	fmt.Print(descFunc())

	descAnother := Employee.Description
	fmt.Println(descAnother(eTrue))
}
