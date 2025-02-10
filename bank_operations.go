package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	balanceFile = "balance.txt"
	welcome     = "Welcome to Go Bank!"
)

func getBalanceFile() float64 {
	data, _ := os.ReadFile(balanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func deposit(accountbalance *float64) {
	var depositAmount float64
	fmt.Print("Your deposit: ")
	fmt.Scan(&depositAmount)

	//validator rules
	if depositAmount <= 0 {
		fmt.Println("Invalid amount! Must be greater than 0.")
		return
	}

	*accountbalance += depositAmount
	writeBalanceToFile(*accountbalance)
	fmt.Println("Balance updated! New amount: ", *accountbalance, "R$")
}

func withdraw(accountbalance *float64) {
	var withdrawAmount float64
	fmt.Print("Your whithdraw: ")
	fmt.Scan(&withdrawAmount)

	//validator rules
	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount! Must be greater than 0.")
		return
	}
	//validator rules
	if withdrawAmount > *accountbalance {
		fmt.Println("Invalid amount! You can't withdraw more than you have")
		return
	}

	*accountbalance -= withdrawAmount
	writeBalanceToFile(*accountbalance)
	fmt.Println("Balance updated! New amount: ", *accountbalance, "R$")
}

func main() {

	var accountbalance = getBalanceFile()
	var choice int

	fmt.Println(welcome)

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Whithdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your amount: ", accountbalance, "R$")
		case 2:
			deposit(&accountbalance)
		case 3:
			withdraw(&accountbalance)
		case 4:
			fmt.Println("Exit your bank")
			return
		default:
			println("Enter a valid value to proceed with the service")
		}
	}
}
