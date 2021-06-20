package testable

import "gorm.io/gorm"

type User struct {
	Name  string
	Phone string
}

// CreateUser 不可测试
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
