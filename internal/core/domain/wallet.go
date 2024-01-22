package domain

import "github.com/google/uuid"

type Wallet struct {
	Id      uuid.UUID
	Balance float64
}
