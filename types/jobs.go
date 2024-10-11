package types

import (
	"database/sql"
	"time"
)

type Job struct {
	ID        int            `db:"id"`
	Name      string         `db:"name"`
	Token     string         `db:"token"`
	Step      int            `db:"step"`
	Error     sql.NullString `db:"error"`
	Status    int            `db:"status"`
	Created   time.Time      `db:"created"`
	Completed sql.NullTime   `db:"completed"`
}
