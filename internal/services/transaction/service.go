package transaction

type Transaction struct {
	Amount        float64 `json:"amount"`
	DebitAddress  string  `json:"debitAddress"`
	CreditAddress string  `json:"creditAddress"`
}
