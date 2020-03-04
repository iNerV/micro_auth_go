package core

import (
	"github.com/jackc/pgtype"
)

type BaseModel struct {
	ID        pgtype.UUID      `faker:"-"`
	CreatedAt pgtype.Timestamp `db:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at"`
	DeletedAt pgtype.Timestamp `db:"deleted_at"`
}
