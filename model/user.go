package model

import "time"

type User struct {
	ID             int64     `json:"id" gorm:"id"`
	Username       string    `json:"username" gorm:"username"`
	Remark         string    `json:"remark" gorm:"remark"`
	LastSubmitTime time.Time `json:"last_submit_time" gorm:"last_submit_time"`
	CreateTime     time.Time `json:"create_time" gorm:"create_time"`
	UpdateTime     time.Time `json:"update_time" gorm:"update_time"`
}

func (u User) TableName() string {
	return "lw_user"
}

func (u *User) IsNeverSubmit() bool {
	return u.LastSubmitTime.IsZero()
}
