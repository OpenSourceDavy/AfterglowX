package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/middleware"
	"github.com/wwkeyboard/sunsetwx/settings"
)

func Setup(routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	// All Public APIs
	NewRegisterRouter(publicRouterV1)
	NewLoginRouter(publicRouterV1)

	protectedRouterV1 := routerV1.Group("")
	// Middleware to verify AccessToken
	protectedRouterV1.Use(middleware.JwtAuthMiddleware(settings.JwtSecret))
	// All Private APIs
}

// func GetSunsetQuality(c *gin.Context) {
// 	data := make(map[string]interface{})
// 	err := c.ShouldBindJSON(&data)
// 	if err != nil {
// 		logs.Log.Error("GetSunsetQuality bind JSON error, err message: %s", err)
// 	}
// 	m, _ := quality.GetMetrics(data["lat"].(float64), data["lon"].(float64), data["type"].(string))
// 	qua := quality.GetQuality(m)
// 	c.JSON(http.StatusOK, gin.H{
// 		"quality": qua,
// 	})
// }

// func GetAlarmRules(c *gin.Context) {
// 	data := make(map[string]interface{})
// 	err := c.BindJSON(&data)
// 	if err != nil {
// 		logs.Log.Error("GetAlarmRules bind JSON error, err message: %s", err)
// 	}

// 	entities, err := models.GetRuleEntities(data)
// 	if err != nil {
// 		logs.Log.Error("GetAlarmRules Error, err message: %s", err)
// 	}

// 	c.JSON(http.StatusOK, entities)
// }

// func CreateAlarmRule(c *gin.Context) {
// 	var data models.RuleEntity
// 	err := c.ShouldBindJSON(&data)
// 	if err != nil {
// 		logs.Log.Error("CreateAlarmRule bind JSON error, err message: %s", err)
// 	}

// 	success := 1
// 	err = models.CreateRuleEntity(data)
// 	if err != nil {
// 		success = -1
// 		logs.Log.Error("CreateRuleEntity error, err message: %s", err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": success,
// 	})
// }

// func UpdateAlarmRule(c *gin.Context) {
// 	var data models.RuleEntity
// 	err := c.ShouldBindJSON(&data)
// 	if err != nil {
// 		logs.Log.Error("UpdateAlarmRule bind JSON error, err message: %s", err)
// 	}

// 	success := 1
// 	err = models.UpdateRuleEntity(data)
// 	if err != nil {
// 		success = -1
// 		logs.Log.Error("UpdateAlarmRule error, err message: %s", err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": success,
// 	})
// }

// func DeleteAlarmRule(c *gin.Context) {
// 	data := make(map[string]interface{})
// 	err := c.BindJSON(&data)
// 	if err != nil {
// 		logs.Log.Error("DeleteAlarmRule bind JSON error, err message: %s", err)
// 	}
// 	success := 1
// 	err = models.DeleteRuleEntity(data)
// 	if err != nil {
// 		success = -1
// 		logs.Log.Error("DeleteAlarmRule error, err message: %s", err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": success,
// 	})
// }
