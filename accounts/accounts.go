package accounts

import (
	"errors"
	"fmt"
)

// Account로 시작하는 주석이 달려야 export 가능.
type Account struct {
	owner   string // 대문자로 해줘야 public이 됨.
	balance int
}

var errNoMoney = errors.New("Can't withdraw you are poor") // err로 시작하는 변수명이 에러명 정할 때의 컨벤션.

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) { // 저 a Account는 receiver라고 함. class method라고 생각하면 됨.
	a.balance += amount
}

// Balance of your account
func (a *Account) Balance() int { // 별을 붙이지 않으면 복사본을 만들게 됨. 따라서 진짜를 변경하기 위해 별을 붙임.
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string { // 파이썬의 __str__메소드 같은 거.
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
