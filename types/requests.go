package types

import (
	"database/sql"
	"encoding/json"
	"time"
)

var (
	RequestStatusNew        RequestStatus = 0
	RequestStatusInProgress RequestStatus = 1
	RequestStatusCompleted  RequestStatus = 2
	RequestStatusFailed     RequestStatus = 3

	JobStatusNew        JobStatus = 0
	JobStatusInProgress JobStatus = 1
	JobStatusCompleted  JobStatus = 2
	JobStatusFailed     JobStatus = 3
	JobStatusRetry      JobStatus = 4

	ActionDelete ActionType = "delete"
)

type RequestStatus int
type JobStatus int
type ActionType string

type Extras map[string]string

type Request struct {
	Token        string         `db:"token"`
	RequestToken string         `db:"request_token"`
	Action       string         `db:"action"`
	Params       Params         `db:"params"`
	Extras       sql.NullString `db:"extras"`
	Steps        Steps          `db:"steps"`
	Step         int            `db:"step"`
	Status       int            `db:"status"`
	Created      time.Time      `db:"created"`
	Completed    sql.NullTime   `db:"completed"`
}

type Params struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (p *Params) Scan(value interface{}) error {
	if nil == value {
		p = &Params{}
		return nil
	}

	return json.Unmarshal(value.([]byte), p)
}
