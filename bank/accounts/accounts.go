package accounts

import (
	"errors"
	"fmt"
)

var errCantWithDraw error = errors.New("引き出しできません,残高不足")

//Account struct
type Account struct {
	//fieldもlowerCaseとaccessできない。
	//UpperCase : Public, LowerCase : Private
	owner   string
	balance int
}

//goにはconstructorがない為にfunctionかstructを作ってconstructorの役割を果たさせる。

//NewAccount of Account's Constructor
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//methodの作り方
// a <- receiver, Account <- aのtypeを付けるだけで完成
// receiverはどんな名前でもいいがconventionとしてはstructの頭文字の小文字で書く
// pointer receiver => (a *Account) 複製じゃなくてDepositを呼び出した本物を使え

//Deposit method of Account
func (a *Account) Deposit(amount int) {
	fmt.Println(&a)
	a.balance += amount
	fmt.Println(&a.balance, "a.balance")
}

//Balance method of Account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw method of Account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errCantWithDraw
	}
	a.balance -= amount
	//errorは2つのvalueを持っている error, nil
	return nil
}
