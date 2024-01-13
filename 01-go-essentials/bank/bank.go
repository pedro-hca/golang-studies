package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/pedro-hca/go-studies/01-go-essentials/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		panic("Cant continue, sorry")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			CheckBalance(accountBalance)

		} else if choice == 2 {
			var amount float64 = 0
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("invalid amount")
				continue
			}
			accountBalance = DepositMoney(amount, accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		} else if choice == 3 {
			var withdrawAmount float64 = 0
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 || withdrawAmount > accountBalance {
				fmt.Println("invalid withdrawAmount")
				continue
			}
			accountBalance = WithdrawMoney(withdrawAmount, accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		} else {
			fmt.Println("Exit")
			break
		}

		fmt.Println("Your choice:", choice)
	}
}

func CheckBalance(accountBalance float64) {
	fmt.Println("Your balance is", accountBalance)
}
func DepositMoney(amount, accountBalance float64) float64 {
	newBalance := accountBalance + amount
	fmt.Println("Balance updated! New amount:", newBalance)
	return newBalance
}
func WithdrawMoney(amount, accountBalance float64) float64 {
	newBalance := accountBalance - amount
	fmt.Println("Balance updated! New amount:", newBalance)
	return newBalance
}
