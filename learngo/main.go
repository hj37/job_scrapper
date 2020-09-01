package main

import (
	"fmt"

	"github.com/hj37/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition:", hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// definition, err := dictionary.Search("second")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	// account := accounts.NewAccount("nico")
	// account.Deposit(10)
	// // err := account.Withdraw(20)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // } //err를 체크하도록 강제함
	// fmt.Println(account)
}
