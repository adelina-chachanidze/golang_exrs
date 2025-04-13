package main

import (
	"fmt"
)

func main() {
	var balance float64 = 1000.00

	fmt.Println("Welcome to the ATM!")

	for {

		fmt.Println("\nPlease select the action")
		fmt.Println("1. Check your balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Type 1-4 to choose: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance:", balance)

		case 2:
			fmt.Println("Enter your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Invalid amount. Minimal deposit: 0.01")
				continue
			}

			balance += deposit
			fmt.Println("Deposit complete!\nYour new balance:", balance)

		case 3:
			fmt.Println("Enter your withdrawal amount: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid amount. Minimal withdrawal: 0.01")
				continue
			}

			if withdraw > balance {
				fmt.Println("Not enough to withdraw, available balance:", balance)
				continue
			}

			balance -= withdraw
			fmt.Println("Withdrawal complete!\nYour new balance:", balance)

		default:
			fmt.Println("Thank you for using our ATM! Come back soon!")
			return
		}
	}
}