package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	GetMessage() string
}

func GetMessageI(printInfo PrintInfo) {
	fmt.Println(printInfo.GetMessage())
}

func (fullTimeEmployee FullTimeEmployee) GetMessage() string {
	return "FullTimeEmployee"
}

func (temporaryEmployee TemporaryEmployee) GetMessage() string {
	return "TemporaryEmployee"
}

func GetMessage(person Person) {
	fmt.Printf("%s with age %d", person.name, person.age)
}

func main() {
	fullTimeEmployee := FullTimeEmployee{}
	fullTimeEmployee.id = 1
	fullTimeEmployee.name = "Richi"
	fullTimeEmployee.age = 34
	fmt.Println(fullTimeEmployee)

	temporaryEmployee1 := TemporaryEmployee{}
	GetMessageI(temporaryEmployee1)
	GetMessageI(fullTimeEmployee)
}
