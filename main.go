package main

import (
	"fmt"

	"github.com/rkh/bankAccount/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
