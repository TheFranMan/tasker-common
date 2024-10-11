package types

import "encoding/json"

var (
	StepsDelete = Steps{
		{
			Name: "service1_retrieve_user",
			Jobs: []string{
				"service1GetUser",
			},
		},
		{
			Name: "delete_user_accounts",
			Jobs: []string{
				"service1DeleteUser",
				"service2DeleteUser",
				"service3DeleteUser",
			},
		},
	}
)

type Steps []Step

type Step struct {
	Name string   `json:"name"`
	Jobs []string `json:"jobs"`
}

func (s *Steps) Scan(value interface{}) error {
	if nil == value {
		s = &Steps{}
		return nil
	}

	return json.Unmarshal(value.([]byte), s)
}
