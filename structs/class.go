package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (employee *Employee) GetIdEmployee() int {
	return employee.id
}

func (employee *Employee) GetNameEmployee() string {
	return employee.name
}

func (employee *Employee) SetIdEmployee(id int) {
	employee.id = id
}

func (employee *Employee) SetNameEmployee(name string) {
	employee.name = name
}

func newEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{id: id, name: name, vacation: vacation}
}

func main() {

	myemployee := Employee{id: 1, name: "Richi"}

	myemployee2 := Employee{}

	fmt.Println(myemployee2)

	myemployee2.SetIdEmployee(5)
	myemployee2.SetNameEmployee("Ivon")

	fmt.Println("Hola Mundo", myemployee, myemployee2)

	myemployee3 := Employee{id: 4, name: "Richard", vacation: true}

	fmt.Println("myemployee3 ", myemployee3)

	myemployee4 := new(Employee)

	fmt.Println("myemployee4 ", *myemployee4)

	myemployee5 := newEmployee(6, "Richi", true)

	fmt.Println("myemployee5 ", *myemployee5)

}
