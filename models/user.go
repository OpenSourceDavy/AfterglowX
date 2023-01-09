package models

import (
	"log"
	"time"

	"github.com/google/uuid"
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
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (user *User) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func CreateUser(data map[string]interface{}) (string, error) {
	id := uuid.NewString()

	err := db.Create(&User{
		UserID:   id,
		UserName: data["name"].(string),
		Country:  data["country"].(string),
		Cell:     data["cell"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
		State:    data["state"].(int),
	}).Error

	if err != nil {
		return "", err
	}

	return id, err
}

func GetUser(data map[string]interface{}) (res User, err error) {
	userID := data["user_id"].(string)

	if len(userID) > 0 {
		err = db.Where("user_id = ?", userID).First(&res).Error
	} else {
		log.Fatal("GetUser error")
		return res, err
	}

	if err != nil {
		log.Fatal("GetUser error")
	}

	return res, err
}
