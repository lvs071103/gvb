package models

import (
	"time"
)

type MODEL struct {
	ID       uint      `gorm:"primaryKey" json:"id"` // 主键ID
	CreateAt time.Time `json:"creat_at"`             // 创建时间
	UpdateAt time.Time `json:"update_at"`            //更新时间
}
