package practice

import (
	"fmt"
	"log"

	"github.com/KangJunewoo/hi-go/accounts"
)

func bank() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

}
