package entities

import "time"

type GlobalEntity struct {
	Id       *int64     `json:"id" db:"id"`
	CreateAt *time.Time `json:"creat_at" db:"create_at"`
	ModifyAt *time.Time `json:"modify_at" db:"modify_at"`
}
