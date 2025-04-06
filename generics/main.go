package main

import "fmt"

func main() {

	value := add(1, 3)
	fmt.Println(value)
}

// This gives us an error because the + operator may not be defined on some data types. Since a and b can
// be any data type including structs, Go throws an error.
// func add(a, b any) {
// 	return a + b
// }

// This is where generics come into play
// Generic type lets you define a generic type which can have concrete data type
// To declare a generic type we use [] before the () of a function
// Here, T is the name of the generic type and int,string,float64 are the concrete values this type can have
// These values are seperated by | also known as a pipe.
func add[T int | string | float64](a, b T) T {
	return a + b
}

// Now that our function can have only string, float64 or int data type, we can use this function without encountering an error
