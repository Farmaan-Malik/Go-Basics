package main

import (
	"fmt"

	"example.com/go-bank/customIo"
	"example.com/go-bank/interaction"
)

const filename = "Balance.txt"

func main() {
	currentBalance, err := customIo.ReadContent(filename)
	if err != nil {
		fmt.Println(err)
	}
	var choice int = 0
	for {
		interaction.InitialAsk(&choice)
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
				err := customIo.WriteContent(currentBalance, filename)
				if err != nil {
					fmt.Print((err))
				}
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
				err := customIo.WriteContent(currentBalance, filename)
				if err != nil {
					fmt.Print((err))
				}
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
