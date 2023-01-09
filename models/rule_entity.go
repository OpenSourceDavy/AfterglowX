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

func (rule *RuleEntity) GerRuleEntities(data map[string]interface{}) (res []RuleEntity, err error) {

	err = db.Where("name = ?", data["rule_id"].(string)).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return
}

func (rule *RuleEntity) CreateRuleEntity(data map[string]interface{}) error {
	id := uuid.NewString()

	err := db.Create(&RuleEntity{
		RuleID:     id,
		UserID:     data["user_id"].(string),
		Coordinate: data["coordinate"].(string),
		Distance:   data["distance"].(float32),
		Quality:    data["quality"].(float32),
		Type:       data["type"].(string),
		Time:       data["time"].(string),
		Frquency:   data["frequency"].(int),
		State:      data["state"].(int),
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (rule *RuleEntity) UpdateRuleEntity(data map[string]interface{}) error {
	id := uuid.NewString()

	err := db.Update(&RuleEntity{
		RuleID:     id,
		UserID:     data["user_id"].(string),
		Coordinate: data["coordinate"].(string),
		Distance:   data["distance"].(float32),
		Quality:    data["quality"].(float32),
		Type:       data["type"].(string),
		Time:       data["time"].(string),
		Frquency:   data["frequency"].(int),
		State:      data["state"].(int),
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (rule *RuleEntity) DeleteRuleEntity(data map[string]interface{}) error {

	err := db.Where("rule_id = ?", data["rule_id"].(string)).Delete(&rule).Error

	if err != nil {
		return err
	}

	return nil
}
