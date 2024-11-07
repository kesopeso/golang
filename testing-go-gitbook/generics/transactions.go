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
	return Reduce(transactions, func(b int, t Transaction) int {
		if t.From == name {
			return b - t.Amount
		}
		if t.To == name {
			return b + t.Amount
		}
		return b
	}, 0)
}
