package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Wallet struct {
	Id      uuid.UUID
	Balance float32
}

type Transaction struct {
	Time   time.Time
	From   uuid.UUID
	To     uuid.UUID
	Amount float32
}
