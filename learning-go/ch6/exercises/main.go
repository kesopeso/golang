package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	p := makePerson("John", "Deer", 100)
	fmt.Println(p)
	pP := makePersonPointer("John", "Deer", 100)
	fmt.Println(pP)
}

func makePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func makePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

// why this happens
// slices are implemented as structs with 3 fields
// pointer to data array
// int len
// int cap
// when passed as a parameter, a copy is made
// this copy keeps the pointer address to the data array
// but it creates new values for len and cap
// so original len anc cap from outside stay the same
// which means that items can only be modified and not added!
func exercise2() {
	slice := []string{"one", "two", "three"}
	updateSlice(slice, "four")
	fmt.Println("after update slice", slice)
	growSlice(slice, "five")
	fmt.Println("after grow slice", slice)
}

func updateSlice(sliceToUpdate []string, newValue string) {
	sliceLen := len(sliceToUpdate)
	if sliceLen == 0 {
		return
	}
	sliceToUpdate[sliceLen-1] = newValue
	fmt.Println("in update slice", sliceToUpdate)
}

func growSlice(sliceToGrow []string, newValue string) {
	if sliceToGrow == nil {
		sliceToGrow = make([]string, 1)
	}
	sliceToGrow = append(sliceToGrow, newValue)
	fmt.Println("in grow slice", sliceToGrow)
}

func exercise3() {
	numberOfPersons := 100
	var persons []Person
	// if we comment this out, GC will run a lot of times
	// slice will be copied to a new location, once if exceeds the capacity
	// if we preallocate the capacity, there will be no slice copying
	persons = make([]Person, 0, numberOfPersons)
	for i := 0; i < numberOfPersons; i++ {
		persons = append(persons, Person{"John", "Doe", 25})
	}
	fmt.Println(persons)
}
