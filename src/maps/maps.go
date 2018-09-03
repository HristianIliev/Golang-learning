package main

import "fmt"

func main() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary contents:", personSalary)

	// Check if a key exists
	value, ok := personSalary["joe"]
	if ok == true {
		fmt.Println("Salary of", "joe", "is", value)
	} else {
		fmt.Println("joe", "not found")
	}

	// Iterating
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	// Deleting
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)

	// Length of the map
	fmt.Println("Length of the map", len(personSalary))
}
