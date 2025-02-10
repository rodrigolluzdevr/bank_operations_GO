package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	balanceFile = "balance.txt"
	welcome     = "Welcome to Go Bank!"
)

func getBalanceFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		err := errors.New("balance file not found. A new one will be created with a balance of 0")
		return 0, err
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		err := errors.New("failed to parse stored balance value")
		return 0, err
	}

	return balance, err
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

	var accountbalance, err = getBalanceFile()
	if err != nil {
		fmt.Println("An error occurred:")
		fmt.Println(err)
		fmt.Println("----------------------------------")
	}

	var choice int

	fmt.Println(welcome)

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Whithdraw money")
		fmt.Println("4. Exit")
		fmt.Println("----------------------------------")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		fmt.Println(" ")

		switch choice {
		case 1:
			fmt.Println("Your amount: ", accountbalance, "R$")
			fmt.Println("----------------------------------")
		case 2:
			deposit(&accountbalance)
			fmt.Println("----------------------------------")
		case 3:
			withdraw(&accountbalance)
			fmt.Println("----------------------------------")
		case 4:
			fmt.Println("Exit your bank")
			return
		default:
			println("Enter a valid value to proceed with the service")
			fmt.Println("----------------------------------")
		}
	}
}
