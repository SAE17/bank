package types

// Money представляет собой денежную сумму в минимальных единицах

type Money int64

// Currency представляет код валюты
type Currency string
type Category string

// Код валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string

type Status string

const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Card представляет информацию о платежной карте

type Card struct {
	ID       int
	PAN      PAN
	Balance  int64
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
