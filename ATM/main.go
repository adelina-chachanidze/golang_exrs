package main

import "fmt"

func main () {
	var balance float64 = 1000.00


	fmt.Println("Welcome to the ATM!\nPlease select the action")
	fmt.Println("\n1. Check your balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("\nType 1-4 to choose: ")
	fmt.Scan(&choice)

	//toCheckBalance := choice == 1

	//= is for assing a value to var. == for check for the quality
	if choice == 1 {
		fmt.Println("Your balance:", balance)

	} else if choice == 2 {
		fmt.Println("Enter your deposit: ")
		var deposit float64
		fmt.Scan(&deposit)
			if deposit <=0 {
				fmt.Println("Invalid amount. Minimal deposit: 0.01")
				return
			}
		balance += deposit
		fmt.Println("Deposit complete!\nYour new balance:", balance)

	} else if choice == 3 {
		fmt.Println("Enter your withdrawl amount: ")
		var withdraw float64
		fmt.Scan(&withdraw)
			if withdraw <=0 {
				fmt.Println("Invalid amount. Minimal withdrawl: 0.01")
				return
			}

			if withdraw > balance {
				fmt.Println("Not enough to withdraw, availabe balance:", balance)
				return
			} 
			
		balance -= withdraw
		fmt.Println("Withdrawl complete!\nYour new balance:", balance)

	} else {
		fmt.Println("Thank you for using our ATM! Come back soon!")
	}
}