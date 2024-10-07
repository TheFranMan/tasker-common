package types

type Steps []Step

type Step struct {
	Name string   `json:"name"`
	Jobs []string `json:"jobs"`
}

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
