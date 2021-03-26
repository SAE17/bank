package types

type Money int64

// Currency представляет код валюты
type Currency string

// Код валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string
type Card struct {
	ID         int
	PAN        PAN
	Balance    int64
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance int64
}
