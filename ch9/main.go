package main

import (
	bank "GoTheProgrammingLanguage/ch9/bank1"
	"fmt"
)

func main() {
	bank.Deposit(100)
	fmt.Println(bank.Balance())

	withdrawal := bank.Withdraw(99)
	withdrawal2 := bank.Withdraw(99)
	fmt.Println(withdrawal)
	fmt.Println(withdrawal2)
	fmt.Println(bank.Balance())
}
