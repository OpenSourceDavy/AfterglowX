package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/logs"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		DB: db11,
	}
}

func (u *userRepository) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())

}

func (u *userRepository) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func (u *userRepository) CreateUser(data *domain.User) error {

	err := u.DB.Create(data).Error

	if err != nil {
		logs.Log.Error("CreateUser error, error message: %s", err)
		return err
	}

	return err
}

func (u *userRepository) GetUser(data map[string]interface{}) (res domain.User, err error) {
	email := data["email"].(string)
	cell := data["cell"].(string)

	if len(email) > 0 {
		err = u.DB.Where("email = ?", email).First(&res).Error
	} else if len(cell) > 0 {
		err = u.DB.Where("cell = ?", cell).First(&res).Error
	} else {
		logs.Log.Error("GetUser empty, error message: %s", err)
	}

	if err != nil {
		logs.Log.Error("GetUser error, error message: %s", err)
	}

	return res, err
}

func (u *userRepository) GetUserByEmail(email string) (res domain.User, err error) {
	err = u.DB.Where("email = ?", email).First(&res).Error

	if err != nil {
		logs.Log.Error("GetUser error, error message: %s", err)
	}

	return res, err
}
