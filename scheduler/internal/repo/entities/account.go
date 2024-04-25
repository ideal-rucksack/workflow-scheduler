package entities

type AccountStatus string

const (
	// ACTIVE 用户已激活，可以正常使用系统。
	ACTIVE AccountStatus = "active"
	// INACTIVE 用户未激活，可能是因为用户未激活邮箱。
	INACTIVE AccountStatus = "inactive"
	// SUSPENDED 用户被暂时挂起，可能是因为违反了系统的使用规定。
	SUSPENDED AccountStatus = "suspended"
)

type Account struct {
	GlobalEntity
	Nickname     *string        `json:"nickname" db:"nickname"`
	Code         *string        `json:"code" db:"code"`
	Email        *string        `json:"email" db:"email"`
	Username     *string        `json:"username" db:"username"`
	Password     *string        `json:"password" db:"password"`
	Secret       *string        `json:"secret" db:"secret"`
	Status       *AccountStatus `json:"status" db:"status"`
	RefreshToken *string        `json:"refresh_token" db:"refresh_token"`
	Deleted      *bool          `json:"deleted" db:"deleted"`
}
