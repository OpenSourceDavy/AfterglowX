package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	UserName   string `json:"name"`
	Country    string `json:"country"`
	Cell       string `json:"cell"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	State      int    `json:"state"`
	CreatedOn  int64  `json:"created_on"`
	ModifiedOn int64  `json:"modified_on"`
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUser(data map[string]interface{}) (res User, err error)
	GetUserByEmail(email string) (user User, err error)
}

func (u *User) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (u *User) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}
