package requests

import (
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
)

type Signup struct {
	Nickname string                 `json:"nickname" validate:"required" message:"Nickname is illegal"`
	Username string                 `json:"username" validate:"required" message:"Username is illegal"`
	Email    string                 `json:"email" validate:"required,email" message:"Email is illegal"`
	Code     *string                `json:"code" message:"Code is illegal"`
	Status   entities.AccountStatus `json:"status"`
}

type Verify struct {
	Username   string `json:"username" validate:"required" message:"Username is illegal"`
	Password   string `json:"password" validate:"required" message:"Password is illegal"`
	VerifyCode string `json:"verify_code" validate:"required" message:"Verify code is illegal"`
}

type SignIn struct {
	Username string `json:"username" validate:"required" message:"Username is illegal"`
	Password string `json:"password" validate:"required" message:"Password is illegal"`
}

func (s Signup) Validate() error {
	return Validate(s)
}

func (s Verify) Validate() error {
	return Validate(s)
}

func (s SignIn) Validate() error {
	return Validate(s)
}
