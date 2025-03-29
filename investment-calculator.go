// package main tells the go compiler that this package is the entry point of the app
package main

// "fmt" is a package that contains the Print function. "fmt" is known as format package
import (
	"fmt"
	"math"
)

// main function is the entry point of the app and therefore, is executed first
// Only one main function is allowed per package
func main() {
	// constant is a data container whose value, unlike var, can't be changed
	const inflationRate = 6.6

	// explicit type declaration of a variable
	var investmentAmount, years float64 = 1000, 10
	// or
	// investmentAmount, years, expectedReturnRate  := 1000.0, 10.0, 5.5
	// this way we can declare all our variables in one line without explicit type declaration

	// inferred type assignment
	expectedReturnRate := 5.5

	// Scan function is used to take user input. &investmentAmount is a pointer to investmentAmount variable
	// this function will populate the investmentAmount variable with the value typed in by the user
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years: ")
	fmt.Scan((&years))
	fmt.Print("Expected Return Rate: ")
	fmt.Scan((&expectedReturnRate))

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print("Future Value-> ")
	fmt.Println(futureValue)

	fmt.Print("Real Future Value-> ")
	fmt.Println(realFutureValue)
}
