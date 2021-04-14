package main

import (
	"fmt"

	"github.com/KangJunewoo/hi-go/mydict"
)

func main() {

	// bank project
	// account := accounts.NewAccount("nico")
	// account.Deposit(10)
	// fmt.Println(account.Balance())
	// err := account.Withdraw(20)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// fmt.Println(account)

	// Dictionary project
	// dictionary := mydict.Dictionary{"hello": "world"}
	// fmt.Println(dictionary)
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

}
