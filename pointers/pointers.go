package main

import "fmt"

func main() {
	age := 40
	// age passed to Println function is the copy of the local variable "age"
	fmt.Println(age)
	// &age is the address at which the variable "age" is stored in memory
	// therefore, what we are passing here is a pointer to the memory location of variable "age"
	changeAgeToAdultyears(&age)
	fmt.Println(age)
}

// this function accepts a pointer to an int as a parameter.
func changeAgeToAdultyears(number *int) {
	// arguments are passed by value in Go, therefore a local variable named number is created inside this function
	// this variable holds the memory address of the age variable
	// if we print *number, it prints the value held at the address that is stored inside of this "number" variable
	// if we print number, it prints the value of the "age" variable passed to the function
	// if we print &number, it will print the address of this local variable named "number".

	fmt.Println(*number)
	*number = *number - 18
}
