package response

import "github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"

type Signup struct {
	Nickname string                 `json:"nickname"`
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	Status   entities.AccountStatus `json:"status"`
}

type Verify struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}

type SignIn struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}

type Current struct {
	Username string                 `json:"username"`
	Nickname string                 `json:"nickname"`
	Email    string                 `json:"email"`
	Code     *string                `json:"code"`
	Status   entities.AccountStatus `json:"status"`
}

type RefreshToken struct {
	AccessToken string `json:"access_token"`
}
