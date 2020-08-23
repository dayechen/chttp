package model

import "github.com/jinzhu/gorm"

// User 用户结构体
type User struct {
	*Model
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

// Create 创建用户
func (u User) Create(db *gorm.DB) (int, error) {
	if err := db.Create(&u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}
