package main

import "fmt"

func main () {
	var balance float64 = 1000.00
	fmt.Println("Welcome to the ATM!")

	for {
	fmt.Println("Please select the action")
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
				continue
			}
		balance += deposit
		fmt.Println("\nDeposit complete!\nYour new balance:", balance)

	} else if choice == 3 {
		fmt.Println("\nEnter your withdrawl amount: ")
		var withdraw float64
		fmt.Scan(&withdraw)
			if withdraw <=0 {
				fmt.Println("\nInvalid amount. Minimal withdrawl: 0.01")
				continue
			}

			if withdraw > balance {
				fmt.Println("\nNot enough to withdraw, availabe balance:", balance)
				continue
			} 
			
		balance -= withdraw
		fmt.Println("\nWithdrawl complete!\nYour new balance:", balance)

	} else {
		fmt.Println("Thank you for using our ATM! Come back soon!")
		break
	}
	}
	
}