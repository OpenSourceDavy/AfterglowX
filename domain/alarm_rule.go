package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AlarmRule struct {
	ID         int     `json:"id"`
	RuleID     string  `json:"rule_id"`
	UserID     string  `json:"user_id"`
	Coordinate string  `json:"coordinate"`
	Distance   float32 `json:"distance"`
	Quality    float32 `json:"quality"`
	Type       string  `json:"type"`
	Time       int64   `json:"time"`
	Frequency  int     `json:"frequency"`
	State      int     `json:"state"`
	CreatedOn  int64   `json:"created_on"`
	ModifiedOn int64   `json:"modified_on"`
}

type AlarmRuleRepository interface {
	CreateRule(rule *AlarmRule) error
	GetAlarmRuleByUserID(userID string) (res []AlarmRule, err error)
	UpdateAlarmRule(rule *AlarmRule) error
	DeleteAlarmRule(ruleID string) error
}

type AlarmRuleUseCase interface {
	CreateRule(rule *AlarmRule) error
	GetAlarmRuleByUserID(userID string) (res []AlarmRule, err error)
	UpdateAlarmRule(rule *AlarmRule) error
	DeleteAlarmRule(ruleID string) error
}

func (r *AlarmRule) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (r *AlarmRule) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}
