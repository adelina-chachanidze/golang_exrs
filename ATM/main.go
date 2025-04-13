package main

import "fmt"

func main () {
	balance := 1000


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
	}

	fmt.Println("Your choice:", choice)

}