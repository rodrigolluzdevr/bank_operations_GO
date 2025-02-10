package main

import "fmt"

func main() {

	accountbalance := 6500.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Whithdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your amount: ", accountbalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Your deposit: ")
		fmt.Scan(&depositAmount)

		//validator rules
		if accountbalance <= 0 {
			fmt.Println("Invalid amount! Must be greater than 0.")
			return
		}

		accountbalance += depositAmount
		fmt.Println("Balance updated! New amount: ", accountbalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Your whithdraw: ")
		fmt.Scan(&withdrawAmount)

		//validator rules

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount! Must be greater than 0.")
		}

		if withdrawAmount < accountbalance {
			fmt.Println("Invalid amount! You can't withdraw more than you have")
		}

		accountbalance -= withdrawAmount
		fmt.Println("Balance updated! New amount: ", accountbalance)
	} else {
		fmt.Println("Exit your bank")
	}
}
