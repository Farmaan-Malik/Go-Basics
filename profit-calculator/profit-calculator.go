package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	fmt.Println("Your EBT is:", ebt)
	profit := ebt * (1 - taxRate/100)

	// %v is used in Printf to print the value of a variable and %T is used to print the type
	// %f can be used to print floats and if you want to control how many decimal points to print you can use
	// %.2f to print 2 decimal points, %.0f to print no decimal points or 0.3f and so on.
	fmt.Printf("Your profit is: %v \n", profit)
	// we can use `` instead of "" and then add line breaks by simply pressing enter key.
	// `` doesnt recognise \n and would simply print \n on the screen if used.
	ratio := ebt / profit
	fmt.Printf("Ratio of profit and ebt is: %.3f \n", ratio)

	// fmt package also has a Sprintf and Sprintln method that can be used but instead of immediately printing
	// the value to the console, it instead stores it inside of a variable.
	// formattedString := fmt.Sprintf("Your profit is: %v \n", profit)
	// now we can print formattedString using fmt.Println(formattedString) or fmt.Print(formattedString)

}
