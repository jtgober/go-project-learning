package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------------")
		// panic(err)
	}

	fmt.Println("Welcome to the Bank of Golang")
	fmt.Println(randomdata.PhoneNumber())
	for {

		presentOptions()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)
		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("How much are you depositing?: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount entered. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount // same as accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated. New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("How much are you withdrawing?: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount entered. Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Withdraw is greater than the account balance.")
				continue
			}

			accountBalance -= withdrawAmount // same as accountBalance = accountBalance - withdrawAmount
			fmt.Println("Balance updated. New amount: ", accountBalance)
		default:
			fmt.Println("Finished. Current account balance: ", accountBalance)
			fmt.Println("Thanks for choosing our bank")
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			return
		}

	}

}
