package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wwkeyboard/sunsetwx/models"
)

func GetSunsetQuality(c *gin.Context) {
	quality := 0.55
	c.JSON(http.StatusOK, gin.H{
		"quality": quality,
	})
}

func CreateUser(c *gin.Context) {
	data := make(map[string]interface{})
	data["name"] = c.Query("name")
	data["country"] = c.Query("country")
	data["cell"] = c.Query("cell")
	data["email"] = c.Query("email")
	data["password"] = c.Query("password")
	data["state"] = c.Query("state")

	userID, err := models.CreateUser(data)
	if err != nil {
		log.Printf("CreateUser Error, err message: %s", err)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "failed",
			"data": -1,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": userID,
		})
	}
}

func UserLogin(c *gin.Context) {
	data := make(map[string]interface{})
	data["user_id"] = c.Query("user_id")

	user, err := models.GetUser(data)
	if err != nil {
		log.Printf("UserLogin Error, err message: %s", err)
	}

	pwd := user.Password

	success := -1
	if pwd == c.Query("password") {
		success = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

func GetAlarmRules(c *gin.Context) {
	data := make(map[string]interface{})
	data["user_id"] = c.Query("user_id")

	entities, err := models.GetRuleEntities(data)
	if err != nil {
		log.Printf("GetAlarmRules Error, err message: %s", err)
	}

	c.JSON(http.StatusOK, entities)
}

func CreateAlarmRule(c *gin.Context) {
	var data models.RuleEntity
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("CreateAlarmRule bind JSON error, err message: %s", err)
	}

	success := 1
	err = models.CreateRuleEntity(data)
	if err != nil {
		success = -1
		log.Printf("CreateRuleEntity error, err message: %s", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

func UpdateAlarmRule(c *gin.Context) {
	var data models.RuleEntity
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("UpdateAlarmRule bind JSON error, err message: %s", err)
	}

	success := 1
	err = models.UpdateRuleEntity(data)
	if err != nil {
		success = -1
		log.Printf("UpdateAlarmRule error, err message: %s", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}

func DeleteAlarmRule(c *gin.Context) {
	data := make(map[string]interface{})
	data["rule_id"] = c.Query("rule_id")
	err := models.DeleteRuleEntity(data)
	if err != nil {
		log.Printf("DeleteAlarmRule error, err message: %s", err)
	}
}
