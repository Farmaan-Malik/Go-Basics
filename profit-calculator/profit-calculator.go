package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, taxRate float64

	// fmt.Print("Enter your revenue: ")
	// fmt.Scan(&revenue)
	err := printAndScan("Revenue", &revenue)
	if err != nil {
		panic(err)
	}
	// fmt.Print("Enter your expenses: ")
	// fmt.Scan(&expenses)
	err = printAndScan("Expenses", &expenses)
	if err != nil {
		panic(err)
	}
	// fmt.Print("Enter tax rate: ")
	// fmt.Scan(&taxRate)
	err = printAndScan("Tax Rate", &taxRate)
	if err != nil {
		panic(err)
	}

	// ebt := revenue - expenses
	// fmt.Println("Your EBT is:", ebt)
	// profit := ebt * (1 - taxRate/100)

	// %v is used in Printf to print the value of a variable and %T is used to print the type
	// %f can be used to print floats and if you want to control how many decimal points to print you can use
	// %.2f to print 2 decimal points, %.0f to print no decimal points or 0.3f and so on.

	// fmt.Printf("Your profit is: %v \n", profit)

	// we can use `` instead of "" and then add line breaks by simply pressing enter key.
	// `` doesnt recognise \n and would simply print \n on the screen if used.

	// ratio := ebt / profit
	// fmt.Printf("Ratio of profit and ebt is: %.3f \n", ratio)

	// fmt package also has a Sprintf and Sprintln method that can be used but instead of immediately printing
	// the value to the console, it instead stores it inside of a variable.
	// formattedString := fmt.Sprintf("Your profit is: %v \n", profit)
	// now we can print formattedString using fmt.Println(formattedString) or fmt.Print(formattedString)

	ebt, profit, ratio, error := calculateValues(revenue, expenses, taxRate)
	if error != nil {
		panic(error)
	}
	fmt.Printf("Calculated EBT is %.2f \nCalculated Profit is %.2f \nRatio of EBT and Profit is %.2f\n", ebt, profit, ratio)
}

func printAndScan(text string, inputVar *float64) error {
	fmt.Printf("Enter %v:", text)
	var input float64
	fmt.Scan(&input)
	if input <= 0 {
		err := errors.New("invalid input value")
		return err
	}
	*inputVar = input
	return nil

}

func calculateValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64, err error) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	formattedEbt := fmt.Sprintf("ebt = %.2f\nprofit = %.2f\nratio = %.2f\n", ebt, profit, ratio)
	err = os.WriteFile("Calculations.txt", []byte(formattedEbt), 0644)
	if err != nil {
		return ebt, profit, ratio, err
	}
	return ebt, profit, ratio, nil
}
