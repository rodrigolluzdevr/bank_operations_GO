package main

import "fmt"

func main() {

	var accountbalance float64

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
		accountbalance += depositAmount
		fmt.Println("Balance updated! New amount: ", accountbalance)
	} else if choice == 3 {
		var withdraw float64
		fmt.Print("Your withdraw: ")
		fmt.Scan(&withdraw)
		accountbalance -= withdraw
		fmt.Println("Balance updated! New amount: ", accountbalance)
	} else {
		fmt.Println("Exit your bank")
	}
}
