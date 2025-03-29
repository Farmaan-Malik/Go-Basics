package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	currentBalance, err := readBalance()
	if err != nil {
		fmt.Println("An error occured while fetching account balance.")
		return
	}
	var choice int = 0
	for {
		initialAsk(&choice)
		switch choice {
		case 1:
			fmt.Println("Your account balance is", currentBalance)
		case 2:
			var amount float64
			fmt.Print("Enter Deposit amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Not a valid amount")
			} else {
				currentBalance += amount
				fmt.Println("Money Deposited successfully")
				fmt.Println("Your account balance is", currentBalance)
				writeBalance(currentBalance)
			}
		case 3:
			var amount float64
			fmt.Print("Enter Withdrawal amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Not a valid amount")
			} else if amount <= currentBalance {
				currentBalance -= amount
				fmt.Println("Money Withdrawn successfully")
				fmt.Println("Your account balance is", currentBalance)
				writeBalance(currentBalance)
			} else {
				fmt.Println("Not enough balance in your account")
			}
		case 4:
			fmt.Printf("Thank you for choosing our services.\nGoodbye!\n")
			return
		default:
			fmt.Printf("Please choose a valid option\n")
		}
	}
}

func initialAsk(choice *int) {
	intro(choice)
	fmt.Printf("Enter your choice: ")
	fmt.Scan(choice)
	fmt.Println("You chose", *choice)
}
func intro(choice *int) {
	if *choice == 0 {
		fmt.Printf("Welcome to Go bank!\nWhat do you want to do today?\n1.Check balance\n2.Deposit Money\n3.Withdraw Money\n4.Exit\n")
	} else {
		fmt.Printf("1.Check balance\n2.Deposit Money\n3.Withdraw Money\n4.Exit\n")
	}
}
func writeBalance(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile("Balance.txt", []byte(balanceString), 0644)
}
func readBalance() (balance float64, err error) {
	data, err := os.ReadFile("Balance.txt")
	dataString := string(data)
	balance, err = strconv.ParseFloat(dataString, 64)
	return balance, err

}
