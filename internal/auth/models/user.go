package models

import (
	"micro_auth/core"
)

type User struct {
	core.BaseModel
	IsActive bool   `db:"is_active"`
	IsAdmin  bool   `db:"is_admin"`
	IsStaff  bool   `db:"is_staff"`
	Username string `db:"username" faker:"username"`
	Email    string `db:"email" faker:"email"`
	Password string `db:"password" faker:"password"`
}
