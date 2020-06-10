package accounts

//Account struct
type Account struct {
	owner   string
	balance int
}

//NewAccount is function
//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit is Method of Account struct
//Deposit receiver는 struct의 첫글자를 따서 소문자로 해야함(ex a Account)
//Deposit x amount on your account
func (a Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int {
	return a.balance
}
