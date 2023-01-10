package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wwkeyboard/sunsetwx/logs"
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

func (user *User) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())

}

func (user *User) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func CreateUser(data User) error {

	err := db.Create(&data).Error

	if err != nil {
		logs.Log.Error("CreateUser error, error message: %s", err)
		return err
	}

	return err
}

func GetUser(data map[string]interface{}) (res User, err error) {
	email := data["email"].(string)
	cell := data["cell"].(string)

	if len(email) > 0 {
		err = db.Where("email = ?", email).First(&res).Error
	} else if len(cell) > 0 {
		err = db.Where("cell = ?", cell).First(&res).Error
	} else {
		logs.Log.Error("GetUser empty, error message: %s", err)
	}

	if err != nil {
		logs.Log.Error("GetUser error, error message: %s", err)
	}

	return res, err
}
