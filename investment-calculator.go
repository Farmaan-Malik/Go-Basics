// package main tells the go compiler that this package is the entry point of the app
package main

// "fmt" is a package that contains the Print function. "fmt" is known as format package
import (
	"fmt"
	"math"
)

// constant is a data container whose value, unlike var, can't be changed
// constant is declared outside of a function making its scope global
const inflationRate = 6.6

// main function is the entry point of the app and therefore, is executed first
// Only one main function is allowed per package
func main() {

	// explicit type declaration of a variable
	var investmentAmount, years float64 = 1000, 10
	// or
	// investmentAmount, years, expectedReturnRate  := 1000.0, 10.0, 5.5
	// this way we can declare all our variables in one line without explicit type declaration

	// inferred type assignment
	expectedReturnRate := 5.5

	// Scan function is used to take user input. &investmentAmount is a pointer to investmentAmount variable
	// this function will populate the investmentAmount variable with the value typed in by the user
	// Limitation: You can't (easily) fetch multi-word input values.
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years: ")
	fmt.Scan((&years))
	fmt.Print("Expected Return Rate: ")
	fmt.Scan((&expectedReturnRate))

	futureValue, realFutureValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print("Future Value-> ")
	fmt.Println(futureValue)

	fmt.Print("Real Future Value-> ")
	fmt.Println(realFutureValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// alternatively, you can give the return values names which would then create variables of that name and type
// this means you no longer need to declare fv and rfv with := , instead you can just use =
// this also means that you can simply write return and go would know what variables or values to return

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }
