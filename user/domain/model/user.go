package model

type User struct {
	// 主键
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	// 用户名称
	Username string `gorm:"unique_index;not_null"`
	// 需要字段
	FristName string

	// 密码
	HashPassword string
}
