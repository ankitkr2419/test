package db

import (
	"github.com/google/uuid"
)

const ContextKeyUsername = "username"

type ProcessSequence struct {
	ID             uuid.UUID `db:"id" json:"process_id" validate:"required"`
	SequenceNumber int64     `db:"sequence_num" json:"sequence_num" validate:"required,gte=1"`
}
