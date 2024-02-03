package models

import (
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
