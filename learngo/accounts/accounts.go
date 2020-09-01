package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

//NewAccount creates Account 생성자
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account //리시버는 소문자의 첫글자를 따서 지어야함
//포인터 리시버
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance your Account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account try catch문이 없음
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}

//기본적으로 account라고 했을때 출력되는 문장
