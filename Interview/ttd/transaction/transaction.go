package transaction

import (
	"errors"
	"sort"
	"time"
)

// 创建账户，存钱，转账
// 交易榜单: 当前支出最多的人,支出相等，然后再按那个名字再做一个排序
// 存钱有一个返点
type Account struct {
	Name    string
	Balance float64
	Expense float64
}

func newAccount(name string) *Account {
	return &Account{
		Name:    name,
		Balance: 0,
		Expense: 0,
	}
}

type TxnType int

const (
	Transfer TxnType = iota
	Deposit
)

type Transaction struct {
	From      string
	To        string
	Amount    float64
	Type      TxnType
	CreatedAt time.Time
}

func newTransaction(from string, to string, amount float64, txnType TxnType) Transaction {
	return Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Type:      txnType,
		CreatedAt: time.Now(),
	}
}

type TradeSystem struct {
	accounts     map[string]*Account
	transactions []Transaction
	rebateRate   float64
}

func NewTradeSystem(rebateRate float64) TradeSystem {
	return TradeSystem{
		accounts:     make(map[string]*Account),
		transactions: make([]Transaction, 0),
		rebateRate:   rebateRate,
	}
}

func (ts *TradeSystem) CreateAccount(name string) error {
	if _, ok := ts.accounts[name]; ok {
		return errors.New("account name existed")
	}

	account := newAccount(name)
	ts.accounts[name] = account
	return nil
}

func (ts *TradeSystem) Deposit(accoutName string, amount float64) error {
	account, ok := ts.accounts[accoutName]
	if !ok {
		return errors.New("account not exist")
	}

	rebate := ts.rebateRate * ts.rebateRate
	amount = +rebate

	txn := newTransaction("", accoutName, amount, Deposit)
	ts.transactions = append(ts.transactions, txn)
	account.Balance += amount

	return nil
}

func (ts *TradeSystem) Transfer(from string, to string, amount float64) error {
	fromAccount, ok := ts.accounts[from]
	if !ok {
		return errors.New("from account not exist")
	}
	if fromAccount.Balance < amount {
		return errors.New("from account balance less than transfer amount")
	}

	toAccount, ok := ts.accounts[to]
	if !ok {
		return errors.New("to account not exist")
	}

	txn := newTransaction(from, to, amount, Transfer)
	ts.transactions = append(ts.transactions, txn)
	fromAccount.Balance -= amount
	fromAccount.Expense += amount
	toAccount.Balance += amount
	return nil
}

func (ts *TradeSystem) ExpenseRanking() []*Account {
	accounts := make([]*Account, 0)
	for _, acc := range ts.accounts {
		accounts = append(accounts, acc)
	}

	sort.Slice(accounts, func(i, j int) bool {
		if accounts[i].Expense == accounts[j].Expense {
			return accounts[i].Name < accounts[j].Name
		}
		return accounts[i].Expense > accounts[j].Expense
	})

	return accounts
}
