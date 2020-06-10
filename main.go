package main

import (
	"fmt"

	"github.com/rkh/bankAccount/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	account.ChangeOwner("nicolas")
	fmt.Println(account)
}
