package domain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id       uuid.UUID
	From     uuid.UUID
	To       uuid.UUID
	Amount   float64
	DateTime time.Time
}
