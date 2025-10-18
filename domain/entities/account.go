package entities

type Account struct {
	ID       uint32
	Owner    string
	Currency string
	Balance  float32
}
