package service

import (
	"github.com/wwkeyboard/sunsetwx/domain"
)

type alarmRuleUseCase struct {
	alarmRuleRepository domain.AlarmRuleRepository
}

func NewAlarmRuleUseCase(alarmRuleRepository domain.AlarmRuleRepository) domain.AlarmRuleUseCase {
	return &alarmRuleUseCase{
		alarmRuleRepository: alarmRuleRepository,
	}
}

func (as *alarmRuleUseCase) CreateRule(rule *domain.AlarmRule) error {
	return as.alarmRuleRepository.CreateRule(rule)
}

func (as *alarmRuleUseCase) GetAlarmRuleByUserID(userID string) (res domain.AlarmRule, err error) {
	return as.alarmRuleRepository.GetAlarmRuleByUserID(userID)
}

func (as *alarmRuleUseCase) UpdateAlarmRule(rule *domain.AlarmRule) error {
	return as.alarmRuleRepository.UpdateAlarmRule(rule)
}

func (as *alarmRuleUseCase) DeleteAlarmRule(ruleID string) error {
	return as.alarmRuleRepository.DeleteAlarmRule(ruleID)
}
