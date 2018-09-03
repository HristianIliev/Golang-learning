package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

// Anonymous fields
type Person struct {
	string
	int
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// Anonymous structures
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	// Individual values of a field
	emp4 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Printf("Salary: $%d", emp4.salary)

	p := Person{"Naveen", 5}
	fmt.Println(p)
}
