package generics

type Transaction struct {
	From   string
	To     string
	Amount int
}

func BalanceOf(transactions []Transaction, name string) int {
	// var balance int
	// for _, t := range transactions {
	// 	if t.From == name {
	// 		balance -= t.Amount
	// 	} else if t.To == name {
	// 		balance += t.Amount
	// 	}
	// }
	// return balance
	return Reduce(transactions, func(balance int, transaction Transaction) int {
		if transaction.From == name {
			return balance - transaction.Amount
		}
		if transaction.To == name {
			return balance + transaction.Amount
		}
		return balance
	}, 0)
}
