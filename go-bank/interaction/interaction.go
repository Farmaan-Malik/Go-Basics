package interaction

import "fmt"

func InitialAsk(choice *int) {
	Intro(choice)
	fmt.Printf("Enter your choice: ")
	fmt.Scan(choice)
	fmt.Println("You chose", *choice)
}
func Intro(choice *int) {
	if *choice == 0 {
		fmt.Printf("Welcome to Go bank!\nWhat do you want to do today?\n1.Check balance\n2.Deposit Money\n3.Withdraw Money\n4.Exit\n")
	} else {
		fmt.Printf("1.Check balance\n2.Deposit Money\n3.Withdraw Money\n4.Exit\n")
	}
}
