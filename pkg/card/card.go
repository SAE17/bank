package card

import "github.com/SAE17/bank/pkg/types"

func Withdraw(card types.Card, amount int64) types.Card {
	if !card.Active {
		return card
	}

	if amount > card.Balance {
		return card
	}

	if amount < 0 {
		return card
	}

	if amount > 20_000_00 {
		return card
	}

	card.Balance -= amount
	return card
}
func Issue(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}
