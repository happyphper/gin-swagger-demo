package models

import "database/sql"

type User struct {
	ID           int            `json:"id" gorm:"" gorm:"primaryKey" example:"11012341234"`
	Username     string         `json:"username" gorm:"not null;uniqueIndex;comment:用户名" example:"11012341234"`
	Password     sql.NullString `json:"password" gorm:"comment:密码" swaggertype:"string"`
	WechatOpenid sql.NullString `json:"wechat_openid" gorm:"comment:微信 OPEN ID" swaggertype:"primitive,string" example:"asdqwez@qsdasdqwe"`
	Status       int            `json:"status" gorm:"not null;default:1;comment:状态/1正常/2禁用" example:"1"`
	CreatedAt    string         `json:"created_at" example:"2006-01-02 15:04:05"`
	UpdatedAt    string         `json:"updated_at" example:"2006-01-02 15:04:05"`
}
