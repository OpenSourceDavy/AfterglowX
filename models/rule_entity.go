package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type RuleEntity struct {
	ID         int     `json:"id"`
	RuleID     string  `json:"rule_id"`
	UserID     string  `json:"user_id"`
	Coordinate string  `json:"coordinate"`
	Distance   float32 `json:"distance"`
	Quality    float32 `json:"quality"`
	Type       string  `json:"type"`
	Time       string  `json:"time"`
	Frquency   int     `json:"frequency"`
	State      int     `json:"state"`
	CreatedOn  string  `json:"created_on"`
	ModifiedOn string  `json:"modified_on"`
}

func (rule *RuleEntity) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (rule *RuleEntity) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func GetRuleEntities(data map[string]interface{}) (res []RuleEntity, err error) {

	err = db.Where("user_id = ?", data["user_id"].(string)).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return
}

func CreateRuleEntity(data RuleEntity) error {
	data.RuleID = uuid.NewString()

	err := db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateRuleEntity(data RuleEntity) error {

	err := db.Where("rule_id = ?", data.RuleID).Update(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteRuleEntity(data map[string]interface{}) error {

	err := db.Where("rule_id = ?", data["rule_id"].(string)).Delete(&RuleEntity{}).Error

	if err != nil {
		return err
	}

	return nil
}
