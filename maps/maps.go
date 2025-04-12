package main

import "fmt"

// type alias for a map of key= string and value=string
type stringMap map[string]string

// type alias for an array of strings

type stringArr []string

func main() {
	// Maps let you store data in a key-value pair (or dictionary type).
	// The keys and the values can have any type e.g a struct, a string, a number etc

	websites := map[string]string{"Google": "https://www.google.com", "LinkedIn": "https://www.linkedIn.com"}
	fmt.Println(websites)
	websites["AWS"] = "https://www.aws.com"
	fmt.Println(websites)
	// delete is a prebuilt function in go that allows you to delete an element with the specified key from a map
	delete(websites, "Google")
	fmt.Println(websites)
	// make is used to allot memory to an array or a map.
	// the arguments that make() takes are:
	// type, size , capacity
	// Make() on takes type and size as arguments in case of a map
	books := make(stringArr, 2, 5)
	fmt.Println(books)
	books = append(books, "three", "four", "five")
	books = append(books, "six")
	// for loop to go through all the elements of a map or an array
	for i, v := range books {
		fmt.Println("Key:", i)
		fmt.Println("Value:", v)
	}
}
