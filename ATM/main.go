package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func getBalanceFromFile () (float64, error) {
	data, err :=os.ReadFile(balanceFile)

	//nil - absence of a useful value
	if err != nil {
		errors.New("Failed to read file")
		return 0, err
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		errors.New("Failed to parse the data")
	}

	return balance, nil
}

func writeBalanceToFile (balance float64) {
	balanceText := fmt.Sprint(balance)
	//0644: owner of the file can read and write, and others only read
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func main() {
	var balance, err  = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error:", err)
		return	
	}

	fmt.Println("Welcome to the ATM!")

	for {

		fmt.Println("\nPlease select the action")
		fmt.Println("1. Check your balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\nYour balance:", balance)

		case 2:
			fmt.Println("\nEnter your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Invalid amount. Minimal deposit: 0.01")
				continue
			}

			balance += deposit
			fmt.Println("\nDeposit complete!\nYour new balance:", balance)
			writeBalanceToFile(balance)

		case 3:
			fmt.Println("\nEnter your withdrawal amount: ")
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
			fmt.Println("\nWithdrawal complete!\nYour new balance:", balance)
			writeBalanceToFile(balance)
		
		case 4: 
			fmt.Println("\nThank you for using our ATM! Come back soon!")
			return


		default:
			fmt.Println("\nInvalid input, please choose 1-4")
		}
	}
}