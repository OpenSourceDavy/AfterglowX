package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wwkeyboard/sunsetwx/logs"
	"github.com/wwkeyboard/sunsetwx/models"
	"github.com/wwkeyboard/sunsetwx/quality"
)

func GetSunsetQuality(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.ShouldBindJSON(&data)
	if err != nil {
		logs.Log.Error("GetSunsetQuality bind JSON error, err message: %s", err)
	}
	m, _ := quality.GetMetrics(data["lat"].(float64), data["lon"].(float64), data["type"].(string))
	qua := quality.GetQuality(m)
	c.JSON(http.StatusOK, gin.H{
		"quality": qua,
	})
}

func CreateUser(c *gin.Context) {
	var data models.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		logs.Log.Error("CreateUser bind JSON error, err message: %s", err)
	}
	data.State = 1
	data.UserID = uuid.NewString()
	err = models.CreateUser(data)
	if err != nil {
		logs.Log.Error("CreateUser Error, err message: %s", err)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "failed",
			"data": -1,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": data.UserID,
		})
	}
}

// todo: user not registered
// todo: jwt
func UserLogin(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.BindJSON(&data)
	if err != nil {
		logs.Log.Error("UserLogin bind JSON error, err message: %s", err)
	}

	user, err := models.GetUser(data)
	if err != nil {
		logs.Log.Error("UserLogin Error, err message: %s", err)
	}

	pwd := user.Password

	success := -1
	if pwd == data["password"] {
		success = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

func GetAlarmRules(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.BindJSON(&data)
	if err != nil {
		logs.Log.Error("GetAlarmRules bind JSON error, err message: %s", err)
	}

	entities, err := models.GetRuleEntities(data)
	if err != nil {
		logs.Log.Error("GetAlarmRules Error, err message: %s", err)
	}

	c.JSON(http.StatusOK, entities)
}

func CreateAlarmRule(c *gin.Context) {
	var data models.RuleEntity
	err := c.ShouldBindJSON(&data)
	if err != nil {
		logs.Log.Error("CreateAlarmRule bind JSON error, err message: %s", err)
	}

	success := 1
	err = models.CreateRuleEntity(data)
	if err != nil {
		success = -1
		logs.Log.Error("CreateRuleEntity error, err message: %s", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

func UpdateAlarmRule(c *gin.Context) {
	var data models.RuleEntity
	err := c.ShouldBindJSON(&data)
	if err != nil {
		logs.Log.Error("UpdateAlarmRule bind JSON error, err message: %s", err)
	}

	success := 1
	err = models.UpdateRuleEntity(data)
	if err != nil {
		success = -1
		logs.Log.Error("UpdateAlarmRule error, err message: %s", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

func DeleteAlarmRule(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.BindJSON(&data)
	if err != nil {
		logs.Log.Error("DeleteAlarmRule bind JSON error, err message: %s", err)
	}
	success := 1
	err = models.DeleteRuleEntity(data)
	if err != nil {
		success = -1
		logs.Log.Error("DeleteAlarmRule error, err message: %s", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}
