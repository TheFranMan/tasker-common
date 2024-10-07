package types

type ActionType string

var (
	ActionDelete ActionType = "delete"
)

type Extras map[string]string
