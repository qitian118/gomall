package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
	Rule           string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func Creat(db *gorm.DB, user *User) error {
	user.Rule = "user"
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetByUserid(db *gorm.DB, userId uint) (*User, error) {
	var user User
	err := db.Where("id = ?", userId).First(&user).Error
	return &user, err
}
