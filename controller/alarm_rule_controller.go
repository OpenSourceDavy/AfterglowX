package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wwkeyboard/sunsetwx/domain"
)

type AlarmRuleController struct {
	AlarmRuleUseCase domain.AlarmRuleUseCase
}

func (ac *AlarmRuleController) CreateAlarmRule(c *gin.Context) {
	var alarmRule domain.AlarmRule

	err := c.ShouldBind(&alarmRule)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	alarmRule.RuleID = uuid.NewString()

	err = ac.AlarmRuleUseCase.CreateRule(&alarmRule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Alarm rule created successfully",
	})
}

func (ac *AlarmRuleController) DeleteAlarmRule(c *gin.Context) {
	var alarmRule domain.AlarmRule

	err := c.ShouldBind(&alarmRule)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = ac.AlarmRuleUseCase.DeleteAlarmRule(alarmRule.RuleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Alarm rule deleted successfully",
	})
}

func (ac *AlarmRuleController) UpdateAlarmRule(c *gin.Context) {
	var alarmRule domain.AlarmRule

	err := c.ShouldBind(&alarmRule)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = ac.AlarmRuleUseCase.UpdateAlarmRule(&alarmRule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Alarm rule updated successfully",
	})
}

func (ac *AlarmRuleController) GetAlarmRuleByUserID(c *gin.Context) {
	var alarmRule domain.AlarmRule

	err := c.ShouldBind(&alarmRule)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	res, err := ac.AlarmRuleUseCase.GetAlarmRuleByUserID(alarmRule.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
