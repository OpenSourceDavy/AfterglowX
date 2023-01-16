package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wwkeyboard/sunsetwx/domain"
	"github.com/wwkeyboard/sunsetwx/logs"
)

type alarmRuleRepository struct {
	DB *gorm.DB
}

func NewAlarmRuleRepository() domain.AlarmRuleRepository {
	return &alarmRuleRepository{
		DB: db11,
	}
}

func (r *alarmRuleRepository) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (r *alarmRuleRepository) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func (r *alarmRuleRepository) CreateRule(data *domain.AlarmRule) error {

	err := r.DB.Create(data).Error

	if err != nil {
		logs.Log.Error("CreateRule error, error message: %s", err)
		return err
	}

	return err
}

func (r *alarmRuleRepository) GetAlarmRuleByUserID(userID string) (res []domain.AlarmRule, err error) {
	err = r.DB.Where("user_id = ?", userID).Find(&res).Error

	if err != nil {
		logs.Log.Error("GetAlarmRuleByUserID error, error message: %s", err)
	}

	return res, err
}

func (r *alarmRuleRepository) UpdateAlarmRule(rule *domain.AlarmRule) error {
	err := r.DB.Model(&domain.AlarmRule{}).Where("rule_id = ?", rule.RuleID).Update(&rule).Error

	if err != nil {
		logs.Log.Error("UpdateAlarmRule error, error message: %s", err)
		return err
	}

	return nil
}

func (r *alarmRuleRepository) DeleteAlarmRule(ruleID string) error {
	err := r.DB.Where("rule_id = ?", ruleID).Delete(&domain.AlarmRule{}).Error

	if err != nil {
		logs.Log.Error("DeleteAlarmRule error, error message: %s", err)
		return err
	}

	return nil
}
