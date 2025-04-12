package main

import "fmt"

type Products struct {
	title string
	id    int
	price int
}

func New(title string, id, price int) *Products {
	return &Products{
		title: title,
		id:    id,
		price: price,
	}
}

func main() {
	var hobbies [3]string = [3]string{"Coding", "Drawing", "Gaming"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	firstTwo := hobbies[:2]
	fmt.Println("First Two ", firstTwo)
	firstTwo = firstTwo[0:3]
	fmt.Println("Modified First Two ", firstTwo)
	var goals []string = []string{"Learn", "Create"}
	fmt.Println(goals)
	goals[1] = "MakeProjects"
	fmt.Println(goals)
	goals = append(goals, "Job")
	fmt.Println(goals)
	var products []Products = []Products{{title: "Chair", id: 1, price: 160}, {title: "Table", id: 2, price: 400}}
	fmt.Println(products)
	var newProduct *Products = New("Keyboard", 3, 290)
	products = append(products, *newProduct)
	fmt.Println(products)
	newProduct = New("Mouse", 3, 120)
	products = append(products, *newProduct)
	fmt.Println(products)
	brandedProducts := []Products{*New("Razer", 4, 290), *New("Nixon", 5, 390), *New("Gaia", 6, 490)}
	// Spread operator to spread the elements of an array/slice/map
	products = append(products, brandedProducts...)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
