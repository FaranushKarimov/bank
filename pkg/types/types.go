package types

// Money представляет собой денежную сумму
type Money int64

// Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.)
type Category string

// Payment представляет информацию о платёжной карте
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

// Currency представляет код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN предсталяет собой номер карты
type PAN string

// Card представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // использовали Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money // использовали Money
}

// PaymentSource struct
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
