package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Wallet struct {
	Id      uuid.UUID `json:"id"`
	Balance float32   `json:"balance"`
}

type Transaction struct {
	Time   time.Time `json:"time"`
	From   uuid.UUID `json:"from"`
	To     uuid.UUID `json:"to"`
	Amount float32   `json:"amount"`
}

// For possibility specify time format

func (t Transaction) MarshalJSON() ([]byte, error) {
	formatted := t.Time.Format(time.RFC3339)

	type jsonTimeFormattedTransactions struct {
		Time   string    `json:"time"`
		From   uuid.UUID `json:"from"`
		To     uuid.UUID `json:"to"`
		Amount float32   `json:"amount"`
	}

	tr := jsonTimeFormattedTransactions{
		Time:   formatted,
		From:   t.From,
		To:     t.To,
		Amount: t.Amount,
	}

	return json.Marshal(&tr)
}
