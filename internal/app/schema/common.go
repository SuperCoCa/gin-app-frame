package schema

import "gorm.io/gorm"

// ID主键
type Id struct {
	Id uint `json:"id" gorm:"primaryKey"`
}

// 添加/更新时间
type Timestamps struct {
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

// 软删除
type SoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
