package main

import (
	"fmt"

	"github.com/const-takeo/learning_go/bank/accounts"
)

func main() {
	account := accounts.NewAccount("Yun")
	fmt.Println(&account)
	account.Deposit(99999)
	fmt.Println(account.Balance())
	err := account.Withdraw(100000)
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
