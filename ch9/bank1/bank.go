// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan withdrawal) // withdraw certain funds possibility

type withdrawal struct {
	amount int
	withdrawals chan bool
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	isWithdrawable := make(chan bool)
	withdrawals <- withdrawal{amount, isWithdrawable}

	return <-isWithdrawable
}

func teller() {
	var balance int // balance is confined to teller goroutine

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case wd := <- withdrawals:
			if wd.amount < balance {
				balance -= wd.amount
				wd.withdrawals <- true
			} else {
				wd.withdrawals <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
