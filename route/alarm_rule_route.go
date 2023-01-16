package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/controller"
	"github.com/wwkeyboard/sunsetwx/repository"
	"github.com/wwkeyboard/sunsetwx/service"
)

func NewAlarmRuleRouter(group *gin.RouterGroup) {
	rp := repository.NewAlarmRuleRepository()
	rc := controller.AlarmRuleController{
		AlarmRuleUseCase: service.NewAlarmRuleUseCase(rp),
	}
	group.POST("/createRule", rc.CreateAlarmRule)
	group.POST("/deleteRule", rc.DeleteAlarmRule)
	group.POST("/rule", rc.GetAlarmRuleByUserID)
	group.POST("/updateRule", rc.UpdateAlarmRule)
}
